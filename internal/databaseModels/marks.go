package databaseModels

//go:generate reform

// Marks represents a row in marks table.
//reform:marks
type Marks struct {
	StudentID int32  `reform:"student_id,pk"`
	Credit1   *int32 `reform:"credit_1"`
	Credit2   *int32 `reform:"credit_2"`
	Credit3   *int32 `reform:"credit_3"`
	Credit4   *int32 `reform:"credit_4"`
	Credit5   *int32 `reform:"credit_5"`
	Credit6   *int32 `reform:"credit_6"`
	Credit7   *int32 `reform:"credit_7"`
	Credit8   *int32 `reform:"credit_8"`
	Credit9   *int32 `reform:"credit_9"`
	Credit10  *int32 `reform:"credit_10"`
	Exam1     *int32 `reform:"exam_1"`
	Exam2     *int32 `reform:"exam_2"`
	Exam3     *int32 `reform:"exam_3"`
	Exam4     *int32 `reform:"exam_4"`
	Exam5     *int32 `reform:"exam_5"`
	Exam6     *int32 `reform:"exam_6"`
	Exam7     *int32 `reform:"exam_7"`
	Exam8     *int32 `reform:"exam_8"`
}

func (s *Marks) ConvertToMassives() ([]int32, []int32) {
	creditSlice := make([]int32, 0)
	if s.Credit1 != nil {
		creditSlice = append(creditSlice, *s.Credit1)
	}
	if s.Credit2 != nil {
		creditSlice = append(creditSlice, *s.Credit2)
	}
	if s.Credit3 != nil {
		creditSlice = append(creditSlice, *s.Credit3)
	}
	if s.Credit4 != nil {
		creditSlice = append(creditSlice, *s.Credit4)
	}
	if s.Credit5 != nil {
		creditSlice = append(creditSlice, *s.Credit5)
	}
	if s.Credit6 != nil {
		creditSlice = append(creditSlice, *s.Credit6)
	}
	if s.Credit7 != nil {
		creditSlice = append(creditSlice, *s.Credit7)
	}
	if s.Credit8 != nil {
		creditSlice = append(creditSlice, *s.Credit8)
	}
	if s.Credit9 != nil {
		creditSlice = append(creditSlice, *s.Credit9)
	}
	if s.Credit10 != nil {
		creditSlice = append(creditSlice, *s.Credit10)
	}
	examSlice := make([]int32, 0)
	if s.Exam1 != nil {
		examSlice = append(examSlice, *s.Exam1)
	}
	if s.Exam2 != nil {
		examSlice = append(examSlice, *s.Exam2)
	}
	if s.Exam3 != nil {
		examSlice = append(examSlice, *s.Exam3)
	}
	if s.Exam4 != nil {
		examSlice = append(examSlice, *s.Exam4)
	}
	if s.Exam5 != nil {
		examSlice = append(examSlice, *s.Exam5)
	}
	if s.Exam6 != nil {
		examSlice = append(examSlice, *s.Exam6)
	}
	if s.Exam7 != nil {
		examSlice = append(examSlice, *s.Exam7)
	}
	if s.Exam8 != nil {
		examSlice = append(examSlice, *s.Exam8)
	}
	return creditSlice, examSlice
}
