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
  result: resultTyoe & {
    subjects: resultTyoe[];
  };
}

export const API_URL: string = import.meta.env.VITE_API_URL || '';
