<script lang="ts">
  let counter = 0;
  $: squared = counter ** 2;
  const handleClick = (): void => {
    ++counter;
  };
</script>

<button on:click={handleClick}>Click me</button>
<h1>Button is clicked {counter} times!</h1>

<h2>The squared number of that is: {squared}</h2>
