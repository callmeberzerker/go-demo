<script lang="ts">
  let name;
  const handleClick = () => {
    alert(`Hello ${name}`);
  };

  const handleInput = (event) => {
    name = event.target.value;
  };
</script>

<h1>Introduce yourself</h1>
<input type="text" on:change={handleInput} />
<button on:click|once={handleClick}>Click me</button>
