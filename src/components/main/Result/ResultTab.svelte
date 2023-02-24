<script lang="ts">
  import { fly } from 'svelte/transition';
  import type { ResponseDataType } from '@components/main/type';
  import { lekhAH } from '@state/main';

  $: lekh = $lekhAH.result.result_tab;

  export let data: ResponseDataType['data'];
  export let result: ResponseDataType['result'];
</script>

<div transition:fly>
  <div>
    <h2 class="text-xl font-bold">{lekh.summary}</h2>
    <div class="flex justify-start">
      <div class="mr-4 flex flex-col">
        <div class="text-neutral-800">{lekh.attempted_ques}</div>
        <div class="text-green-800">{lekh.correct_ques}</div>
        <div class="text-red-800">{lekh.incorrect_ques}</div>
        <div class="text-blue-800">{lekh.unattempted_ques}</div>
        <div class="font-bold text-green-900">{lekh.total_score}</div>
      </div>
      <div class="flex flex-col">
        <div class="text-zinc-800">{result.correct_count + result.incorrect_count}</div>
        <div class="text-green-800">{result.correct_count}</div>
        <div class="text-red-800">{result.incorrect_count}</div>
        <div class="text-blue-800">{result.unattempted_count}</div>
        <div class="font-bold text-blue-900">{result.score}</div>
      </div>
    </div>
    <h3 class="text-lg font-bold text-amber-800 underline">{lekh.subject_wise.title}</h3>
    <div class="overflow-x-scroll text-center text-sm sm:text-base">
      <table class="border-collapse rounded-xl border-2 border-blue-800">
        <thead>
          <tr>
            <th class="border border-slate-900 bg-gray-100 px-1 text-sm"
              >{lekh.subject_wise.table.subject}</th
            >
            <th class="border border-slate-900 bg-gray-100 px-1 text-xs text-green-800 sm:text-sm"
              >{lekh.subject_wise.table.correct}</th
            >
            <th class="border border-slate-900 bg-gray-100 px-1 text-xs text-red-700 sm:text-sm"
              >{lekh.subject_wise.table.incorrect}</th
            >
            <th class="border border-slate-900 bg-gray-100 px-1 text-xs text-blue-900 sm:text-sm"
              >{lekh.subject_wise.table.unattempted}</th
            >
            <th class="border border-slate-900 bg-gray-100 px-1 text-xs text-amber-900 sm:text-sm"
              >{lekh.subject_wise.table.score}</th
            >
          </tr>
        </thead>
        <tbody>
          {#each { length: 3 } as _, i (i)}
            <tr>
              <td class="border border-zinc-400 px-1 text-violet-900"
                >{lekh.subject_wise.subjects[i]}</td
              >
              <td class="border border-zinc-400 px-1 text-zinc-800"
                >{result.subjects[i].correct.length}</td
              >
              <td class="border border-zinc-400 px-1 text-zinc-800"
                >{result.subjects[i].incorrect.length}</td
              >
              <td class="border border-zinc-400 px-1 text-zinc-800"
                >{result.subjects[i].unattempted.length}</td
              >
              <td class="border border-zinc-400 px-1 text-zinc-800">{result.subjects[i].score}</td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  </div>
  <div>
    <h2 class="mt-2 text-xl font-bold">{lekh.detailed_report}</h2>
    <div class="list-disc">
      <li>
        <span class="font-semibold text-green-700">{lekh.correct_ques}</span> :-
        <span class="text-slate-700"
          >{result.subjects[0].correct
            .concat(result.subjects[1].correct)
            .concat(result.subjects[2].correct)
            .join(', ')}</span
        >
      </li>
      <li>
        <span class="font-semibold text-red-700">{lekh.incorrect_ques}</span> :-
        <span class="text-slate-700"
          >{result.subjects[0].incorrect
            .concat(result.subjects[1].incorrect)
            .concat(result.subjects[2].incorrect)
            .join(', ')}</span
        >
      </li>
      <li>
        <span class="font-semibold text-blue-700">{lekh.unattempted_ques}</span> :-
        <span class="text-slate-700"
          >{result.subjects[0].unattempted
            .concat(result.subjects[1].unattempted)
            .concat(result.subjects[2].unattempted)
            .join(', ')}</span
        >
      </li>
    </div>
  </div>
  <div class="overflow-x-scroll text-center">
    <table class="mt-4 border-collapse rounded-xl border-2 border-blue-800">
      <thead>
        <tr>
          <th class="border border-slate-900 bg-zinc-200 px-1" />
          <th class="border border-slate-900 bg-zinc-200 px-1">{lekh.table.given_answer}</th>
          <th class="border border-slate-900 bg-zinc-200 px-2">{lekh.table.correct_answer}</th>
          <th class="border border-slate-900 bg-zinc-200 px-2">{lekh.table.question_id}</th>
          <th class="border border-slate-900 bg-zinc-200 px-2">{lekh.table.correct_answer_id}</th>
          <th class="border border-slate-900 bg-zinc-200 px-2">{lekh.table.given_answer_id}</th>
        </tr>
      </thead>
      <tbody>
        {#each { length: data.GivenAnswer.length } as _, i (i)}
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
</div>
