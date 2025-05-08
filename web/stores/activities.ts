import { pb, type Item, shallowRef, ref, watchDebounced } from "#imports";
import { defineStore } from "pinia";
import type { ItemType } from "~/utils/types";

export const useActivitiesStore = defineStore("activitiesStore", () => {
  const items = shallowRef<Item[]>([]);
  const itemTypes = shallowRef<Set<ItemType>>(new Set());
  const item = shallowRef<Item | undefined>(undefined);
  const query = shallowRef("");

  const load = async () => {
    let filter = query.value;
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

  watchDebounced(
    query,
    async () => {
      await load();
    },
    { debounce: 300 },
  );

  return { items, item, itemTypes, load, query };
});
