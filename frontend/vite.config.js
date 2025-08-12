import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

export default defineConfig({
  plugins: [vue()],
  server: {
    port: 5173,
    proxy: {
      "/api": {
        target: "http://localhost:8080", // ton backend
        changeOrigin: true,
        ws: true, // si tu utilises websockets un jour
      },
    },
  },
});
