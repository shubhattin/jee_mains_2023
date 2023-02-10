<script lang="ts">
  import type { PageData } from './$types';
  import LangMetaTags from '@components/tags/LangMetaTags.svelte';
  import type { ResponseDataType } from './data_type';
  import { fetch_post } from '@tools/fetch';

  export let data: PageData;

  $: lekh = data.lekh;

  let mode: 'main' | 'result' = 'main';
  let resp_data: ResponseDataType;

  const fetch_data = async () => {
    const req = await fetch_post('/api/get_result');
    if (!req.ok) return;
    const resp: ResponseDataType = await req.json();
    resp_data = resp;
    mode = 'result';
  };
</script>

<LangMetaTags title={lekh.page_title} description={lekh.page_description} />
<svelte:head>
  <title>{lekh.html_title}</title>
</svelte:head>

<div class="fixed inset-0 h-full w-full bg-gradient-to-r from-amber-50 via-lime-100 to-blue-50 p-2">
  <h1 class="text-center text-3xl font-bold">
    <span
      class="bg-gradient-to-r from-red-600 via-pink-500 to-blue-800 bg-clip-text text-transparent"
      >{lekh.page_title}</span
    >
  </h1>

  <div>
    {#if mode === 'main'}
      <button
        on:click={fetch_data}
        class="rounded-xl border-2 border-red-700 px-1 text-2xl text-green-700">Get Result</button
      >
    {:else if mode === 'result'}
      {JSON.stringify(resp_data)}
    {/if}
  </div>
</div>
