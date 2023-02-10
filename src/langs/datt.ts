import { locales, default_locale } from './locales.json';
import { load } from 'js-yaml';
import { set_val_from_adress, process_json } from '@tools/json';
import * as fs from 'fs';
import { getMarkdownValues, replace_link_text, markdownify_keys } from './process';
import type { dattType } from './model';
export type { dattType } from './model';
export { default_locale } from './locales.json';

export type langKey = keyof typeof locales; // Language list
export const langNames = Object.values(locales);
export const get_locale = (locale: string) => (locale || default_locale) as langKey;

const reuseValues = (datt: dattType) => {
  const reuseMap: [string, string[]][] = [];
  for (let x of reuseMap)
    for (let y of x[1]) set_val_from_adress(y.substring(4).split('.').join('/'), datt, x[0]);
  return datt;
};
const db: { [x in langKey]?: dattType } = {};

const main = (locale: langKey, type = '') => {
  if (import.meta.env.PROD && locale in db) return db[locale];
  const STR_LOADED = replace_link_text(
    fs.readFileSync(`./src/langs/data/${locales[locale]}.yaml`).toString()
  );
  const MARKDOWN = getMarkdownValues(
    replace_link_text(fs.readFileSync(`./src/langs/data/markdown/${locales[locale]}.md`).toString())
  );
  let dt = process_json(reuseValues(load(STR_LOADED) as dattType), (vl: string) =>
    markdownify_keys(vl, MARKDOWN)
  ) as dattType;
  db[locale] = dt;
  return dt;
};
export default main;
