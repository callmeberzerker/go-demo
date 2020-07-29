<script lang="ts">
  const getRandomWarsCharacter = async () => {
    const randomNumber = Math.floor(Math.random() * 10) + 1;
    const apiResponse = await fetch(
      `https://swapi.dev/api/people/${randomNumber}`
    );
    return await apiResponse.json();
  };

  let promise = getRandomWarsCharacter();
</script>

{#await promise}
  <h1>Loading</h1>
{:then character}
  <h1>{character.name}</h1>
{/await}
