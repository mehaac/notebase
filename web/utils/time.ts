import { parseDateTime } from '@internationalized/date'

export function parseDate(date: string) {
  return parseDateTime(`${new Date(date).toISOString().split('.')[0]}`)
}

export function formatDateShort(date: Date) {
  return date.toLocaleDateString('en-CA', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
  })
}
