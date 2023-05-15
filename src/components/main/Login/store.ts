import { writable } from 'svelte/store';
import type { ResponseDataType } from '@components/main/type';

export const mode = writable<'login' | 'submit_data'>('login');
export const appl_numb_not_found_err_msg = writable(false);

export const normalize_data = (data: ResponseDataType) => {
  // this function is used to tackle the problem of empty arrays being serialized into null
  // this might be fixed in go(/api) folder, but is here to suuport old cached results
  for (let i = 0; i < 3; i++) {
    const sbjct = data.result.subjects[i];
    if (!sbjct.correct || sbjct.correct === null || sbjct.correct === undefined) sbjct.correct = [];
    if (!sbjct.incorrect || sbjct.incorrect === null || sbjct.incorrect === undefined)
      sbjct.incorrect = [];
    if (!sbjct.unattempted || sbjct.unattempted === null || sbjct.unattempted === undefined)
      sbjct.unattempted = [];
  }
  return data;
};
