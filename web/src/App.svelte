<script>
    import { Router, Route } from "svelte-routing";
    import NavLink from "./components/NavLink.svelte";
    import Country from "./routes/country/Country.svelte";
    import Region from "./routes/region/Region.svelte";
    import Worldwide from "./routes/worldwide/Worldwide.svelte";
  
    // Used for SSR. A falsy value is ignored by the Router.
    export let url = "";
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