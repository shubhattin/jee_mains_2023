<script lang="ts">
  import '../app.postcss';
  import { onMount } from 'svelte';
  import PartyTown from '@components/tags/PartyTown.svelte';
  import { change_html_lang_tag } from '@tools/i18n';

  onMount(change_html_lang_tag);
</script>

<slot />
<PartyTown />
