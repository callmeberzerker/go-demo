<script lang="ts">
  import { store } from "./StoreConfig";
  import Incrementer from "./Incrementer.svelte";
  import Decrementer from "./Decrementer.svelte";
  import Setter from "./Setter.svelte";
</script>

<h1>The count value is {$store}</h1>

<Incrementer />

<Decrementer />

<Setter />
