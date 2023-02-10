<script lang="ts">
  import type { PageData } from './$types';
  import LangMetaTags from '@components/tags/LangMetaTags.svelte';
  import type { ResponseDataType } from './data_type';
  import { fetch_post } from '@tools/fetch';
  import Icon from '@tools/Icon.svelte';
  import AiOutlineGithub from 'svelte-icons-pack/ai/AiOutlineGithub';

  export let data: PageData;

  $: lekh = data.lekh;

  let mode: 'main' | 'result' = 'main';
  let datt: ResponseDataType;

  const fetch_data = async () => {
    const req = await fetch_post('/api/get_result');
    if (!req.ok) return;
    const resp: ResponseDataType = await req.json();
    datt = resp;
    mode = 'result';
  };
</script>

<LangMetaTags title={lekh.page_title} description={lekh.page_description} />
<svelte:head>
  <title>{lekh.html_title}</title>
</svelte:head>
<div
  class="fixed inset-0 flex h-full w-full flex-col justify-center overflow-y-scroll bg-gradient-to-r from-amber-50 via-lime-50 to-blue-50"
>
  <div class="flex-grow p-2">
    <h1 class="text-center text-2xl font-bold">
      <span class="bg-gradient-to-r from-red-600 to-blue-800 bg-clip-text text-transparent"
        >{lekh.page_title}</span
      >
    </h1>

    <div>
      {#if mode === 'main'}
        <button
          on:click={fetch_data}
          class="cursor-button rounded-xl border-2 border-red-700 px-1 text-lg text-green-700"
          >Get Result</button
        >
      {:else if mode === 'result'}
        {#await import('./Result.svelte') then Result}
          <Result.default {datt} />
        {/await}
      {/if}
    </div>
  </div>
  <footer class="flex justify-center border-t-2 border-t-gray-400 bg-zinc-200 px-2 pt-2 pb-3">
    <div>
      <span class="font-bold text-red-600">
        Made by <a
          class="text-blue-700"
          rel="noreferrer"
          href="https://www.github.com/shubhattin"
          target="_blank">@shubhattin</a
        >
      </span>
      <a
        href="https://github.com/shubhattin/jee_mains_2023_score_calculator"
        class="ml-4"
        rel="noreferrer"
        target="_blank"
      >
        <Icon className="text-2xl" src={AiOutlineGithub} />
      </a>
    </div>
  </footer>
</div>
