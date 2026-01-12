import { fileURLToPath, URL } from 'node:url';

import { defineConfig, loadEnv } from 'vite';
import vue from '@vitejs/plugin-vue';
import tailwindcss from '@tailwindcss/vite';

// https://vite.dev/config/
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd());
  console.info(`[${mode}]: `, env);

  const config = {
    plugins: [vue(), tailwindcss()],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url)),
      },
    },
    define: {
      API_URL: JSON.stringify(env.VITE_APP_API_URL),
      STATIC_URL: JSON.stringify(env.VITE_APP_API_STATIC),
    },
    server: {
      port: 8080,
      // host: true,
    },
  };

  return config;
});
