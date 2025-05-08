export type Item = {
  id: string
  title: string
  content: string
  completed: string
  type: ItemType
  frontmatter: object
}

export type DebtFrontmatter = {
  currency: string
}

export enum ItemType {
  Track = 'track',
  Debt = 'debt',
  Task = 'task',
}
