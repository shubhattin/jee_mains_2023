from fastapi import APIRouter
from .process_data.load_data import load_data
from .process_data.get_result import get_result

router = APIRouter(prefix="/api")


@router.post("/get_result")
async def result():
    data = load_data(
        "./data/temp_data/answer_key.csv", "./data/temp_data/question_paper_html"
    )
    result = get_result(data)
    return {"result": result.__dict__, "data": data.__dict__}
