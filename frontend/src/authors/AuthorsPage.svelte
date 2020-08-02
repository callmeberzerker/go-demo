<script lang="ts">
  import { onMount } from "svelte";
  import type { Author } from "../@types";
  import { AuthorService } from "./AuthorService";
  import { createForm } from "svelte-forms-lib";
  import AuthorForm from "./AuthorForm.svelte";
  let authors: Author[] = [];

  onMount(async () => {
    authors = await AuthorService.getAllAuthors();
  });
  export let location: any;

  let authorToBeEdited: Author | null = null;

  const handleEdit = (id: number | string) => () => {
    authorToBeEdited = authors.find((x) => x.id === id);
  };

  const handleSave = async (authorValues: Author) => {
    console.log(authorValues);
    if (authorValues.id) {
      // if there is an id -> it's a existing record
      const updatedAuthor = await AuthorService.updateAuthor(authorValues);
      authors = authors.map((x) => {
        if (x.id === updatedAuthor.id) {
          return updatedAuthor;
        }
        return x;
      });
      authorToBeEdited = updatedAuthor;
    } else {
      const newAuthor = await AuthorService.createAuthor(authorValues);
      authors = [...authors, newAuthor];
      authorToBeEdited = newAuthor;
    }
  };

  const newAuthor = () => {
    authorToBeEdited = {
      id: "",
      firstName: "",
      lastName: "",
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
  <h1 data-path={location.pathname}>Hello from Authors Page!</h1>

  <div class="row">
    <div class="column">
      <ul>
        {#each authors as { id, firstName, lastName }}
          <li>
            {id} - {firstName} - {lastName}
            <button on:click={handleEdit(id)}>Edit</button>
          </li>
        {/each}
      </ul>
      <button on:click={newAuthor}>Add new author</button>
    </div>

    <div class="column">
      {#if authorToBeEdited !== null}
        <AuthorForm author={authorToBeEdited} onSave={handleSave} />
      {/if}
    </div>
  </div>
</main>
