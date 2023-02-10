<script lang="ts">
  import { get_current_locale, change_locale, locales, locale_keys } from '@tools/i18n';
  import type { langKey } from '@langs/datt';
  import { clsx } from '@tools/clsx';
  import { page } from '$app/stores';
  import { onDestroy } from 'svelte';
  import Select from '@components/Select.svelte';
  import { writable } from 'svelte/store';

  let value = get_current_locale($page.params.lang);
  const unsubscribe = page.subscribe((page) => {
    const locale = get_current_locale(page.params.lang);
    if (locale !== value) value = locale;
  });
  onDestroy(unsubscribe);

  const val = writable(value);
  $: options = (() => {
    const opts: any = {};
    for (let x in locales)
      opts[x] = {
        text: locales[x as langKey],
        className: clsx('bg-black font-semibold', $val === x ? 'text-yellow-400' : 'text-white')
      };
    return opts;
  })();
</script>

<Select
  value={val}
  className="px-1 font-bold bg-black text-white border-2 border-lime-500 outline-none rounded-lg"
  onChange={() => change_locale($val)}
  {options}
/>
