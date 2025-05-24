// @vitest-environment nuxt
import { describe, it, expect, beforeEach } from 'vitest'
import { useUser } from '../../../composables/useUser'

describe('useUser', () => {
  beforeEach(() => {
    // Clear any state between tests
    const user = useUser()
    user.value = { isAuthenticated: false }
  })

  it('returns a state with default values', () => {
    const user = useUser()
    expect(user.value).toEqual({ isAuthenticated: false })
  })

  it('can update the state', () => {
    const user = useUser()
    user.value = { isAuthenticated: true }
    expect(user.value).toEqual({ isAuthenticated: true })
  })
})
