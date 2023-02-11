import csv
from pyquery import PyQuery as pq
from . import AnswerKeyType, MainDataType
from typing import List

OPTIONS = ["A", "B", "C", "D"]
QUESTION_URL_PREFIX = "https://cdn3.digialm.com"


def load_answer_key(filedata: str) -> AnswerKeyType:
    data = AnswerKeyType()
    reader = csv.DictReader(filedata.splitlines())
    for row in reader:
        data.QuestionID.append(row["QuestionID"])
        data.CorrectAnswerID.append(row["Correct Option(s)/ Answers"])
    return data


def load_data(anwer_key_data: str, response_sheet_data: str) -> MainDataType:
    data = MainDataType()
    ANSWER_KEY = load_answer_key(anwer_key_data)
    HTML_DATA = pq(response_sheet_data)

    # This contains the selected option and OptionIDs
    ANSWERS = HTML_DATA("td.rw table.menu-tbl")

    # This contains the question text, images and answer in case of Numerical Questions
    QUESTIONS = HTML_DATA("td.rw table.questionRowTbl")

    i = 0
    for table in ANSWERS:
        tbl = pq(table)

        # Question Type
        typ = tbl("tr").eq(0)("td").eq(1).text()
        # Question ID
        q_id = tbl("tr").eq(1)("td").eq(1).text()
        # Option Number
        awns_number: str = ""
        # Option ID
        awns_id: str = "--"
        given_ans: str = "--"
        # ques_index = data[data["QuestionID"] == q_id].index[0]

        if True:
            # Setting the QuestionID and CorrectAnswerID in main data as well
            data.QuestionID[i] = q_id
            data.CorrectAnswerID[i] = ANSWER_KEY.CorrectAnswerID[
                ANSWER_KEY.QuestionID.index(q_id)
            ]

        if typ == "MCQ":
            awns_number = tbl("td").eq(15).text()
            awns_ids: List[str] = [
                tbl("tr").eq(x + 2)("td").eq(1).text() for x in range(4)
            ]
            # QuestionID followd by OptionID of the 4 options
            data.CorrectAnswer[i] = OPTIONS[
                awns_ids.index(data.CorrectAnswerID[i])
            ]  # doing this as the first is quwstion id
            if awns_number != "--":
                awns_id = awns_ids[int(awns_number) - 1]
                given_ans = OPTIONS[int(awns_number) - 1]
            IMG_IDS: List[str] = [
                x.attrib["src"] for x in QUESTIONS.eq(i).eq(0).eq(0)("img")
            ]
            data.QuestionIMG[i] = QUESTION_URL_PREFIX + IMG_IDS[0]
            data.Option1IMG[i] = QUESTION_URL_PREFIX + IMG_IDS[1]
            data.Option2IMG[i] = QUESTION_URL_PREFIX + IMG_IDS[2]
            data.Option3IMG[i] = QUESTION_URL_PREFIX + IMG_IDS[3]
            data.Option4IMG[i] = QUESTION_URL_PREFIX + IMG_IDS[4]
        else:
            id12 = QUESTIONS.eq(i).eq(0).text().split("\n")[-1]
            if id12 != "--":
                awns_id = id12
                given_ans = id12
        data.GivenAnswerID[i] = awns_id
        data.GivenAnswer[i] = given_ans
        data.Type[i] = typ if typ == "MCQ" else "NUM"
        if typ == "SA":  # Numerical Questions
            data.CorrectAnswer[i] = data.CorrectAnswerID[i]
            data.QuestionIMG[i] = (
                QUESTION_URL_PREFIX
                + QUESTIONS.eq(i).eq(0).eq(0)("img")[0].attrib["src"]
            )
        i += 1

    return data
