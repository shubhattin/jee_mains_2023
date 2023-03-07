from pydantic import BaseModel
from typing import List


class AnswerKeyType(BaseModel):
    QuestionID: List[str] = [None] * 90
    CorrectAnswerID: List[str] = [None] * 90


class MainDataType(BaseModel):
    GivenAnswer: List[str] = [None] * 90
    CorrectAnswer: List[str] = [None] * 90
    Type: List[str] = [None] * 90
    QuestionID: List[str] = [None] * 90
    CorrectAnswerID: List[str] = [None] * 90
    GivenAnswerID: List[str] = [None] * 90
    QuestionIMG: List[str] = [None] * 90
    Option1IMG: List[str] = [None] * 90
    Option2IMG: List[str] = [None] * 90
    Option3IMG: List[str] = [None] * 90
    Option4IMG: List[str] = [None] * 90


class ResultMetrics(BaseModel):
    unattempted: List[str] = []
    correct: List[str] = []
    incorrect: List[str] = []
    score: int = 0


class ResultDataType(BaseModel):
    score: int = 0
    unattempted_count: int = 0
    correct_count: int = 0
    incorrect_count: int = 0
    subjects: "List[ResultMetrics]" = [
        ResultMetrics(),  # Physics
        ResultMetrics(),  # Chemistry
        ResultMetrics(),  # Mathematics
    ]


class InfoDataType(BaseModel):
    ApplicationNumber: str = None
    Name: str = None
    RollNumber: str = None
    TestDate: str = None
    TestTime: str = None
