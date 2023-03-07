package jee

type ResultMetricsType struct {
	Unattempted []string `json:"unattempted"`
	Correct     []string `json:"correct"`
	Incorrect   []string `json:"incorrect"`
	Score       int16    `json:"score"`
}
type ResultDataType struct {
	Score            int16                `json:"score"`
	UnattemptedCount int16                `json:"unattempted_count"`
	CorrectCount     int16                `json:"correct_count"`
	IncorrectCount   int16                `json:"incorrect_count"`
	Subjects         [3]ResultMetricsType `json:"subjects"`
}
type MainDataType struct {
	GivenAnswer     [90]string `json:"GivenAnswer"`
	CorrectAnswer   [90]string `json:"CorrectAnswer"`
	Type            [90]string `json:"Type"`
	QuestionID      [90]string `json:"QuestionID"`
	CorrectAnswerID [90]string `json:"CorrectAnswerID"`
	GivenAnswerID   [90]string `json:"GivenAnswerID"`
	QuestionIMG     [90]string `json:"QuestionIMG"`
	Option1IMG      [90]string `json:"Option1IMG"`
	Option2IMG      [90]string `json:"Option2IMG"`
	Option3IMG      [90]string `json:"Option3IMG"`
	Option4IMG      [90]string `json:"Option4IMG"`
}

type UserInfoType struct {
	ApplicationNumber string `json:"key"`
	Name              string `json:"name"`
	RollNumber        string `json:"roll"`
	DOB               string `json:"DOB"`
	DateOfExam        string `json:"date"`
	TimeOfExam        string `json:"time"`
}

type AnswerKeyType struct {
	QuestionID      [90]string
	CorrectAnswerID [90]string
}
