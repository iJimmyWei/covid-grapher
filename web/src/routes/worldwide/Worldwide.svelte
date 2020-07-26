
<script lang="ts">
	import { ApolloClient, gql, InMemoryCache, HttpLink } from "apollo-boost";
	import { setClient, query, getClient } from "svelte-apollo";
	import type { Query } from "../../generated/graphql";
	import Records from "./Records.svelte";
	import Select from 'svelte-select';

	const link = new HttpLink({uri: "http://0.0.0.0:8085/query"});
	const client = new ApolloClient({link, cache: new InMemoryCache()});
	setClient(client);

	const updateRecordsObserver = (isInit?: boolean) => {
		const RECORD = gql`
		{
			getRecords{
				countriesAndTerritories,
				cumulative_14d_per_10000,
				dateRep
			}
		}`;

		const observer = query<Query>(client, { query: RECORD });
		if (isInit) {
			return observer;
		}
		
		if (records) {
			records = observer;
		}
	}
		
	let records = updateRecordsObserver(true);
</script>

<main>
	<h1>Covid 19 - Worldwide</h1>
	<h2>Cumulative Cases over 14 days, per 10000</h2>
	{#await $records}
		<p>Loading worldwide data....</p>
	{:then result}
		<Records data={result.data} />
	{:catch error}
		<p>{error}</p>
	{/await}
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

	h2 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 2em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>