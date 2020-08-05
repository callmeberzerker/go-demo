<script lang="ts">
  import { onMount } from "svelte";
  import type { Book, BookCreate } from "../@types";
  import { BookService } from "./BookService";
  import { createForm } from "svelte-forms-lib";
  import BookForm from "./BookForm.svelte";
  let books: Book[] = [];

  onMount(async () => {
    books = await BookService.getAllBooks();
  });
  export let location: any;

  let bookToBeEdited: Book | null = null;

  const handleEdit = (id: number | string) => () => {
    bookToBeEdited = books.find((x) => x.id === id);
  };

  const handleSave = async (bookValues: Book) => {
    console.log(bookValues);
    if (bookValues.id) {
      // if there is an id -> it's a existing record
      const updatedBook = await BookService.updateBook(bookValues);
      books = books.map((x) => {
        if (x.id === updatedBook.id) {
          return updatedBook;
        }
        return x;
      });
      bookToBeEdited = updatedBook;
    } else {
      const { id, ...rest } = bookValues;
      const newBookValues = { ...rest } as BookCreate;
      const newBook = await BookService.createBook(newBookValues);
      books = [...books, newBook];
      bookToBeEdited = newBook;
    }
  };

  const newBook = () => {
    bookToBeEdited = {
      id: "",
      isbn: "",
      title: "",
      authorId: 1,
    };
  };
</script>

<style>
  .row {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    width: 100%;
  }

  .column {
    display: flex;
    flex-direction: column;
    flex-basis: 100%;
    flex: 1;
  }

  main {
    margin: 15px;
    background-color: red;
  }
</style>

<main>
  <h1 data-path={location.pathname}>Hello from Books Page!</h1>

  <div class="row">
    <div class="column">
      <ul>
        {#each books as { id, title }}
          <li>
            {id} - {title}
            <button on:click={handleEdit(id)}>Edit</button>
          </li>
        {/each}
      </ul>
      <button on:click={newBook}>Add new book</button>
    </div>

    <div class="column">
      {#if bookToBeEdited !== null}
        <BookForm book={bookToBeEdited} onSave={handleSave} />
      {/if}
    </div>
  </div>
</main>
