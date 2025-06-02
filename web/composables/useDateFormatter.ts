import { useNotebaseConfig } from '#imports'
import { DateFormatter } from '@internationalized/date'

export function useDateFormatter() {
  const notebaseConfig = useNotebaseConfig()

  function formatShortDate(date: Date) {
    const formatter = new DateFormatter(notebaseConfig.config.value.dateLocales, {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
    })
    return formatter.format(date)
  }

  return {
    formatShortDate,
  }
}
