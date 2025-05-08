import { pb, type Item } from "#imports";
import { defineStore } from "pinia";
import type { ItemType } from "~/utils/types";

export const useActivitiesStore = defineStore("activitiesStore", {
  state: () => ({
    items: [] as Item[],
    item: undefined as Item | undefined,
    itemTypes: new Set<ItemType>(),
  }),
  actions: {
    async load(filter: string) {
      if (filter.length === 0) {
        filter = "path ~ 'inbox/activities/%'";
      }
      const result = await pb.collection("files").getFullList({
        filter: filter,
      });
      this.items = result.map((item): Item => {
        this.itemTypes.add(item.frontmatter.type);
        return {
          id: item.id,
          title: item.frontmatter.title || item.frontmatter.summary,
          content: item.content,
          completed: item.done,
          type: item.frontmatter.type,
          frontmatter: item.frontmatter,
        };
      });
    },
    setItem(id: string) {
      this.item = this.items.find((item) => item.id === id);
    },
  },
});
