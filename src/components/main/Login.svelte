<script lang="ts">
  import type { ResponseDataType } from '@components/main/type';
  import { lekhAH, mode, datt } from '@state/main';
  import { fetch_post } from '@tools/fetch';
  import { clsx } from '@tools/clsx';

  $: lekh = $lekhAH.home;
  let dob: string;
  let appl_numb: number;

  const fetch_sample_data_result = async () => {
    const req = await fetch_post('/api/get_sample_result');
    if (!req.ok) return;
    const resp: ResponseDataType = await req.json();
    $datt = resp;
    $mode = 'result';
  };
</script>

<div class="my-4">
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
          'mb-1 block w-44 rounded-md border-2 px-1 outline-none transition-all duration-200 focus:ring-2',
          'border-blue-800 ring-green-500 placeholder:text-zinc-400'
        )}
      />
      {#if dob == ''}
        <div class="mt-1 text-sm text-red-500">{lekh.verify.dob_verify}</div>
      {/if}
    </div>
    <button
      class={clsx(
        'cursor-button rounded-xl border-2 border-red-700 px-1 text-base text-green-700',
        'ring-orange-400 active:border-black active:bg-zinc-100 active:text-black active:ring-2'
      )}
    >
      {lekh.get_result}
    </button>
  </div>
</div>
<button
  on:click={fetch_sample_data_result}
  class={clsx(
    'cursor-button mt-2 block rounded-xl border-2 border-purple-600 bg-purple-700 px-1 text-base text-white',
    'active:border-amber-700 active:bg-purple-300 active:text-black',
    'ring-purple-300 active:ring-2'
  )}
>
  {lekh.get_sample_result}
</button>
