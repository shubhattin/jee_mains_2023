<script lang="ts">
  import '../app.postcss';
  import '@fontsource/roboto/latin.css';
  import { onMount } from 'svelte';
  import PartyTown from '@components/tags/PartyTown.svelte';
  import { change_html_lang_tag } from '@tools/i18n';
  import GoogleAnalytics from '@components/tags/GoogleAnalytics.svelte';

  onMount(change_html_lang_tag);
</script>

<slot />
<PartyTown />
<GoogleAnalytics />
