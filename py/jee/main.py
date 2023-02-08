from fastapi import APIRouter
from fastapi.responses import HTMLResponse
from .process_data.load_data import load_data
from .process_data.get_result import (
    get_result,
    get_questions_markdown,
    get_result_markdown,
)
import markdown

router = APIRouter(prefix="")


def markdown_to_html(vl: str) -> str:
    opt = ["codehilite", "fenced_code", "tables"]
    return markdown.markdown(vl, extensions=opt)


@router.get("/")
async def result():
    data = load_data(
        "./data/temp_data/answer_key.csv", "./data/temp_data/question_paper_html"
    )
    result = get_result(data)
    html = markdown_to_html(get_result_markdown(data, result))
    return HTMLResponse(html)


@router.get("/questions")
async def questions():
    data = load_data(
        "./data/temp_data/answer_key.csv", "./data/temp_data/question_paper_html"
    )
    result = get_result(data)
    html = markdown_to_html(get_questions_markdown(data, result))
    return HTMLResponse(html)
