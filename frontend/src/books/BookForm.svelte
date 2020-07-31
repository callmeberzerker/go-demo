<script lang="ts">
  import type { Book } from "../@types";
  import { createForm } from "svelte-forms-lib";
  import { onMount, onDestroy } from "svelte";

  export let book: Book;
  export let onSave: any;
  const {
    // observables state
    form,
    errors,
    state,
    touched,
    isValid,
    isSubmitting,
    isValidating,
    // handlers
    updateInitialValues,
    handleBlur,
    handleChange,
    handleSubmit,
  } = createForm({
    initialValues: book,
    validate: (values) => {
      return {};
    },
    onSubmit: onSave,
  });

  let prevBook: Book;

  $: {
    if (prevBook !== book) {
      updateInitialValues(book);
    }
    prevBook = book;
  }
</script>

<style>
  form {
    border-left: 5px solid grey;
    background-color: grey;
  }
  form.valid {
    border-color: green;
  }
  button:disabled {
    background: grey;
  }
</style>

<form class:valid={$isValid} on:submit={handleSubmit}>

  <label>ID</label>
  <input name="id" on:keyup={handleChange} bind:value={$form.id} disabled />
  {#if $errors.id && $touched.id}
    <small>{$errors.id}</small>
  {/if}

  <label>title</label>
  <input name="title" on:keyup={handleChange} bind:value={$form.title} />
  {#if $errors.title && $touched.title}
    <small>{$errors.title}</small>
  {/if}

  <label>isbn</label>
  <input name="isbn" on:keyup={handleChange} bind:value={$form.isbn} />
  {#if $errors.isbn && $touched.isbn}
    <small>{$errors.isbn}</small>
  {/if}

  <label>author id</label>
  <input name="authorId" on:keyup={handleChange} bind:value={$form.authorId} />
  {#if $errors.authorId && $touched.authorId}
    <small>{$errors.authorId}</small>
  {/if}

  <!-- 
    can't use disabled={!$isValid} for submit button
    since the functionality is broken due to $isValid checking
    if ALL fields are touched (which is wrong). See:

    * https://github.com/tjinauyeung/svelte-forms-lib/issues/47
    * https://github.com/tjinauyeung/svelte-forms-lib/pull/28
  
  -->
  <button type="submit">
    {#if $isSubmitting}loading...{:else}submit{/if}
  </button>
</form>
