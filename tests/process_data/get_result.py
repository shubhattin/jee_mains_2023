from . import ResultDataType, MainDataType


def get_result(data: MainDataType) -> ResultDataType:
    dt = ResultDataType()
    for i in range(len(data.QuestionID)):
        if data.GivenAnswer[i] == "--":
            dt.unattempted.append(str(i + 1))
            dt.subjects[i // 30].unattempted.append(str(i + 1))
        elif data.GivenAnswer[i] == data.CorrectAnswer[i]:
            dt.correct.append(str(i + 1))
            dt.subjects[i // 30].correct.append(str(i + 1))
        else:
            dt.incorrect.append(str(i + 1))
            dt.subjects[i // 30].incorrect.append(str(i + 1))
    dt.score = len(dt.correct) * 4 - len(dt.incorrect) * 1
    for i in range(3):
        dt.subjects[i].score = (
            len(dt.subjects[i].correct) * 4 - len(dt.subjects[i].incorrect) * 1
        )
    return dt
