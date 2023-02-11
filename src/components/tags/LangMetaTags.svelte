<script lang="ts">
  import { default_locale, get_link, locale_keys } from '@tools/i18n';
  import { page } from '$app/stores';

  const SITE_URL: string = import.meta.env.VITE_SITE_URL || '';

  export let title: string;
  export let description: string;

  /** Normalised URL without its locale */
  const asPath = (() => {
    let vl = get_link($page.url.pathname, default_locale);
    if (vl[vl.length - 1] == '/') vl = vl.substring(0, vl.length - 1);
    vl = vl == '/' ? '' : vl;
    return vl;
  })();
  const LINKS = locale_keys.map((lcl) => {
    let lnk = SITE_URL;
    let lcl_part = lcl !== default_locale ? lcl : '';
    lnk += (lcl_part.length == 0 ? '' : '/') + lcl_part + asPath;
    return lnk;
  });
</script>

<svelte:head>
  <meta property="og:title" content={title} />
  <meta name="description" content={description} />
  <meta property="og:description" content={description} />
  <meta property="og:site_name" content={title} />
  {#each locale_keys as lcl, i}
    <link rel="alternate" href={LINKS[i]} hreflang={lcl} />
  {/each}
</svelte:head>
