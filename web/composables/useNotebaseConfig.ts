import { useLocalStorage } from '#imports'

interface NotebaseConfig {
  dateLocales: string
  showFilters: boolean
  showExtra: boolean
}
const DEFAULT_DATE_LOCALE = 'en-CA'
const DEFAULT_SHOW_FILTERS = false
const DEFAULT_SHOW_EXTRA = false

const DEFAULT_CONFIG: NotebaseConfig = {
  dateLocales: DEFAULT_DATE_LOCALE,
  showFilters: DEFAULT_SHOW_FILTERS,
  showExtra: DEFAULT_SHOW_EXTRA,
}
const useNotebaseLocalConfig = () =>
  useLocalStorage<NotebaseConfig>('notebase-config', DEFAULT_CONFIG, { deep: true })

export function useNotebaseConfig() {
  const config = useNotebaseLocalConfig()

  function setDateLocale(locale: string) {
    config.value.dateLocales = locale
  }

  function setShowFilters(show: boolean) {
    config.value.showFilters = show
  }

  function setShowExtra(show: boolean) {
    config.value.showExtra = show
  }

  return {
    config,
    setDateLocale,
    setShowFilters,
    setShowExtra,
  }
}
