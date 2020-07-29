<script lang="ts">
  import { store } from "./StoreConfig";

  const increment = () => {
    store.update((x) => x + 1);
  };
</script>

<button on:click={increment}>+</button>
