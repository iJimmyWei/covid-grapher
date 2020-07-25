
<script lang="ts">
	import { ApolloClient, gql, InMemoryCache, HttpLink } from "apollo-boost";
	import { setClient, query, getClient } from "svelte-apollo";
	import type { Query, QueryRecordsByCountryCodeArgs } from "./types";
	import Records from "./Records.svelte";
import type { QueryStore } from "apollo-client/data/queries";

	let countryCode = "AFG"

	const link = new HttpLink({uri: "http://localhost:8085/query"});
	const client = new ApolloClient({link, cache: new InMemoryCache()});
	setClient(client);
	
	const updateRecordsObserver = (isInit?: boolean) => {
		const RECORD = gql`
		{
			recordsByCountryCode(countryCode: "${countryCode}"){
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

	let records = updateRecordsObserver(true);

	$: if (countryCode) {
		updateRecordsObserver();
		records.refetch();
	};
</script>

<main>
	<h1>Covid 19 - {countryCode}</h1>
	<input bind:value={countryCode}>
	{#await $records}
		<p>Loading records....</p>
	{:then result}
		<Records data={result.data} />
	{:catch error}
		<p>Error loading records: {error}</p>
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