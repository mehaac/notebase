import type { ItemRecord } from '#pocketbase-imports'

export function toggleItem(item: ItemRecord) {
  const frontmatter = item.frontmatter || {}
  if (frontmatter?.completed) {
    frontmatter.completed = ''
  }
  else {
    frontmatter.completed = new Date().toISOString().split('.')[0]!
  }

  return {
    ...item,
    frontmatter,
  }
}
