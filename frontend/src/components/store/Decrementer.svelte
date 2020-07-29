<script lang="ts">
  import { store } from "./StoreConfig";

  const decrement = () => {
    store.update((x) => x - 1);
  };
</script>

<button on:click={decrement}>-</button>
