<script>
    import { Router, Route } from "svelte-routing";
    import NavLink from "./components/NavLink.svelte";
    import Country from "./routes/country/Country.svelte";
    import Region from "./routes/region/Region.svelte";
    import Worldwide from "./routes/worldwide/Worldwide.svelte";
  	import { ApolloClient, gql, InMemoryCache, HttpLink } from "apollo-boost";
    import { setClient, query, getClient } from "svelte-apollo";

    // Used for SSR. A falsy value is ignored by the Router.
    export let url = "";

    const link = new HttpLink({uri: "http://192.168.8.108:8085/query"});
    const client = new ApolloClient({link, cache: new InMemoryCache()});
    setClient(client);
  </script>
  
<Router url="{url}">
  <nav>
    <NavLink to="/">Home</NavLink> -
    <NavLink to="country">Country</NavLink> -
    <NavLink to="region/asia">Region</NavLink> -
    <NavLink to="worldwide">Worldwide</NavLink>
  </nav>
  <div>
    <Route path="/" component="{Country}" />
    <Route path="/region/:region" component="{Region}" />
    <Route path="/country/:countryId" component="{Country}" />
    <Route path="/country" component="{Country}" />
    <Route path="/worldwide" component="{Worldwide}" />
  </div>
</Router>


<style>
	:global(main) {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
    max-width: 1280px;
	}

	:global(h1) {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

  :global(h2) {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 2em;
		font-weight: 100;
  }
</style>