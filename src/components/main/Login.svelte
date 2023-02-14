<script lang="ts">
  import type { ResponseDataType } from '@components/main/type';
  import { lekhAH, viewCountData, mode, datt } from '@state/main';
  import { fetch_post } from '@tools/fetch';

  $: lekh = $lekhAH.home;
  const fetch_sample_data_result = async () => {
    const req = await fetch_post('/api/get_sample_result');
    if (!req.ok) return;
    const resp: ResponseDataType = await req.json();
    $datt = resp;
    $mode = 'result';
  };
</script>

<button
  on:click={fetch_sample_data_result}
  class="cursor-button mt-2 rounded-xl border-2 border-red-700 px-1 text-lg text-green-700 active:border-green-800 active:text-blue-700"
>
  {lekh.get_result}
</button>
