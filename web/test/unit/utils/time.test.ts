import { describe, it, expect, vi } from 'vitest'
import { parseDate } from '../../../modules/pocketbase/runtime/utils/time'

// Mock the @internationalized/date module
vi.mock('@internationalized/date', () => ({
  parseDateTime: vi.fn().mockImplementation(dateTimeString => ({
    dateTimeString,
  })),
}))

describe('parseDate', () => {
  it('should correctly parse date strings', () => {
    const testDate = '2023-01-15T12:30:45Z'
    const result = parseDate(testDate)

    // The result should contain the ISO string with milliseconds truncated
    expect(result).toEqual({
      dateTimeString: '2023-01-15T12:30:45',
    })
  })
})
