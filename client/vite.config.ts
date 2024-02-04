import million from "million/compiler";
import react from "@vitejs/plugin-react-swc";
import { defineConfig } from "vite";
import path from 'path';
 
export default defineConfig({
  plugins: [million.vite({ auto: true }), react()],
  resolve: {
    alias: {
      '@style': path.resolve(__dirname, 'styles'),
      '@lib': path.resolve(__dirname, 'lib'),
    },
  }
});