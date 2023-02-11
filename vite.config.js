import { sveltekit } from '@sveltejs/kit/vite';
import { partytownVite } from '@builder.io/partytown/utils';
import path from 'path';

const PROD_MODE = process.env.NODE_ENV === 'production';

/** @type {import('vite').UserConfig} */
const config = {
  plugins: [
    sveltekit(),
    partytownVite({
      // `dest` specifies where files are copied to in production
      dest: path.join(process.cwd(), PROD_MODE ? 'build' : 'static', '~partytown')
    })
  ],
  server: {
    port: 3427,
    strictPort: true
  },
  resolve: {
    alias: {
      '@langs': path.resolve('./src/langs'),
      '@tools': path.resolve('./src/tools'),
      '@components': path.resolve('./src/components')
    }
  }
};

export default config;
