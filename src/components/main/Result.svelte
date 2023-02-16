<script lang="ts">
  import { lekhAH, datt, application_number, sample_result_status } from '@state/main';
  import { scale } from 'svelte/transition';
  import { clsx } from '@tools/clsx';
  import QuestionsTab from './Result/QuestionsTab.svelte';
  import ResultTab from './Result/ResultTab.svelte';

  $: data = $datt.data;
  $: result = $datt.result;

  let mode: 'result' | 'questions' = 'result';
</script>

<div class="mt-4" in:scale>
  {#if $sample_result_status}
    <div class="mb-2 font-bold text-amber-700">{$lekhAH.result.sample_result_msg}</div>
  {:else}
    <div class="mb-2">
      <span class="font-bold text-zinc-700">{$lekhAH.home.appl_numb}</span> :
      <span class="text-blue-700">{$application_number}</span>
    </div>
  {/if}
  <div
    class="h-[80vh] w-full overflow-x-hidden overflow-y-scroll rounded-xl border-2 border-gray-700"
  >
    <div
      class="sticky top-0 w-screen rounded-tl-xl rounded-tr-xl border-b-2 border-b-gray-400 bg-zinc-100 text-xl font-bold"
    >
      <button
        on:click={() => (mode = 'result')}
        class={clsx(
          'cursor-button rounded-tl-lg border-r-2 border-r-neutral-500 px-4 text-red-500 transition active:text-black',
          mode === 'result' ? 'border-2 border-black bg-yellow-100' : ''
        )}
        ><span class="text-black" class:hidden={mode !== 'result'}> • </span>{$lekhAH.result
          .result_tab.title}</button
      ><button
        on:click={() => (mode = 'questions')}
        class={clsx(
          'cursor-button border-r-2 border-r-neutral-500 px-4 text-blue-600 transition active:text-black',
          mode === 'questions' ? 'border-2 border-black bg-yellow-100' : ''
        )}
        ><span class="text-black" class:hidden={mode !== 'questions'}> • </span>{$lekhAH.result
          .questions_tab.title}</button
      >
    </div>
    <div class="p-1 text-base">
      {#if mode === 'result'}
        <ResultTab {data} {result} />
      {:else if mode === 'questions'}
        <QuestionsTab {data} />
      {/if}
    </div>
  </div>
</div>
