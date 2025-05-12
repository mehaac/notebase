import type { Frontmatter, ItemType } from '~/modules/pocketbase/types/types'

export type Item = {
  id: string
  title: string
  content: string
  done: boolean
  type: ItemType
  frontmatter: Frontmatter
}
