<script lang="ts">
  import { fly } from 'svelte/transition';
  import type { ResponseDataType } from '@components/main/type';
  import { lekhAH } from '@state/main';

  $: lekh = $lekhAH.result.questions_tab;

  export let data: ResponseDataType['data'];
</script>

<div transition:fly class="list-decimal">
  {#each { length: data.GivenAnswer.length } as _, i (i)}
    <li class="mb-4">
      <img class="inline-block" src={data.QuestionIMG[i]} alt={`Question ${i + 1}`} />
      <div class="mt-2 ml-4 list-disc">
        {#if data.Type[i] === 'MCQ'}
          {#each ['A', 'B', 'C', 'D'] as opt, opt_num (opt)}
            <li class="mb-1">
              {opt}.
              <img
                class="inline-block"
                src={[data.Option1IMG, data.Option2IMG, data.Option3IMG, data.Option4IMG][opt_num][
                  i
                ]}
                alt={`Question ${i + 1} Option ${opt}`}
              />
            </li>
          {/each}
        {/if}
        <li>{lekh.correct_answer} : {data.CorrectAnswer[i]}</li>
        <li>{lekh.marked_answer} : {data.GivenAnswer[i]}</li>
      </div>
    </li>
    <div class="h-1 w-full bg-gray-400" />
  {/each}
</div>
