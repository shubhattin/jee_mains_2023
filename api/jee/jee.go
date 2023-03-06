package jee

import (
	"api/kry"
	"github.com/gin-gonic/gin"
)

func GetRoutes(mainRoute *gin.Engine) {
	router := mainRoute.Group("/api")

	router.POST("/page_view_count", page_view_count)
	router.POST("/get_sample_result", get_sample_result)
}

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

type GetSampleResultType struct {
	Result ResultDataType `json:"result"`
	Data   MainDataType   `json:"data"`
}

func get_sample_result(c *gin.Context) {
	base := kry.Deta.Base("data")
	var data GetSampleResultType
	base.Get("sample_result", &data)
	c.JSON(200, data)
}

type PageViewCountType struct {
	PageViewCount   int `json:"page_view_count"`
	ResultViewCount int `json:"result_view_count"`
}

type KeyValuePairInt struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

func page_view_count(c *gin.Context) {
	base := kry.Deta.Base("counts")

	var page_view_count KeyValuePairInt
	var result_view_count KeyValuePairInt
	base.Get("page_view_count", &page_view_count)
	base.Get("result_view_count", &result_view_count)
	c.JSON(200, &PageViewCountType{
		PageViewCount:   page_view_count.Value,
		ResultViewCount: result_view_count.Value,
	})
}
