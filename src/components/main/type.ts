interface resultTyoe {
  unattempted: string[];
  correct: string[];
  incorrect: string[];
  score: number;
}

export interface ResponseDataType {
  data: {
    GivenAnswer: string[];
    CorrectAnswer: string[];
    Type: string[];
    QuestionID: string[];
    CorrectAnswerID: string[];
    GivenAnswerID: string[];
    QuestionIMG: string[];
    Option1IMG: string[];
    Option2IMG: string[];
    Option3IMG: string[];
    Option4IMG: string[];
  };
  result: {
    score: number;
    unattempted_count: number;
    correct_count: number;
    incorrect_count: number;
    subjects: resultTyoe[];
  };
}

export const API_URL: string = import.meta.env.VITE_API_URL || '';
