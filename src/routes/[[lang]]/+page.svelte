<script lang="ts">
  import type { PageData } from './$types';
  import LangMetaTags from '@components/tags/LangMetaTags.svelte';

  export let data: PageData;

  $: lekh = data.lekh;
</script>

<LangMetaTags title={lekh.page_title} description={lekh.page_description} />
<svelte:head>
  <title>{lekh.html_title}</title>
</svelte:head>

<h1>{lekh.page_title}</h1>
