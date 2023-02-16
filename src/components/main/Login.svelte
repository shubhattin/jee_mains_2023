<script lang="ts">
  import type { ResponseDataType } from '@components/main/type';
  import { lekhAH, mode as mainMode, datt } from '@state/main';
  import { fetch_post } from '@tools/fetch';
  import { clsx } from '@tools/clsx';
  import Spinner from '@components/Spinner.svelte';

  $: lekh = $lekhAH.home;
  let dob: string;
  let appl_numb: number;
  let get_result_loading_status = false;
  let get_sample_result_loading_status = false;
  let submit_data_loading_status = false;
  let dob_wrong_error = false;
  let appl_numb_not_found_err_msg = false;
  let responseData = '';
  let answerKey = '';

  $: {
    if (dob_wrong_error) setTimeout(() => (dob_wrong_error = false), 700);
  }

  const fetch_sample_data_result = async () => {
    get_sample_result_loading_status = true;
    const req = await fetch_post('/api/get_sample_result');
    if (!req.ok) {
      get_sample_result_loading_status = false;
      return;
    }
    const resp: ResponseDataType = await req.json();
    $datt = resp;
    $mainMode = 'result';
  };
  const fetch_data_result = async () => {
    if (dob === '' || !appl_numb || appl_numb.toString().length != 12) return;
    get_result_loading_status = true;
    const req = await fetch_post('/api/get_result', {
      json: {
        ApplicationNumber: appl_numb.toString(),
        DateOfBirth: dob.split('-').reverse().join('/')
      }
    });
    if (req.status === 200) {
      const resp: ResponseDataType = await req.json();
      $datt = resp;
      $mainMode = 'result';
    } else if (req.status === 403) {
      const reason = req.headers.get('x-error');
      if (reason! === 'appl_numb_not_found') {
        get_result_loading_status = false;
        appl_numb_not_found_err_msg = true;
        mode = 'extra';
      } else if (reason === 'dob_did_not_match') {
        dob = '';
        dob_wrong_error = true;
        get_result_loading_status = false;
      }
    }
  };
  const submit_data = async () => {
    if (responseData === '' || answerKey === '') return;
    submit_data_loading_status = true;
    const req = await fetch_post('/api/submit_result_data', {
      json: {
        ResponsePageData: responseData,
        AnswerKeyData: answerKey
      }
    });
    if (req.status === 200) {
      const resp: ResponseDataType = await req.json();
      $datt = resp;
      $mainMode = 'result';
    } else if (req.status === 403) {
      submit_data_loading_status = false;
    }
  };
  let mode: 'login' | 'extra' = 'extra';
  // let mode: 'login' | 'extra' = 'login';
</script>

<div class="my-4">
  {#if mode === 'login'}
    <div>
      <div class="mb-2">
        <input
          type="number"
          bind:value={appl_numb}
          placeholder={lekh.appl_numb}
          class={clsx(
            'block w-44 rounded-md border-2 px-1 outline-none transition-all duration-200 focus:ring-2',
            'border-blue-800 ring-green-500 placeholder:text-zinc-400'
          )}
        />
        {#if appl_numb && appl_numb.toString().length != 12}
          <div class="mt-1 text-sm text-red-500">{lekh.verify.appl_numb_verify}</div>
        {/if}
      </div>
      <div class="mb-2">
        <input
          type="date"
          bind:value={dob}
          placeholder={lekh.DOB}
          class={clsx(
            'mb-1 block w-44 rounded-md border-2 px-1 outline-none transition-all duration-200',
            'border-blue-800 placeholder:text-zinc-400 focus:ring-2 focus:ring-green-500',
            dob_wrong_error ? 'border-4 border-red-600' : ''
          )}
        />
        {#if dob == ''}
          <div class="mt-1 text-sm text-red-500">{lekh.verify.dob_verify}</div>
        {/if}
      </div>
      <div>
        <button
          on:click={fetch_data_result}
          class={clsx(
            'cursor-button rounded-xl border-2 border-red-700 px-1 text-base text-green-700',
            'ring-orange-400 active:border-black active:bg-zinc-100 active:text-black active:ring-2'
          )}
        >
          {#if get_result_loading_status}
            <Spinner className="text-black w-5 h-5 m-1" />
          {/if}
          {lekh.get_result}
        </button>
        <button
          class="cursor-button ml-16 rounded-xl border-2 border-blue-700 px-1 text-sm text-zinc-500"
          on:click={() => (mode = 'extra')}>Get result using data</button
        >
      </div>
    </div>
    <button
      on:click={fetch_sample_data_result}
      class={clsx(
        'cursor-button mt-2 block rounded-xl border-2 border-purple-600 bg-purple-500 px-1 text-base text-white',
        'active:border-amber-700 active:bg-purple-300 active:text-black',
        'ring-purple-300 active:ring-2'
      )}
    >
      {#if get_sample_result_loading_status}
        <Spinner className="text-white w-5 h-5 m-1" />
      {/if}
      {lekh.get_sample_result}
    </button>
  {:else if mode === 'extra'}
    <div class="list-disc text-slate-600">
      {#if appl_numb_not_found_err_msg}
        <div class="text-red-700">
          ⚠️ As you are viewing your results for first time you need to provide this Information
        </div>
      {/if}
      <div>
        <div class="text-lg font-bold text-amber-800">Response Page Data</div>
        <textarea
          spellcheck="false"
          class="resize-none rounded-lg border-2 border-blue-700 p-1 text-sm text-black outline-none transition focus:ring-2 focus:ring-blue-400"
          rows="3"
          bind:value={responseData}
          placeholder="Enter URL or HTML of the Response Page"
          cols="40"
        />
      </div>
      <div class="mt-1">
        <div class="text-lg font-bold text-yellow-800">Answer Key Data</div>
        <textarea
          spellcheck="false"
          class="resize-none rounded-lg border-2 border-purple-700 p-1 text-sm text-black outline-none transition focus:ring-2 focus:ring-purple-400"
          rows="5"
          bind:value={answerKey}
          placeholder="Enter HTML of the Answer Key Page"
          cols="40"
        />
      </div>
      <button
        on:click={submit_data}
        class="mt-1 rounded-lg border-2 border-black bg-green-300 px-1 text-lg font-bold text-red-600 ring-zinc-500 transition-all active:text-black active:ring-2"
      >
        {#if submit_data_loading_status}
          <Spinner className="text-red-600 w-5 h-5 mb-1" />
        {/if}
        Submit
      </button>
    </div>
  {/if}
</div>
