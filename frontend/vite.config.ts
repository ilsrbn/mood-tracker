import { fileURLToPath } from "url";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      "@wails": fileURLToPath(
        new URL("./wailsjs/go/main/App", import.meta.url),
      ),
      // "@api": fileURLToPath(
      // new URL("./src/core/api/index.ts"),
      // import.meta.url,
      // ),
    },
  },
});
