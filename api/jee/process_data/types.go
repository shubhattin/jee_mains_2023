package jee

type ResultMetricsType struct {
	Unattempted []string `json:"unattempted"`
	Correct     []string `json:"correct"`
	Incorrect   []string `json:"incorrect"`
	Score       int      `json:"score"`
}
type ResultDataType struct {
	Score            int                 `json:"score"`
	UnattemptedCount int                 `json:"unattempted_count"`
	CorrectCount     int                 `json:"correct_count"`
	IncorrectCount   int                 `json:"incorrect_count"`
	Subjects         []ResultMetricsType `json:"subjects"`
}
type MainDataType struct {
	GivenAnswer     []string `json:"GivenAnswer"`
	CorrectAnswer   []string `json:"CorrectAnswer"`
	Type            []string `json:"Type"`
	QuestionID      []string `json:"QuestionID"`
	CorrectAnswerID []string `json:"CorrectAnswerID"`
	GivenAnswerID   []string `json:"GivenAnswerID"`
	QuestionIMG     []string `json:"QuestionIMG"`
	Option1IMG      []string `json:"Option1IMG"`
	Option2IMG      []string `json:"Option2IMG"`
	Option3IMG      []string `json:"Option3IMG"`
	Option4IMG      []string `json:"Option4IMG"`
}

type UserInfoType struct {
	ApplicationNumber string `json:"key"`
	Name              string `json:"name"`
	RollNumber        string `json:"roll"`
	DOB               string `json:"DOB"`
	DateOfExam        string `json:"date"`
	TimeOfExam        string `json:"time"`
}
