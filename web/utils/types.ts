import type { Frontmatter, ItemType } from '#pocketbase-imports'

export type Item = {
  id: string
  title: string
  content: string
  done: boolean
  type: ItemType
  frontmatter: Frontmatter
}
