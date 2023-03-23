import { writable } from 'svelte/store';

export const mode = writable<'login' | 'submit_data'>('login');
export const appl_numb_not_found_err_msg = writable(false);
