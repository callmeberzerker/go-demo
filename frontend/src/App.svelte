<script lang="ts">
  import { Link, Route, Router } from "svelte-routing";
  import BooksPage from "./books/BooksPage.svelte";
  import AuthorsPage from "./authors/AuthorsPage.svelte";

  export let url = "";
</script>

<Router {url}>
  <nav>
    <Link to="/authors">Authors</Link>
    <Link to="/books">Books</Link>
  </nav>
  <div>
    <Route path="authors" component={AuthorsPage} />

    <Route path="books" component={BooksPage} />
  </div>
</Router>

<div>TEST!</div>
