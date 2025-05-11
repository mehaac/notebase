// Importing from 'nuxt/kit' is correct for a module file
import { defineNuxtModule, createResolver, addPlugin, addImportsDir } from 'nuxt/kit'
import { defu } from 'defu'

export interface PocketbaseModuleOptions {
  url?: string
  type?: 'mock' | 'pb'
}

export default defineNuxtModule<PocketbaseModuleOptions>({
  meta: {
    name: 'pocketbase',
    configKey: 'pocketbase',
    compatibility: {
      nuxt: '>=3.0.0'
    }
  },
  defaults: {
    type: 'pb',
    url: ''
  },
  setup(options, nuxt) {
    const { resolve } = createResolver(import.meta.url)

    // Get apiBase from config if url is not provided
    const apiBase = (nuxt.options.runtimeConfig.public as any)?.apiBase
    if (!options.url && apiBase) {
      options.url = apiBase
    }

    // Configure runtime config
    nuxt.options.runtimeConfig.pocketbase = defu(
      nuxt.options.runtimeConfig.pocketbase || {},
      options as Record<string, any>
    )
    
    // Make sure url is also available in public config
    nuxt.options.runtimeConfig.public.pocketbase = defu(
      nuxt.options.runtimeConfig.public.pocketbase || {},
      { url: options.url }
    )

    // Register plugin and composables
    addPlugin(resolve('./runtime/plugin'))
    addImportsDir(resolve('./runtime/composables'))
  }
}) 
