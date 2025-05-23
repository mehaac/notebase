import { z } from 'zod'

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
  type: z.enum(Object.values(itemTypes)).nullish(),
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

export const recordSchema = z.object({
  id: z.string(),
  content: z.string(),
  created: z.string(),
  path: z.string(),
  slug: z.string(),
  updated: z.string(),
  frontmatter: z.union([frontmatterSchema, z.string().transform(() => ({} as Record<string, unknown>))]),
})

export type ItemRecord = z.infer<typeof recordSchema>

export type Frontmatter = z.infer<typeof frontmatterSchema>

export type DebtFrontmatter = z.infer<typeof debtFrontmatterSchema>
export type DebtTransaction = z.infer<typeof debtTransactionSchema>

export type TrackFrontmatter = z.infer<typeof trackFrontmatterSchema>
