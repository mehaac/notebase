# PocketBase Module for Nuxt 3

A simple Nuxt module for integrating PocketBase with your Nuxt 3 application.

## Setup

Add the module to your `nuxt.config.ts`:

```ts
export default defineNuxtConfig({
  modules: [
    '~/modules/pocketbase'
  ],
  
  // Configure using the common apiBase (recommended)
  runtimeConfig: {
    public: {
      apiBase: 'http://your-pocketbase-url'
    }
  },
  
  // Optional: Configure with module-specific options
  pocketbase: {
    // The URL will automatically use apiBase if not specified
    type: 'pb' | 'mock'
  }
})
```

## Usage

### Using the Injected Services

The module provides one injected service:

1. `$client` - The PicketBase or Mock client with custom helpers

```js
<script setup>
const { $client } = useNuxtApp()

// Using the raw PocketBase client
const records = await $client.getList(1, 10, filter)

// Using the extended client with custom helpers
const item = await $client.getItem('item_id')
await $client.toggleItem('item_id')
</script>
```

### .env

```
NUXT_PUBLIC_POCKETBASE_TYPE=pb or mock
```

### Using the Composables (Recommended)

The module auto-imports these composables for easier use:
or use #pocketbase-imports
```vue
<script setup>
import { usePocketBaseClient } from '#pocketbase-imports'

const pbClient = usePocketBaseClient()
const item = await pbClient.getItem('item_id')
await pbClient.toggleItem('item_id')
</script>
```

## Custom Helpers

```ts 
/**
 * Base client interface for database operations
 * Provides methods for CRUD operations on items and user authentication
 */
export interface BaseClient {
  /**
   * Retrieves a single item by its unique identifier
   * @param id - The unique identifier of the item to retrieve
   * @returns Promise resolving to the item record
   */
  getItem: (id: string) => Promise<RecordModel>

  /**
   * Toggles the completion status of an item
   * @param id - The unique identifier of the item to toggle
   * @returns Promise resolving to the updated item record
   */
  toggleItem: (id: string) => Promise<RecordModel>

  /**
   * Adds a new debt transaction to an existing debt item
   * @param id - The unique identifier of the debt item
   * @param amount - The transaction amount (positive or negative)
   * @param comment - Optional description of the transaction
   * @returns Promise resolving to the updated debt record
   */
  addDebtTransaction: (id: string, amount: number, comment: string) => Promise<RecordModel>

  /**
   * Checks if the current client has valid authentication
   * @returns Promise resolving to authentication status (true/false)
   */
  isAuthenticated: () => Promise<boolean>

  /**
   * Removes all authentication data from the client
   * @returns Promise that resolves when auth is cleared
   */
  clearAuth: () => Promise<void>

  /**
   * Authenticates a user with email and password credentials
   * @param payload - Object containing email and password
   * @returns Promise resolving to auth response with user data
   */
  authenticatedUser: (payload: { email: string, password: string }) => Promise<RecordAuthResponse<RecordModel>>

  /**
   * Retrieves a paginated list of items with optional filtering
   * @param page - Page number to retrieve (starts at 1)
   * @param pageSize - Number of items per page
   * @param filter - Optional filter query string
   * @returns Promise resolving to paginated list result
   */
  getList: (page: number, pageSize: number, filter: string) => Promise<ListResult<RecordModel>>
}
```
