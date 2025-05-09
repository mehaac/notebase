import { pb } from '#imports'

export function getItem(id: string) {
  return pb.collection('files').getOne(id)
}

export const toggleItem = async (id: string) => {
  const item = await getItem(id)
  if (item.frontmatter.completed) {
    item.frontmatter.completed = ''
  }
  else {
    item.frontmatter.completed = (new Date()).toISOString()
  }
  const record = await pb.collection('files').update(item.id, {
    frontmatter: item.frontmatter,
  })
  console.log(record)
}
