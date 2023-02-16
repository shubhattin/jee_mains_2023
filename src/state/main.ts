import { writable } from 'svelte/store';
import type { dattType } from '@langs/datt';
import type { ResponseDataType } from '@components/main/type';
export const lekhAH = writable<dattType['main']>(null!);
export const viewCountData = writable<[number, number]>([0, 0]);

/** Status of count request been sent */
export const counted = writable(false);

/** Mode of the display og page */
export const mode = writable<'main' | 'result'>('main');

/** Main Responnse Data */
export const datt = writable<ResponseDataType>(null!);
export const application_number = writable<string>(null!);
export const sample_result_status = writable(false);
