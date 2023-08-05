// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ['@nuxtjs/tailwindcss', '@nuxtjs/apollo', 'nuxt-icon', "@nuxtjs/color-mode", '@nuxt/image-edge'],
  css: ["@/assets/css/main.css"],
  image: {
      domains: ['rickandmortyapi.com']
  },
  colorMode: {
      classSuffix: ''
  },
  apollo: {
      clients: {
          default: {
              httpEndpoint: 'http://localhost:8080'
          }
      },
  },
})
