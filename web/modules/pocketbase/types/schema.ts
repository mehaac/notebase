import { z } from 'zod'
import type { RecordModel } from 'pocketbase'
import type { Item } from './types'

export const itemTypes = {
  track: 'track',
  debt: 'debt',
  task: 'task',
  none: 'none',
} as const

export type ItemType = (typeof itemTypes)[keyof typeof itemTypes]

export const debtTransactionSchema = z.object({
  amount: z.number(),
  created: z.string(),
  comment: z.string().nullish(),
})

export const baseFrontmatterSchema = z.looseObject({
  title: z.string().nullish(),
  summary: z.string().nullish(),
  type: z.enum(Object.values(itemTypes)).default(itemTypes.none),
  completed: z.string().nullish(),
  aliases: z.string().nullish(),
  tags: z.array(z.string()).nullish(),
})

export const debtFrontmatterSchema = baseFrontmatterSchema.extend({
  currency: z.string(),
  transactions: z.array(debtTransactionSchema),
})

export const trackFrontmatterSchema = baseFrontmatterSchema.extend({
  season: z.number(),
  episode: z.number(),
  next_episode: z.string(),
  url: z.string(),
})

export const frontmatterSchema = baseFrontmatterSchema.extend({
  currency: z.string().nullish(),
  transactions: z.array(debtTransactionSchema).nullish(),
  season: z.number().nullish(),
  episode: z.number().nullish(),
  next_episode: z.string().nullish(),
  url: z.string().nullish(),
  due: z.string().nullish(),
  status: z.string().nullish(),
  priority: z.string().nullish(),
  modified: z.string().nullish(),
  rating: z.string().nullish(),
})

export type Frontmatter = z.infer<typeof frontmatterSchema>

export type DebtFrontmatter = z.infer<typeof debtFrontmatterSchema>
export type DebtTransaction = z.infer<typeof debtTransactionSchema>

export type TrackFrontmatter = z.infer<typeof trackFrontmatterSchema>

export function transformItem(item: RecordModel): Item {
  const frontmatter = frontmatterSchema.parse(item.frontmatter ? item.frontmatter : {})
  const title = frontmatter.title ?? frontmatter.summary ?? 'None'
  const type = frontmatter.type ?? itemTypes.none
  return {
    id: item.id,
    title,
    content: item.content,
    done: Boolean(frontmatter.completed),
    type,
    frontmatter,
  }
}
