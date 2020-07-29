<script lang="ts">
  import { onMount } from "svelte";
  import type { Book } from "../@types";
  import { BookService } from "./BookService";
  let books: Book[] = [];

  onMount(async () => {
    books = await BookService.getAllBooks();
  });
  export let location: any;
</script>

<h1 data-path={location.pathname}>Hello from Books Page!</h1>

<ul>
  {#each books as { id, title }}
    <li>{id} - {title}</li>
  {:else}
    <li>Loading...</li>
  {/each}
</ul>
