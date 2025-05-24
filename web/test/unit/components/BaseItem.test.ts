import { describe, it, expect, vi } from 'vitest'
import { mountSuspended } from '@nuxt/test-utils/runtime'
import BaseItem from '../../../components/BaseItem.vue'

// Mock Vue's resolveDynamicComponent
vi.mock('vue', async () => {
  const actual = await vi.importActual('vue')
  return {
    ...actual,
    resolveDynamicComponent: vi.fn().mockImplementation(() => 'div'),
  }
})

describe('BaseItem', () => {
  it('renders with props', async () => {
    const mockItem = {
      id: '1',
      content: '',
      created: '2023-01-01',
      hash: 'hash123',
      path: '/path/to/item',
      slug: 'item-slug',
      updated: '2023-01-02',
      frontmatter: {
        type: 'track',
      },
    }

    const wrapper = await mountSuspended(BaseItem, {
      props: {
        item: mockItem,
        isList: true,
      },
    })

    expect(wrapper.vm).toBeDefined()
    expect(wrapper.props().item).toEqual(mockItem)
    expect(wrapper.props().isList).toBe(true)
  })
})
