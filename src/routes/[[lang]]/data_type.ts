import { writable } from 'svelte/store';
import type { PageData } from './$types';

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
    unattempted: string[];
    correct: string[];
    incorrect: string[];
    score: number;
  };
}

export const lekhAH = writable<PageData['lekh']>(null!);
export const viewCountData = writable<[number, number]>([0, 0]);
