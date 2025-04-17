import { writable, type Writable } from "svelte/store";

// place files you want to import through the `$lib` alias in this folder.
export const tasks:any = writable([])
export const user:any = writable({})
export const error:Writable<string> = writable("")

export const API_URL = "https://tasker-back-production.up.railway.app"