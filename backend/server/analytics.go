package backend

import (
	"eu-and-vk-analysis/backend/client_models"
	"github.com/pkg/errors"
	"log"
)

type Analytics struct {
	DataBaseManager *DataBaseManager
	performance     map[string]int
}

func NewAnalytics() (*Analytics, error) {
	db, err := NewDataBaseManager()
	if err != nil {
		return nil, err
	}

	err = db.Connect()
	if err != nil {
		return nil, err
	}

	return &Analytics{
		DataBaseManager: db,
		performance: map[string]int{"excellent": 5,
			"good":  4,
			"three": 3,
			"bad":   0,
		},
	}, nil
}

func (Analytics *Analytics) CloseDB() {
	Analytics.DataBaseManager.CloseDB()
}

func (Analytics *Analytics) AnalyseInterests(filter int) client_models.Response {
	Interests := map[string]int{"total_students": 0}
	students, err := Analytics.DataBaseManager.GetStudentsWithPerformance(filter)
	if err != nil {
		log.Println(err)
		return client_models.Response{Status: "NOT OK"}
	}

	for _, student := range students {
		Interests["total_students"]++
		studentThemeMap := make(map[string]bool)
		for _, group := range student.VKGroups {
			_, ok := studentThemeMap[group.Theme]
			if !ok {
				studentThemeMap[group.Theme] = true
				_, ok := Interests[group.Theme]
				if ok {
					Interests[group.Theme]++
				} else {
					Interests[group.Theme] = 1
				}
			}
		}
	}

	return client_models.Response{
		Statistics: Interests,
		Status:     "OK",
	}
}

func (Analytics *Analytics) AnalyseStudents(GroupID int) client_models.Response {
	Performance := map[string]int{"NA": 0, "three": 0, "good": 0, "excellent": 0}

	Students, err := Analytics.DataBaseManager.GetStudentsPerformanceByGroup(GroupID)
	if err != nil {
		log.Println(err)
		return client_models.Response{Status: "NOT OK"}
	}

	for _, student := range Students {
		studentFlag := false
		for _, credit := range student.Marks.Credits {
			if credit == 0 {
				Performance["NA"]++
				studentFlag = true
				break
			}
		}
		if !studentFlag {
			for _, exam := range student.Marks.Exams {
				if exam == 0 {
					Performance["NA"]++
					break
				}
				if exam == 3 {
					Performance["three"]++
					break
				}
				if exam == 4 {
					Performance["good"]++
					break
				}
			}
		}
		if !studentFlag {
			Performance["excellent"]++
		}
	}

	return client_models.Response{Status: "OK", Statistics: Performance}
}

func (Analytics *Analytics) CheckCorrectPerformance(InputPerformance string) (int, error) {
	status, ok := Analytics.performance[InputPerformance]
	if !ok {
		return -1, errors.Errorf("%s Filter Not Supported", InputPerformance)
	}
	return status, nil
}
