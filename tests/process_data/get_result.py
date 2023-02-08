from . import ResultDataType, MainDataType
import jinja2


def get_result(data: MainDataType) -> ResultDataType:
    dt = ResultDataType()
    for i in range(len(data.QuestionID)):
        if data.GivenAnswer[i] == "--":
            dt.unattempted.append(str(i + 1))
        elif data.GivenAnswer[i] == data.CorrectAnswer[i]:
            dt.correct.append(str(i + 1))
        else:
            dt.incorrect.append(str(i + 1))
    dt.score = len(dt.correct) * 4 - len(dt.incorrect) * 1
    return dt


TEMPLATE_FOLDER: str = "./data"


def set_template_folder(folder: str):
    global TEMPLATE_FOLDER
    TEMPLATE_FOLDER = folder


def render_jinja_temp(fl: str, **o):
    return (
        jinja2.Environment(loader=jinja2.FileSystemLoader(TEMPLATE_FOLDER))
        .get_template(fl)
        .render(**o)
    )


def get_result_markdown(data: MainDataType, sc: ResultDataType) -> str:
    OPTIONS = {
        "correct": len(sc.correct),
        "incorrect": len(sc.incorrect),
        "unattempted": len(sc.unattempted),
        "attempted": len(sc.correct) + len(sc.incorrect),
        "score": sc.score,
        "data": data,
        "correct_ques": sc.correct,
        "incorrect_ques": sc.incorrect,
        "unattempted_ques": sc.unattempted,
    }
    return render_jinja_temp("result_temp.md.j2", **OPTIONS)


def get_questions_markdown(data: MainDataType, sc: ResultDataType) -> str:
    return render_jinja_temp("questions_temp.md.j2", data=data)
