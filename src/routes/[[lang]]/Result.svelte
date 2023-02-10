<script lang="ts">
  import type { ResponseDataType } from './data_type';
  import { lekhAH } from './data_type';
  import { scale, fly } from 'svelte/transition';
  import { clsx } from '@tools/clsx';

  export let datt: ResponseDataType;

  const { data, result } = datt;
  let result_tab: typeof $lekhAH.result.result_tab,
    questions_tab: typeof $lekhAH.result.questions_tab;
  $: {
    result_tab = $lekhAH.result.result_tab;
    questions_tab = $lekhAH.result.questions_tab;
  }

  let mode: 'result' | 'questions' = 'result';
</script>

<div class="mt-4" in:scale>
  <div class="h-[80vh] w-full overflow-scroll rounded-xl border-2 border-gray-700">
    <div
      class="rounded-tl-xl rounded-tr-xl border-b-2 border-b-gray-400 bg-zinc-100 text-xl font-bold"
    >
      <button
        on:click={() => (mode = 'result')}
        class={clsx(
          'cursor-button rounded-tl-xl border-r-2 border-r-neutral-500 px-4 text-red-500 transition active:text-black',
          mode === 'result' ? 'bg-yellow-100' : ''
        )}
        ><span class="text-black" class:hidden={mode !== 'result'}>
          •
        </span>{result_tab.title}</button
      ><button
        on:click={() => (mode = 'questions')}
        class={clsx(
          'cursor-button border-r-2 border-r-neutral-500 px-4 text-blue-600 transition active:text-black',
          mode === 'questions' ? 'bg-yellow-100' : ''
        )}
        ><span class="text-black" class:hidden={mode !== 'questions'}>
          •
        </span>{questions_tab.title}</button
      >
    </div>
    <div class="p-1 text-base">
      {#if mode === 'result'}
        <div transition:fly>
          <h2 class="text-xl font-bold">{result_tab.summary}</h2>
          <div class="flex justify-start">
            <div class="mr-4 flex flex-col">
              <div class="text-neutral-800">{result_tab.attempted_ques}</div>
              <div class="text-green-800">{result_tab.correct_ques}</div>
              <div class="text-red-800">{result_tab.incorrect_ques}</div>
              <div class="text-blue-800">{result_tab.unattempted_ques}</div>
              <div class="font-bold text-green-900">{result_tab.total_score}</div>
            </div>
            <div class="flex flex-col">
              <div class="text-zinc-800">{result.correct.length + result.incorrect.length}</div>
              <div class="text-green-800">{result.correct.length}</div>
              <div class="text-red-800">{result.incorrect.length}</div>
              <div class="text-blue-800">{result.unattempted.length}</div>
              <div class="font-bold text-blue-900">{result.score}</div>
            </div>
          </div>
          <h2 class="mt-2 text-xl font-bold">{result_tab.detailed_report}</h2>
          <div class="list-disc">
            <li>
              <span class="font-semibold text-green-700">{result_tab.correct_ques}</span> :-
              <span class="text-slate-700">{result.correct.join(', ')}</span>
            </li>
            <li>
              <span class="font-semibold text-red-700">{result_tab.incorrect_ques}</span> :-
              <span class="text-slate-700">{result.incorrect.join(', ')}</span>
            </li>
            <li>
              <span class="font-semibold text-blue-700">{result_tab.unattempted_ques}</span> :-
              <span class="text-slate-700">{result.unattempted.join(', ')}</span>
            </li>
          </div>
          <table class="mt-4 border-collapse rounded-xl border-2 border-blue-800">
            <thead>
              <tr>
                <th class="border border-zinc-400 px-1" />
                <th class="border border-zinc-400 px-1">{result_tab.table.given_answer}</th>
                <th class="border border-zinc-400 px-2">{result_tab.table.correct_answer}</th>
                <th class="border border-zinc-400 px-2">{result_tab.table.question_id}</th>
                <th class="border border-zinc-400 px-2">{result_tab.table.correct_answer_id}</th>
                <th class="border border-zinc-400 px-2">{result_tab.table.given_answer_id}</th>
              </tr>
            </thead>
            <tbody>
              {#each { length: data.GivenAnswer.length } as _, i}
                <tr>
                  <td class="border border-zinc-400 px-1">{i + 1}</td>
                  <td class="border border-zinc-400 px-1">{data.GivenAnswer[i]}</td>
                  <td class="border border-zinc-400 px-2">{data.CorrectAnswer[i]}</td>
                  <td class="border border-zinc-400 px-2">{data.QuestionID[i]}</td>
                  <td class="border border-zinc-400 px-2">{data.CorrectAnswerID[i]}</td>
                  <td class="border border-zinc-400 px-2">{data.GivenAnswerID[i]}</td>
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
      {:else if mode === 'questions'}
        <div transition:fly class="list-decimal">
          {#each { length: data.GivenAnswer.length } as _, i}
            <li class="mb-4">
              <img class="inline-block" src={data.QuestionIMG[i]} alt={`Question ${i + 1}`} />
              <div class="mt-2 ml-4 list-disc">
                {#if data.Type[i] === 'MCQ'}
                  {#each ['A', 'B', 'C', 'D'] as opt, opt_num}
                    <li class="mb-1">
                      {opt}.
                      <img
                        class="inline-block"
                        src={[data.Option1IMG, data.Option2IMG, data.Option3IMG, data.Option4IMG][
                          opt_num
                        ][i]}
                        alt={`Question ${i + 1} Option ${opt}`}
                      />
                    </li>
                  {/each}
                {/if}
                <li>{questions_tab.correct_answer} : {data.CorrectAnswer[i]}</li>
                <li>{questions_tab.marked_answer} : {data.GivenAnswer[i]}</li>
              </div>
            </li>
            <div class="h-1 w-full bg-gray-400" />
          {/each}
        </div>
      {/if}
    </div>
  </div>
</div>
