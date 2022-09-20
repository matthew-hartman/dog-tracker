import { writable } from 'svelte/store';
import PocketBase from 'pocketbase'

const addr = window.location.href
const addr1 = "http://localhost:8090"

export const client = writable(new PocketBase(addr));
export const dogs = writable([]);
export const userID = writable(null);
