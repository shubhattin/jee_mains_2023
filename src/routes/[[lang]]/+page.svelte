<script lang="ts">
  import type { PageData } from './$types';
  import LangMetaTags from '@components/tags/LangMetaTags.svelte';
  import type { ResponseDataType } from './data_type';
  import { lekhAH, viewCountData } from './data_type';
  import { fetch_post } from '@tools/fetch';
  import Icon from '@tools/Icon.svelte';
  import AiOutlineGithub from 'svelte-icons-pack/ai/AiOutlineGithub';
  import LangChange from '@components/LangChange.svelte';
  import { onMount } from 'svelte';

  export let data: PageData;

  $: $lekhAH = data.lekh;

  $: lekh = $lekhAH;

  let mode: 'main' | 'result' = 'main';
  let datt: ResponseDataType;
  let counted = false;

  onMount(async () => {
    if (import.meta.env.DEV) return;
    const req = await fetch_post('/api/page_view_count');
    if (!req.ok) return;
    const data: {
      page_view_count: number;
      result_view_count: number;
    } = await req.json();
    $viewCountData = [data.page_view_count, data.result_view_count];
    counted = true;
  });
  const fetch_data = async () => {
    const req = await fetch_post('/api/get_result');
    if (!req.ok) return;
    const resp: ResponseDataType = await req.json();
    $viewCountData[1]++;
    datt = resp;
    mode = 'result';
  };
</script>

<LangMetaTags title={lekh.title} description={lekh.description} />
<svelte:head>
  <title>{lekh.title}</title>
</svelte:head>
<div
  class="flex h-full w-full flex-col justify-between overflow-y-scroll bg-gradient-to-r from-amber-50 via-lime-50 to-yellow-50"
>
  <div class="p-2">
    <h1 class="text-center text-2xl font-bold">
      <span class="bg-gradient-to-r from-red-600 to-blue-800 bg-clip-text text-transparent"
        >{lekh.title}</span
      >
    </h1>

    <div>
      {#if mode === 'main'}
        <button
          on:click={fetch_data}
          class="cursor-button mt-2 rounded-xl border-2 border-red-700 px-1 text-lg text-green-700 active:border-green-800 active:text-blue-700"
          >{lekh.home.get_result}</button
        >
      {:else if mode === 'result'}
        {#await import('./Result.svelte') then Result}
          <Result.default {datt} />
        {/await}
      {/if}
    </div>
  </div>
  <footer class="flex justify-center border-t-2 border-t-gray-400 bg-zinc-200 px-2 pt-2 pb-1">
    <div class="flex flex-col justify-center">
      <div>
        <span class="font-bold text-red-600">
          {lekh.home.made_by}
          <a
            class="text-blue-700"
            rel="noreferrer"
            href="https://www.github.com/shubhattin"
            target="_blank">@shubhattin</a
          >
        </span>
        <a
          href="https://github.com/shubhattin/jee_mains_2023_score_calculator"
          class="ml-2"
          rel="noreferrer"
          target="_blank"
        >
          <Icon className="text-2xl" src={AiOutlineGithub} />
        </a>
      </div>
      <div class="flex justify-center">
        <span class="mr-1 h-6 w-6 bg-cover">
          <svg viewBox="0 0 239 217">
            <g transform="translate(0,217) scale(0.1,-0.1)">
              <path
                d="M485 1923 c-47 -10 -145 -45 -145 -52 0 -6 32 -126 46 -173 4 -15 9,-15 37 -3 64 26 150 38 196 25 62 -16 81 -40 81 -101 0 -28 -5 -59 -10 -70,-16 -30 -79 -56 -153 -64 l-67 -7 0 -93 0 -92 86 -6 c134 -10 188 -56 188,-162 0 -97 -48 -148 -142 -149 -134 -1 -241 110 -323 335 -17 46 -32 85 -34,87 -1 2 -44 -14 -95 -35 -106 -43 -104 -35 -49 -171 114 -286 283 -432 500,-432 140 0 257 54 318 147 41 62 55 113 55 199 l-1 72 41 6 c23 3 51 9 64 12,l22 6 0 -271 0 -271 100 0 c55 0 100 -2 100 -5 0 -3 -43 -119 -95 -258 -52,-140 -95 -259 -95 -265 0 -9 35 -12 118 -12 l119 0 64 173 64 172 275 0 275 0,64 -172 64 -173 119 0 c88 0 118 3 118 13 0 6 -117 323 -259 702 l-259 690,-122 0 -122 0 -154 -410 -154 -410 0 508 0 507 80 0 80 0 0 95 0 95 -190 0,-190 0 0 -259 0 -258 -56 -18 c-76 -24 -179 -25 -222 -2 l-33 18 44 40 c117,103 131 270 32 387 -63 74 -159 113 -279 111 -45 -1 -90 -3 -101 -6z m1355,-965 c45 -123 85 -229 87 -235 4 -10 -35 -13 -177 -13 -142 0 -181 3 -177 13,2 6 42 112 87 235 46 122 86 222 90 222 4 0 44 -100 90 -222z"
              />
            </g>
          </svg>
        </span>
        <LangChange />
      </div>
      <div class="text-xs">
        <span class="invisible">|</span>
        {#if counted}
          <span class="mr-1 text-zinc-500"
            >{lekh.home.page_views} - <span class="text-slate-600">{$viewCountData[0]}</span></span
          >
          <span class="text-zinc-500"
            >{lekh.home.result_views} -
            <span class="text-slate-600">{$viewCountData[1]}</span></span
          >
        {/if}
      </div>
    </div>
  </footer>
</div>
