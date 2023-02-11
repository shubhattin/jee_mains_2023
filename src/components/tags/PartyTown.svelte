<script>
  import { onMount } from 'svelte';
  import { partytownSnippet } from '@builder.io/partytown/integration';

  // Add the Partytown script to the DOM head
  let scriptEl;
  onMount(() => {
    if (scriptEl) {
      scriptEl.textContent = partytownSnippet();
    }
  });
</script>

<svelte:head>
  <script>
    partytown = {
      forward: ['dataLayer.push']
    };
  </script>

  <!-- `partytownSnippet` is inserted here -->
  <script bind:this={scriptEl}></script>
</svelte:head>
