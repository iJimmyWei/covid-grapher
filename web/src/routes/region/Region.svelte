
<script lang="ts">
	import { ApolloClient, gql, InMemoryCache, HttpLink } from "apollo-boost";
	import { setClient, query, getClient } from "svelte-apollo";
	import type { Query } from "../../generated/graphql";
	import Records from "./Records.svelte";
	import { navigate } from "svelte-routing";

	const link = new HttpLink({uri: "http://192.168.8.108:8085/query"});
	const client = new ApolloClient({link, cache: new InMemoryCache()});
	setClient(client);

	const updateRecordsObserver = (isInit?: boolean) => {
		const RECORD = gql`
		{
			getRecords(region: "${region}"){
				id,
				dateRep,
				deaths,
				cases
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
		
	export let region = "Asia";
	let records = updateRecordsObserver(true);
</script>

<main>
	<h1>Covid 19 - {region}</h1>
	{#await $records}
		<p>Loading records....</p>
	{:then result}
		{console.log(result)}
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

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>