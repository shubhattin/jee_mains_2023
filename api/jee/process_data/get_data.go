package jee

import (
	"api/kry"
	"os"
	"strconv"
	"strings"

	"errors"

	"github.com/PuerkitoBio/goquery"
	"github.com/goccy/go-json"
)

var QUESTION_URL_PREFIX = "https://cdn3.digialm.com"
var OPTIONS = [4]string{"A", "B", "C", "D"}

const QUESTION_COUNT = 90

func load_answer_key(filedata string) (AnswerKeyType, error) {
	var data AnswerKeyType
	HTML_DATA, _ := goquery.NewDocumentFromReader(strings.NewReader(filedata))
	DATA := HTML_DATA.Find("#ctl00_LoginContent_grAnswerKey>tbody>tr")
	if DATA.Length()-1 != QUESTION_COUNT {
		return data, errors.New("invalid_answer_key")
	}
	DATA.Each(func(i int, tr *goquery.Selection) {
		if i == 0 {
			// Skip the first row as it is the header
			return
		}
		data.QuestionID[i-1] = tr.Find("td>span").Eq(1).Text()
		data.CorrectAnswerID[i-1] = tr.Find("td>span").Eq(2).Text()
	})
	return data, nil
}

func findIndex(arr []string, val string) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}
func findIndex_90(arr [90]string, val string) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}

func normalize_double_dash(str string) string {
	if str == " -- " {
		return "--"
	}
	return str
}

var TEST_MODE = kry.DEV_ENV && len(os.Args) > 1 && os.Args[1] == "test"

func GetData(answer_key_data, response_sheet_data string) (MainDataType, error) {
	if TEST_MODE {
		fl_answer, _ := os.ReadFile("../tests/data/answer_key_html")
		answer_key_data = string(fl_answer)
		fl_resp, _ := os.ReadFile("../tests/data/question_paper_html")
		response_sheet_data = string(fl_resp)
	}
	ANSWER_KEY, err := load_answer_key(answer_key_data)
	var data MainDataType
	if err != nil {
		return data, errors.New("invalid_answer_key")
	}
	HTML_DATA, _ := goquery.NewDocumentFromReader(strings.NewReader(response_sheet_data))
	ANSWERS := HTML_DATA.Find("td.rw table.menu-tbl")
	QUESTIONS := HTML_DATA.Find("td.rw table.questionRowTbl")
	if !(ANSWERS.Length() == QUESTION_COUNT && QUESTIONS.Length() == QUESTION_COUNT) {
		return data, errors.New("invalid_response_sheet")
	}
	ANSWERS.Each(func(i int, tbl *goquery.Selection) {
		// Question Type
		typ := tbl.Find("tr").Eq(0).Find("td").Eq(1).Text()
		// Question ID
		q_id := tbl.Find("tr").Eq(1).Find("td").Eq(1).Text()

		// Option Number
		awns_number := ""
		// Option ID
		awns_id := "--"
		given_ans := "--"
		{
			// Setting the QuestionID and CorrectAnswerID in main data as well
			data.QuestionID[i] = q_id
			data.CorrectAnswerID[i] = ANSWER_KEY.CorrectAnswerID[findIndex_90(ANSWER_KEY.QuestionID, q_id)]
		}
		if typ == "MCQ" {
			awns_number = normalize_double_dash(tbl.Find("td").Eq(15).Text())
			awns_ids := []string{
				tbl.Find("tr").Eq(2).Find("td").Eq(1).Text(),
				tbl.Find("tr").Eq(3).Find("td").Eq(1).Text(),
				tbl.Find("tr").Eq(4).Find("td").Eq(1).Text(),
				tbl.Find("tr").Eq(5).Find("td").Eq(1).Text(),
			}
			// OptionID of the 4 options
			data.CorrectAnswer[i] = OPTIONS[findIndex(awns_ids, data.CorrectAnswerID[i])]
			if awns_number != "--" {
				answ_numb, _ := strconv.Atoi(awns_number)
				awns_id = awns_ids[answ_numb-1]
				given_ans = OPTIONS[answ_numb-1]
			}
			IMG_IDS := []string{
				QUESTIONS.Eq(i).Find("img").Eq(0).AttrOr("src", ""),
				QUESTIONS.Eq(i).Find("img").Eq(1).AttrOr("src", ""),
				QUESTIONS.Eq(i).Find("img").Eq(2).AttrOr("src", ""),
				QUESTIONS.Eq(i).Find("img").Eq(3).AttrOr("src", ""),
				QUESTIONS.Eq(i).Find("img").Eq(4).AttrOr("src", ""),
			}
			data.QuestionIMG[i] = QUESTION_URL_PREFIX + IMG_IDS[0]
			data.Option1IMG[i] = QUESTION_URL_PREFIX + IMG_IDS[1]
			data.Option2IMG[i] = QUESTION_URL_PREFIX + IMG_IDS[2]
			data.Option3IMG[i] = QUESTION_URL_PREFIX + IMG_IDS[3]
			data.Option4IMG[i] = QUESTION_URL_PREFIX + IMG_IDS[4]
		} else {
			id12 := normalize_double_dash(QUESTIONS.Eq(i).Find("tr td.rw td").Eq(5).Text())
			if id12 != "--" {
				awns_id = id12
				given_ans = id12
			}
		}
		data.GivenAnswerID[i] = awns_id
		data.GivenAnswer[i] = given_ans
		if typ == "MCQ" {
			data.Type[i] = typ
		} else {
			data.Type[i] = "NUM"
		}
		if typ == "SA" { // Numerical Questions
			data.CorrectAnswer[i] = data.CorrectAnswerID[i]
			data.QuestionIMG[i] = QUESTION_URL_PREFIX + QUESTIONS.Eq(i).Find("img").AttrOr("src", "")
		}
	})

	if TEST_MODE {
		json_data, _ := json.MarshalIndent(data, "", "  ")

		// Clearing the output directory
		os.RemoveAll("./jee/process_data/out")
		os.MkdirAll("./jee/process_data/out", 0755)

		os.WriteFile("./jee/process_data/out/data.json", json_data, 0644)
	}
	return data, nil
}

func GetResult(data *MainDataType) ResultDataType {
	var dt ResultDataType
	for i := 0; i < 3; i++ {
		// Initializing empty arrays as slices as normal empty arrays tend to be serializwed as null
		dt.Subjects[i].Correct = make([]string, 0)
		dt.Subjects[i].Incorrect = make([]string, 0)
		dt.Subjects[i].Unattempted = make([]string, 0)
	}
	for i := 0; i < QUESTION_COUNT; i++ {
		if data.GivenAnswer[i] == "--" {
			dt.UnattemptedCount++
			dt.Subjects[i/30].Unattempted = append(dt.Subjects[i/30].Unattempted, strconv.Itoa(i+1))
		} else if data.GivenAnswer[i] == data.CorrectAnswer[i] {
			dt.CorrectCount++
			dt.Subjects[i/30].Correct = append(dt.Subjects[i/30].Correct, strconv.Itoa(i+1))
		} else {
			dt.IncorrectCount++
			dt.Subjects[i/30].Incorrect = append(dt.Subjects[i/30].Incorrect, strconv.Itoa(i+1))
		}
	}
	for i := 0; i < 3; i++ {
		dt.Subjects[i].Score = len(dt.Subjects[i].Correct)*4 - len(dt.Subjects[i].Incorrect)
	}
	dt.Score = 4*dt.CorrectCount - 1*dt.IncorrectCount
	if TEST_MODE {
		json_data, _ := json.MarshalIndent(dt, "", "  ")
		os.WriteFile("./jee/process_data/out/result.json", json_data, 0644)
	}
	return dt
}
func GetMetaData(response_sheet_data string, answer_key_data string) (MetaDataType, error) {
	var data MetaDataType
	if TEST_MODE {
		fl_resp, _ := os.ReadFile("../tests/data/question_paper_html")
		response_sheet_data = string(fl_resp)
		fl_answ, _ := os.ReadFile("../tests/data/answer_key_html")
		answer_key_data = string(fl_answ)
	}
	HTML_DATA, _ := goquery.NewDocumentFromReader(strings.NewReader(response_sheet_data))
	ANSWER_KEY_HTML, _ := goquery.NewDocumentFromReader(strings.NewReader(answer_key_data))
	HTML_QUERY := HTML_DATA.Find(".main-info-pnl table tbody tr")
	if HTML_QUERY.Length() == 0 {
		return data, errors.New("metadata_not_found")
	}

	data.ApplicationNumber = HTML_QUERY.Eq(0).Find("td").Eq(1).Text()
	data.Name = HTML_QUERY.Eq(1).Find("td").Eq(1).Text()
	data.RollNumber = HTML_QUERY.Eq(2).Find("td").Eq(1).Text()
	data.DateOfExam = HTML_QUERY.Eq(3).Find("td").Eq(1).Text()
	data.TimeOfExam = HTML_QUERY.Eq(4).Find("td").Eq(1).Text()
	data.DOB = ANSWER_KEY_HTML.Find("#ctl00_LoginContent_lblDob").Text()
	data.DOB = strings.ReplaceAll(data.DOB, "-", "/")
	if TEST_MODE {
		json_data, _ := json.MarshalIndent(data, "", "  ")
		os.WriteFile("./jee/process_data/out/meta_data.json", json_data, 0644)
	}
	return data, nil
}
