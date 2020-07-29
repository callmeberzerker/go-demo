<script lang="ts">
  import { onMount } from "svelte";
  let characters = [];
  onMount(async () => {
    const apiResponse = await fetch("https://swapi.dev/api/people");
    const swPeopleJson = await apiResponse.json();

    characters = swPeopleJson.results;
  });
</script>

<ul>
  {#each characters as { name, height, birth_year }}
    <li>
      <strong>{name}</strong>
      (height: {height}cm, birth year: {birth_year})
    </li>
  {:else}
    <p>Loading...</p>
  {/each}
</ul>
