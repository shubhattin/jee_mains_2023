from . import ResultDataType, MainDataType
from functools import reduce


def get_result(data: MainDataType) -> ResultDataType:
    dt = ResultDataType()
    for i in range(len(data.QuestionID)):
        if data.GivenAnswer[i] == "--":
            dt.subjects[i // 30].unattempted.append(str(i + 1))
        elif data.GivenAnswer[i] == data.CorrectAnswer[i]:
            dt.subjects[i // 30].correct.append(str(i + 1))
        else:
            dt.subjects[i // 30].incorrect.append(str(i + 1))
    for i in range(3):
        dt.subjects[i].score = (
            len(dt.subjects[i].correct) * 4 - len(dt.subjects[i].incorrect) * 1
        )
    dt.unattempted_count = reduce(
        lambda prev, next: prev + len(next.unattempted), dt.subjects, 0
    )
    dt.correct_count = reduce(
        lambda prev, next: prev + len(next.correct), dt.subjects, 0
    )
    dt.incorrect_count = reduce(
        lambda prev, next: prev + len(next.incorrect), dt.subjects, 0
    )
    dt.score = 4 * dt.correct_count - 1 * dt.incorrect_count
    return dt
