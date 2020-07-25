
<script lang="ts">
	import ApolloClient from "apollo-boost";
	import { setClient, query } from "svelte-apollo";

	export let name: string;

	const client: any = new ApolloClient({uri: "localhost:8085/query"});
	setClient(client);

	const books = query(client, { query: GET_BOOKS });
</script>

<main>
	<h1>derp</h1>
	<p>Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn how to build Svelte apps.</p>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>