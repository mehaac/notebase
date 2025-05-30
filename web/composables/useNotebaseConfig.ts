import { useLocalStorage } from '#imports'

interface NotebaseConfig {
  dateLocales: string
}
const DEFAULT_DATE_LOCALE = 'en-CA'

const DEFAULT_CONFIG: NotebaseConfig = {
  dateLocales: DEFAULT_DATE_LOCALE,
}
const useNotebaseLocalConfig = () =>
  useLocalStorage<NotebaseConfig>('notebase-config', DEFAULT_CONFIG)

export function useNotebaseConfig() {
  const config = useNotebaseLocalConfig()

  function setDateLocale(locale: string) {
    config.value.dateLocales = locale
  }

  return {
    config,
    setDateLocale,
  }
}
