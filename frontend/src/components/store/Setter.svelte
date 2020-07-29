<script lang="ts">
  import { store } from "./StoreConfig";

  const update = (event) => {
    store.set(event.target.value);
  };
</script>

<input type="number" on:change={update} />
