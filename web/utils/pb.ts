import { useRuntimeConfig } from "#app";
import PocketBase from "pocketbase";

const config = useRuntimeConfig();

export const pb = new PocketBase(config.public.apiBase);
