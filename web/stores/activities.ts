import { ref, pb, type Item } from "#imports";
import { defineStore } from "pinia";
import type { ItemType } from "~/utils/types";

export const useActivitiesStore = defineStore("activitiesStore", () => {
  const items = ref<Item[]>([]);
  const item = ref<Item | undefined>(undefined);
  const itemTypes = ref<Set<ItemType>>(new Set());

  const load = async (filter: string) => {
    if (filter.length === 0) {
      filter = "path ~ 'inbox/activities/%'";
    }
    const result = await pb.collection("files").getFullList({
      filter: filter,
    });
    items.value = result.map((item): Item => {
      itemTypes.value.add(item.frontmatter.type);
      return {
        id: item.id,
        title: item.frontmatter.title || item.frontmatter.summary,
        content: item.content,
        completed: item.done,
        type: item.frontmatter.type,
        frontmatter: item.frontmatter,
      };
    });
  };

  return { items, item, itemTypes, load };
});
