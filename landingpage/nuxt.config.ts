// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: [
    '@nuxtjs/tailwindcss',
  ],
  css: [
    '@/assets/css/main.css',
  ],
  tailwindcss: {
    cssPath: './assets/css/main.css',
    configPath: './tailwind.config.ts',
  }
})
