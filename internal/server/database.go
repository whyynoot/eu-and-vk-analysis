package backend

import (
	"database/sql"
	"eu-and-vk-analysis/internal/databaseModels"
	"eu-and-vk-analysis/internal/models"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"log"
	"os"
	"strconv"
)

// TODO: Release database work

type DataBaseManager struct {
	DbURL         string
	sqlConnection *sql.DB
	reform        *reform.DB
}

type DataBaseConfig struct {
	DbURL string `envconfig:"DATABASE_URL" env-required:"true"`
}

func NewDataBaseManager() (*DataBaseManager, error) {
	cfg := DataBaseConfig{}
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &DataBaseManager{DbURL: cfg.DbURL}, nil
}

func (db *DataBaseManager) Connect() error {
	var err error
	db.sqlConnection, err = sql.Open("postgres", db.DbURL)
	if err != nil {
		return err
	}
	err = db.sqlConnection.Ping()
	if err != nil {
		return err
	}

	logger := log.New(os.Stderr, "SQL: ", log.Flags())

	db.reform = reform.NewDB(db.sqlConnection, postgresql.Dialect, reform.NewPrintfLogger(logger.Printf))

	return nil
}

func (db *DataBaseManager) GetStudentsPerformanceByGroup(group int) ([]models.Student, error) {
	ExtraQuery := ` INNER JOIN GroupsStudents g on s.id = g.student_id WHERE g.group_id = $1`

	Students, err := db.getStudentMarksByQuery(ExtraQuery, group)
	if err != nil {
		return nil, err
	}

	return Students, nil
}

func (db *DataBaseManager) GetStudentsWithPerformance(performance int) ([]models.Student, error) {
	var ExtraQuery string
	if performance == 5 {
		ExtraQuery = ` WHERE (m.credit_1 > 0 OR m.credit_1 is null) AND (m.credit_2 > 0 OR m.credit_2 is null) AND (m.credit_3 > 0 OR m.credit_3 is null) AND (m.credit_4 > 0 OR m.credit_4 is null) AND (m.credit_5 > 0 OR m.credit_5 is null) AND (m.credit_6 > 0 OR m.credit_6 is null) AND (m.credit_7 > 0 OR m.credit_7 is null) AND (m.credit_8 > 0 OR m.credit_8 is null) AND (m.credit_9 > 0 OR m.credit_9 is null) AND (m.credit_10 > 0 OR m.credit_10 is null)
			  AND (exam_1 = 5 OR exam_1 is null) AND (exam_2 = 5 OR exam_2 is null) AND (exam_3 = 5 OR exam_3 is null)AND (exam_4 = 5 OR exam_4 is null)AND (exam_5 = 5 OR exam_5 is null)AND (exam_6 = 5 OR exam_6 is null)AND (exam_7 = 5 OR exam_7 is null)AND (exam_8 = 5 OR exam_8 is null)`
	} else {
		ExtraQuery = fmt.Sprintf(` WHERE m.credit_1 = %[1]s OR m.credit_2 = %[1]s OR m.credit_3 = %[1]s OR m.credit_4 = %[1]s OR m.credit_5 = %[1]s OR m.credit_6 = %[1]s OR m.credit_7 = %[1]s OR m.credit_8 = %[1]s OR m.credit_9 = %[1]s OR m.credit_10 = %[1]s
			OR m.exam_1 = %[1]s OR m.exam_2 = %[1]s OR m.exam_3 = %[1]s OR m.exam_4 = %[1]s OR m.exam_5 = %[1]s OR m.exam_6 = %[1]s OR m.exam_7 = %[1]s OR m.exam_8 = %[1]s `, strconv.Itoa(performance))
	}

	Students, err := db.getStudentMarksByQuery(ExtraQuery)
	if err != nil {
		return nil, err
	}

	for i := range Students {
		Groups, err := db.GetStudentGroupsByID(Students[i].ID)
		if err != nil {
			return nil, err
		}
		Students[i].VKGroups = Groups
	}
	return Students, nil
}

func (db *DataBaseManager) GetStudentGroupsByID(id int32) ([]models.VKGroup, error) {
	q := "SELECT g.id, g.category FROM vkgroups g INNER JOIN groupsstudents on groupsstudents.group_id = g.id where groupsstudents.student_id = $1"

	GroupsRows, err := db.sqlConnection.Query(q, id)
	if err != nil {
		return nil, err
	}
	defer GroupsRows.Close()

	var Groups []models.VKGroup

	for GroupsRows.Next() {
		var Group models.VKGroup
		if err := GroupsRows.Scan(&Group.ID, &Group.Theme); err != nil {
			return nil, err
		}
		Groups = append(Groups, Group)
	}

	return Groups, nil
}

// private method
func (db *DataBaseManager) getStudentMarksByQuery(extraQuery string, args ...interface{}) ([]models.Student, error) {
	var Students []models.Student

	q := `SELECT s.id,
    m.credit_1, m.credit_2, m.credit_3, m.credit_4, m.credit_5, m.credit_6, m.credit_7, m.credit_8, m.credit_9, m.credit_10,
    m.exam_1, m.exam_2, m.exam_3, m.exam_4, m.exam_5, m.exam_6, m.exam_7, m.exam_8 FROM students s INNER JOIN marks m ON s.id = m.student_id` + extraQuery

	StudentMarksRows, err := db.sqlConnection.Query(q, args...)
	if err != nil {
		return nil, err
	}
	defer StudentMarksRows.Close()

	for StudentMarksRows.Next() {
		var marks databaseModels.Marks
		var student models.Student

		if err := StudentMarksRows.Scan(&student.ID, &marks.Credit1, &marks.Credit2, &marks.Credit3, &marks.Credit4, &marks.Credit5,
			&marks.Credit6, &marks.Credit7, &marks.Credit8, &marks.Credit9, &marks.Credit10, &marks.Exam1, &marks.Exam2, &marks.Exam3,
			&marks.Exam4, &marks.Exam5, &marks.Exam6, &marks.Exam7, &marks.Exam8); err != nil {
			return nil, err
		}

		creditSlice, examSlice := marks.ConvertToMassives()
		student.Marks = models.Marks{
			Credits: creditSlice,
			Exams:   examSlice,
		}
		Students = append(Students, student)
	}
	return Students, nil
}

func (db *DataBaseManager) CloseDB() {
	_ = db.sqlConnection.Close()
}
