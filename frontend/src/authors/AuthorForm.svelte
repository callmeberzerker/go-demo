<script lang="ts">
  import type { Book, Author } from "../@types";
  import { createForm } from "svelte-forms-lib";
  import { onMount, onDestroy } from "svelte";

  export let author: Author;
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
    initialValues: author,
    validate: (values) => {
      return {};
    },
    onSubmit: onSave,
  });

  let prevAuthor: Author;

  $: {
    if (prevAuthor !== author) {
      updateInitialValues(author);
    }
    prevAuthor = author;
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

  <label>first name</label>
  <input
    name="firstName"
    on:keyup={handleChange}
    bind:value={$form.firstName} />
  {#if $errors.firstName && $touched.firstName}
    <small>{$errors.firstName}</small>
  {/if}

  <label>last name</label>
  <input name="lastName" on:keyup={handleChange} bind:value={$form.lastName} />
  {#if $errors.lastName && $touched.lastName}
    <small>{$errors.lastName}</small>
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
