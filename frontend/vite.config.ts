import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

export default defineConfig({
  plugins: [svelte()],
  server: {
    port: 8080
  },
   build: {
    rollupOptions: {
      external: "/static/cropper.css"
    }
  }, 
  optimizeDeps: {exclude:['svelte-navigator']}
})
