<script>
  const cities = [
    {
      id: "1",
      name: "Skopje",
    },
    {
      id: "2",
      name: "Stockholm",
    },
    {
      id: "3",
      name: "Kumanovo",
    },
  ];
</script>

<h1>European cities</h1>
<ul>
  {#each cities as { name, id }}
    <li>{id}: {name}</li>
  {/each}
</ul>
