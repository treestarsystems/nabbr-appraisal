import { defineConfig, loadEnv } from 'vite';
import vue from '@vitejs/plugin-vue';

export default defineConfig(({ mode }) => {
  process.env = { ...process.env, ...loadEnv(mode, process.cwd()) };
  return {
    plugins: [vue()],
    server: {
      proxy: {
        '/api': {
          target: process.env.VITE_BACKEND_API_BASE_URL,
          changeOrigin: true,
          rewrite: path => path.replace(/^\//, ''),
        },
      },
    },
    build: {
      outDir: '../dist/frontend/',
      chunkSizeWarningLimit: 1000,
      rollupOptions: {
        output: {
          manualChunks(id) {
            if (id.includes('node_modules')) {
              return id.toString().split('node_modules/')[1].split('/')[0].toString();
            }
          },
        },
      },
    },
  };
});
