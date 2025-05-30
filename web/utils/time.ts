import { parseDateTime, DateFormatter } from '@internationalized/date'

export function parseDate(date: string) {
  return parseDateTime(`${new Date(date).toISOString().split('.')[0]}`)
}

export function formatDateShort(date: Date) {
  const df = new DateFormatter('en-CA', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
  })
  return df.format(date)
}
