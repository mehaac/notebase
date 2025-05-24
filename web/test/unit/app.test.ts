import { describe, it, expect } from 'vitest'
import { setup } from '@nuxt/test-utils/e2e'

describe('App Builds', () => {
  it('has a valid app.vue', async () => {
    // Only check that app.vue exists and can be parsed
    await setup()

    // Test passes if setup doesn't throw
    expect(true).toBe(true)
  })
})
