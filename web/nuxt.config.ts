export default defineNuxtConfig({
  compatibilityDate: "2024-11-01",
  future: {
    compatibilityVersion: 4,
  },

  imports: {
    autoImport: false,
  },

  devtools: { enabled: true },
  ssr: false,

  runtimeConfig: {
    public: {
      apiBase: "http://127.0.0.1:8090", // NUXT_PUBLIC_API_BASE=/ in production
    },
  },

  router: {
    options: {
      hashMode: true,
    },
  },

  app: {
    head: {
      link: [
        {
          rel: "stylesheet",
          href: "https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css",
        },
      ],
    },
  },

  modules: ["@vueuse/nuxt", "@pinia/nuxt"],

  nitro: {
    output: {
      // { dir: '.output', serverDir: '.output/server', publicDir: '.output/public' }
      publicDir: "../pb_public",
    },
  },
});
