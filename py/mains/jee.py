from fastapi import APIRouter
from .process_data.load_data import load_data
from .process_data.get_result import get_result
from kry.datt import Drive, deta_val, Base, PROD_ENV

router = APIRouter(prefix="/api")


@router.post("/get_result")
async def result():
    data = load_data(
        Drive("temp_data").get("answer_key.csv").read().decode("utf-8"),
        Drive("temp_data").get("question_paper_html").read().decode("utf-8"),
    )
    # Doing work with temporary data to demonstrate the use
    result = get_result(data)

    if PROD_ENV:
        # Tracking the number of times the result is viewed
        result_count = deta_val("result_view_count", "counts")
        result_count += 1
        Base("counts").put(result_count, "result_view_count")

    return {"result": result.__dict__, "data": data.__dict__}


@router.post("/page_view_count")
def page_view_count():
    # This will be called every time the page is loaded only in production

    # Tracking the number of times the page is viewed
    page_view_count = deta_val("page_view_count", "counts")
    page_view_count += 1
    Base("counts").put(page_view_count, "page_view_count")
    return {
        "page_view_count": page_view_count,
        "result_view_count": deta_val("result_view_count", "counts"),
    }
