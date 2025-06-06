export default defineNuxtConfig({
  modules: [
    '@vueuse/nuxt',
    '@pinia/nuxt',
    '@nuxt/eslint',
    '@nuxt/ui',
    '@nuxtjs/mdc',
    '~/modules/pocketbase/module',
    '@pinia/colada-nuxt',
    '@nuxt/icon',
    '@vite-pwa/nuxt',
  ],
  ssr: false,
  imports: {
    autoImport: false,
  },
  devtools: { enabled: true },
  app: {
    head: {
      title: 'Notebase',
      viewport: 'width=device-width, initial-scale=1, maximum-scale=1, viewport-fit=cover',
      htmlAttrs: {
        lang: 'en-US',
      },
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.svg' },
      ],
      meta: [
        { name: 'apple-mobile-web-app-status-bar-style', content: 'black-translucent' },
        { name: 'mobile-web-app-capable', content: 'yes' },
        { name: 'apple-mobile-web-app-capable', content: 'yes' },
      ],
    },
  },
  css: ['~/assets/css/main.css'],
  router: {
    options: {
      hashMode: true,
    },
  },
  mdc: {
    highlight: {
      shikiEngine: 'oniguruma',
      theme: {
        default: 'material-theme-palenight',
      },
      langs: ['js', 'ts', 'yaml', 'markdown', 'json'],
    },
  },
  runtimeConfig: {
    public: {
      apiBase: 'http://127.0.0.1:8090', // NUXT_PUBLIC_API_BASE=/ in production
    },
  },
  future: {
    compatibilityVersion: 4,
  },
  compatibilityDate: '2024-11-01',
  eslint: {
    config: {
      stylistic: true,
      formatters: true,
    },
  },
  icon: {
    provider: 'none',
    clientBundle: {
      scan: true,
    },
  },
  pwa: {
    manifest: {
      name: 'Notebase',
      short_name: 'NB',
      background_color: '#18181b',
      theme_color: '#51a2ff',
      display: 'fullscreen',
    },
  },
})
