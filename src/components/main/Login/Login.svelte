<script lang="ts">
  import type { ResponseDataType } from '@components/main/type';
  import { API_URL } from '@components/main/type';
  import {
    lekhAH,
    mode as mainMode,
    datt,
    viewCountData,
    application_number as appl_numb,
    sample_result_status
  } from '@state/main';
  import { fetch_post } from '@tools/fetch';
  import { clsx } from '@tools/clsx';
  import Spinner from '@components/Spinner.svelte';
  import { mode, appl_numb_not_found_err_msg, normalize_data } from './store';
  import { DateInput } from 'date-picker-svelte';

  $: lekh = $lekhAH.home;
  let result_loading_status = false;
  let sample_result_loading_status = false;
  let dob_wrong_error = false;

  let date: Date = null!;
  let dob: string = null!;

  $: {
    if (date) dob = date.toLocaleDateString();
  }

  let empty_status = {
    appl_numb: false,
    dob: false
  };

  $: {
    if (empty_status.appl_numb) setTimeout(() => (empty_status.appl_numb = false), 600);
    if (empty_status.dob) setTimeout(() => (empty_status.dob = false), 600);
  }

  $: {
    if (dob_wrong_error) setTimeout(() => (dob_wrong_error = false), 700);
  }

  const fetch_sample_data_result = async () => {
    sample_result_loading_status = true;
    const req = await fetch_post(API_URL + '/api/get_sample_result');
    if (!req.ok) {
      sample_result_loading_status = false;
      return;
    }
    const resp: ResponseDataType = await req.json();
    $sample_result_status = true;
    $datt = normalize_data(resp);
    $mainMode = 'result';
  };
  const fetch_data_result = async () => {
    if (true) {
      let return_status = false;
      if (!$appl_numb || $appl_numb.toString().length != 12) {
        empty_status.appl_numb = true;
        return_status = true;
      }
      if (!dob || dob === '') {
        empty_status.dob = true;
        return_status = true;
      }
      if (return_status) return;
    }
    result_loading_status = true;
    const req = await fetch_post(API_URL + '/api/get_result', {
      json: {
        ApplicationNumber: $appl_numb.toString(),
        DateOfBirth: dob
      }
    });
    if (req.status === 200) {
      const resp: ResponseDataType = await req.json();
      $datt = normalize_data(resp);
      $mainMode = 'result';
      $viewCountData[1]++;
    } else if (req.status === 403) {
      const reason: string = (await req.json()).detail.error;
      if (reason === 'appl_numb_not_found') {
        result_loading_status = false;
        $appl_numb_not_found_err_msg = true;
        $mode = 'submit_data';
      } else if (reason === 'dob_did_not_match') {
        date = null!;
        dob_wrong_error = true;
        result_loading_status = false;
      }
    }
  };
</script>

<div>
  <!-- ApplicationNumber Input -->
  <div class="mb-2">
    <input
      type="number"
      bind:value={$appl_numb}
      placeholder={lekh.appl_numb}
      class={clsx(
        'block w-44 rounded-md border-2 px-1 outline-none transition-all duration-200 focus:ring-2',
        'border-blue-800 ring-green-500 placeholder:text-zinc-400',
        empty_status.appl_numb ? 'border-4 border-red-400' : ''
      )}
    />
    {#if $appl_numb && $appl_numb.toString().length != 12}
      <div class="mt-1 text-sm text-red-500">{lekh.verify.appl_numb_verify}</div>
    {/if}
  </div>
  <!-- DOB Input -->
  <div class="mb-2">
    <DateInput
      bind:value={date}
      placeholder="DD/MM/YYYY"
      format="dd/MM/yyyy"
      closeOnSelection={true}
      class={clsx(
        'mb-1 inline-block rounded-md border-2 text-base outline-none transition-all duration-200',
        'border-blue-800 placeholder:text-zinc-400 focus:ring-2 focus:ring-green-500',
        dob_wrong_error ? 'border-4 border-red-600' : '',
        empty_status.dob ? 'border-4 border-red-400' : ''
      )}
    />
  </div>
  <!-- Submit Button -->
  <div>
    <button
      on:click={fetch_data_result}
      class={clsx(
        'cursor-button rounded-xl border-2 border-red-700 px-1 text-base text-green-700',
        'ring-orange-400 active:border-black active:bg-zinc-100 active:text-black active:ring-2'
      )}
    >
      {#if result_loading_status}
        <Spinner className="text-black w-5 h-5 m-1" />
      {/if}
      {lekh.get_result}
    </button>
    <button
      class="cursor-button ml-16 rounded-xl border-2 border-blue-700 px-1 text-sm text-zinc-500"
      on:click={() => ($mode = 'submit_data')}>{lekh.get_data_result}</button
    >
  </div>
</div>
<button
  on:click={fetch_sample_data_result}
  class={clsx(
    'cursor-button mt-8 block rounded-xl border-2 border-purple-600 bg-purple-500 px-1 text-base text-white',
    'active:border-amber-700 active:bg-purple-300 active:text-black',
    'ring-purple-300 active:ring-2'
  )}
>
  {#if sample_result_loading_status}
    <Spinner className="text-white w-5 h-5 m-1" />
  {/if}
  {lekh.get_sample_result}
</button>
