<script lang="ts">
  import BooksPage from "./books/BooksPage.svelte";
  import { Router, Link, Route } from "svelte-routing";
  import AuthorsPage from "./authors/AuthorsPage.svelte";
</script>

<Router>
  <nav>
    <Link to="/authors">Authors</Link>
    <Link to="/books">Books</Link>
  </nav>
  <main>
    <Route path="/authors" component={AuthorsPage} />
    <Route path="/books" component={BooksPage} />
  </main>
</Router>
