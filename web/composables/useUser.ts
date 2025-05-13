import { useState } from '#app'

export const useUser = () => useState('user', () => ({
  isAuthenticated: false,
}))
