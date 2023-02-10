<script lang="ts">
  import type { ResponseDataType } from './data_type';
  import { clsx } from '@tools/clsx';
  export let datt: ResponseDataType;

  const { data, result } = datt;

  let mode: 'result' | 'questions' = 'result';
</script>

<div class="mt-4">
  <div class="h-[80vh] w-full rounded-xl border-2 border-gray-700">
    <div
      class="rounded-tl-xl rounded-tr-xl border-b-2 border-b-gray-400 bg-zinc-100 text-xl font-bold"
    >
      <button
        on:click={() => (mode = 'result')}
        class={clsx(
          'cursor-button rounded-tl-xl border-r-2 border-r-neutral-500 px-4 text-red-500 transition',
          mode === 'result' ? 'bg-yellow-100' : ''
        )}>Result</button
      ><button
        on:click={() => (mode = 'questions')}
        class={clsx(
          'cursor-button border-r-2 border-r-neutral-500 px-4 text-blue-600 transition',
          mode === 'questions' ? 'bg-yellow-100' : ''
        )}>Questions</button
      >
    </div>
    <div class="overflow-scroll p-1 text-base">
      {#if mode === 'result'}
        <h2 class="text-xl font-bold">Summary</h2>
        <div class="flex justify-start">
          <div class="mr-4 flex flex-col">
            <div class="text-neutral-800">Unattempted Questions</div>
            <div class="text-neutral-800">Attempted Questions</div>
            <div class="text-neutral-800">Correct Answers</div>
            <div class="text-neutral-800">Incorrect Answers</div>
            <div class="font-bold text-green-900">Total Score</div>
          </div>
          <div class="flex flex-col">
            <div class="text-zinc-800">{result.unattempted.length}</div>
            <div class="text-zinc-800">{result.correct.length + result.incorrect.length}</div>
            <div class="text-zinc-800">{result.correct.length}</div>
            <div class="text-zinc-800">{result.incorrect.length}</div>
            <div class="font-bold text-blue-900">{result.score}</div>
          </div>
        </div>
        <h2 class="mt-2 text-xl font-bold">Detailed Report</h2>
        <div class="list-disc">
          <li>
            <span class="font-semibold text-green-700">Correct Questions</span> :-
            <span class="text-slate-700">{result.correct.join(', ')}</span>
          </li>
          <li>
            <span class="font-semibold text-red-700">Incorrect Questions</span> :-
            <span class="text-slate-700">{result.incorrect.join(', ')}</span>
          </li>
          <li>
            <span class="font-semibold text-blue-700">Unattempted Questions</span> :-
            <span class="text-slate-700">{result.unattempted.join(', ')}</span>
          </li>
        </div>
      {/if}
    </div>
  </div>
</div>
