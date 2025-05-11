# PocketBase Module for Nuxt 3

A simple Nuxt module for integrating PocketBase with your Nuxt 3 application.

## Setup

Add the module to your `nuxt.config.ts`:

```ts
export default defineNuxtConfig({
  modules: [
    '~/modules/pocketbase'
  ],
  
  // Configure with module options
  pocketbase: {
    url: 'http://your-pocketbase-url'
  },
  
  // Or use runtime config
  runtimeConfig: {
    pocketbase: {
      url: 'http://your-pocketbase-url'
    },
    public: {
      // To make URL available on client-side
      pocketbase: {
        url: 'http://your-pocketbase-url'
      }
    }
  }
})
```

## Usage

### Using the Injected Services

The module provides two injected services:

1. `$pb` - The raw PocketBase client
2. `$pbClient` - The extended client with custom helpers

```vue
<script setup>
const { $pb, $pbClient } = useNuxtApp()

// Using the raw PocketBase client
const records = await $pb.collection('my_collection').getList(1, 10)

// Using the extended client with custom helpers
const item = await $pbClient.getItem('item_id')
await $pbClient.toggleItem('item_id')
</script>
```

### Using the Composables (Recommended)

The module auto-imports these composables for easier use:

```vue
<script setup>
// Use the raw PocketBase client
const pb = usePocketBase()
const records = await pb.collection('my_collection').getList(1, 10)

// Use the extended client with custom helpers
const pbClient = usePocketBaseClient()
const item = await pbClient.getItem('item_id')
await pbClient.toggleItem('item_id')
</script>
```

## Custom Helpers

- `getItem(id)` - Get a file item by ID
- `toggleItem(id)` - Toggle the completed status of an item
- `addDebtTransaction(id, amount, comment)` - Add a debt transaction to an item 
