<script lang="ts">
  import type { ResponseDataType } from '@components/main/type';
  import { API_URL } from '@components/main/type';
  import {
    lekhAH,
    mode as mainMode,
    datt,
    viewCountData,
    application_number as appl_numb
  } from '@state/main';
  import { fetch_post } from '@tools/fetch';
  import { clsx } from '@tools/clsx';
  import Spinner from '@components/Spinner.svelte';
  import { slide } from 'svelte/transition';
  import { appl_numb_not_found_err_msg, normalize_data } from '@components/main/Login/store';

  $: lekh = $lekhAH.home;
  let sumbit_loading_status = false;
  let parse_error_status = false;
  let responseData = '';
  let answerKey = '';

  let empty_status = {
    responseData: false,
    answerKey: false
  };

  $: {
    if (empty_status.responseData) setTimeout(() => (empty_status.responseData = false), 700);
    if (empty_status.answerKey) setTimeout(() => (empty_status.answerKey = false), 700);
  }

  $: {
    if (parse_error_status) setTimeout(() => (parse_error_status = false), 4800);
  }

  const submit_data = async () => {
    if (true) {
      let return_status = false;
      if (!responseData || responseData === '') {
        empty_status.responseData = true;
        return_status = true;
      }
      if (!answerKey || answerKey === '') {
        empty_status.answerKey = true;
        return_status = true;
      }
      if (return_status) return;
    }
    sumbit_loading_status = true;
    const req = await fetch_post(API_URL + '/api/submit_result_data', {
      json: {
        ResponsePageData: responseData,
        AnswerKeyData: answerKey
      }
    });
    if (req.status === 200) {
      const resp: ResponseDataType & {
        key: string;
      } = await req.json();
      $appl_numb = resp.key;
      $datt = normalize_data(resp);
      $mainMode = 'result';
      $viewCountData[1]++;
    } else if (req.status === 403) {
      sumbit_loading_status = false;
      parse_error_status = true;
    }
  };
</script>

<div in:slide>
  {#if $appl_numb_not_found_err_msg}
    <div class="text-red-700">
      ⚠️ {lekh.data_page.first_time_info}
    </div>
  {/if}
  <div>
    <div class="text-lg font-bold text-amber-800">{lekh.data_page.response_page_data}</div>
    <textarea
      spellcheck="false"
      class={clsx(
        'max-w-full resize-none rounded-lg border-2 border-blue-700 p-1 text-sm text-black outline-none transition-all focus:ring-2 focus:ring-blue-400',
        empty_status.responseData ? 'border-4 border-red-400' : ''
      )}
      rows="3"
      bind:value={responseData}
      placeholder={lekh.data_page.response_page_data_msg}
      cols="60"
    />
  </div>
  <div class="mt-1">
    <div class="text-lg font-bold text-yellow-800">{lekh.data_page.answer_key_data}</div>
    <textarea
      spellcheck="false"
      class={clsx(
        'max-w-full resize-none rounded-lg border-2 border-purple-700 p-1 text-sm text-black outline-none transition-all focus:ring-2 focus:ring-purple-400',
        empty_status.answerKey ? 'border-4 border-red-400' : ''
      )}
      rows="4"
      bind:value={answerKey}
      placeholder={lekh.data_page.answer_key_data_msg}
      cols="60"
    />
  </div>
  {#if parse_error_status}
    <div class="font-bold text-red-700">{lekh.data_page.error_msg}</div>
  {/if}
  <button
    on:click={submit_data}
    class="mt-1 rounded-lg border-2 border-black bg-green-300 px-1 text-lg font-bold text-red-600 ring-zinc-500 transition-all active:text-black active:ring-2"
  >
    {#if sumbit_loading_status}
      <Spinner className="text-red-600 w-5 h-5 mb-1" />
    {/if}
    {lekh.data_page.submit_btn}
  </button>
  <div class="list-dis mt-8 text-sm text-zinc-600">
    <li class="font-bold text-zinc-700">
      {lekh.data_page.notes.get_html_page} :
      <div class="ml-4 list-disc">
        <li>
          <span class="text-blue-700">{lekh.data_page.notes.on_mobile} : </span>
          {@html lekh.data_page.notes.on_mobile_msg}
        </li>
        <li>
          <span class="text-blue-700">{lekh.data_page.notes.on_computer} : </span>
          {@html lekh.data_page.notes.on_computer_msg}
        </li>
      </div>
    </li>
    {#each lekh.data_page.notes.other_notes as notes}
      <li>{@html notes}</li>
    {/each}
  </div>
</div>
