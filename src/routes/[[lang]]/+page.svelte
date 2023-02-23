<script lang="ts">
  import type { PageData } from './$types';
  import { lekhAH, viewCountData, mode, counted } from '@state/main';
  import { fetch_post } from '@tools/fetch';
  import { onMount } from 'svelte';
  import { API_URL } from '@components/main/type';
  import { getTime } from '@tools/cookies';
  import LangMetaTags from '@components/tags/LangMetaTags.svelte';
  import Footer from '@components/main/Footer.svelte';
  import Login from '@components/main/Login.svelte';

  export let data: PageData;

  $: $lekhAH = data.lekh;

  $: lekh = $lekhAH;

  onMount(async () => {
    if (import.meta.env.DEV) return;
    let update_count = true;
    if (localStorage.getItem('page_count_time')) {
      const page_count_time = parseInt(localStorage.getItem('page_count_time')!);
      // count will be updated only if page_count_time is greater than current time
      // so that we dont update count on every page load
      if (getTime() - page_count_time < 0) update_count = false;
    }
    const req = await fetch_post(API_URL + `/api/page_view_count?update_count=${update_count}`);
    if (!req.ok) return;
    const {
      page_view_count,
      result_view_count,
      page_count_time
    }: {
      page_view_count: number;
      result_view_count: number;
      page_count_time: number;
    } = await req.json();
    $viewCountData = [page_view_count, result_view_count];
    localStorage.setItem('page_count_time', (getTime() + page_count_time).toString());
    $counted = true;
  });
</script>

<LangMetaTags title={lekh.title} description={lekh.description} />
<svelte:head>
  <title>{lekh.title}</title>
</svelte:head>
<div class="flex h-full w-full flex-col justify-between overflow-y-scroll bg-orange-50">
  <div class="p-2">
    <h1 class="text-center text-2xl font-bold">
      <span class="bg-gradient-to-r from-red-600 to-blue-800 bg-clip-text text-transparent"
        >{lekh.title}</span
      >
    </h1>
    <div>
      {#if $mode === 'main'}
        <Login />
      {:else if $mode === 'result'}
        {#await import('@components/main/Result.svelte') then Result}
          <Result.default />
        {/await}
      {/if}
    </div>
  </div>
  <Footer />
</div>
