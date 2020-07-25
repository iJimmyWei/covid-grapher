
<script lang="ts">
	import { ApolloClient, gql, InMemoryCache, HttpLink } from "apollo-boost";
	import { setClient, query, getClient } from "svelte-apollo";
	import type { Query } from "./generated/graphql";
	import Records from "./Records.svelte";
	import Select from 'svelte-select';

	const link = new HttpLink({uri: "http://localhost:8085/query"});
	const client = new ApolloClient({link, cache: new InMemoryCache()});
	setClient(client);
	
	const COUNTRIES = gql`
	{
		getAllCountries
	}`;
	const countryList = query<Query>(client, { query: COUNTRIES });

	const updateRecordsObserver = (isInit?: boolean) => {
		const RECORD = gql`
		{
			recordsByCountryName(countryName: "${countryName}"){
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
		
	let countryName = "Japan"
	let records = updateRecordsObserver(true);
		
	$: if (countryName) {
		updateRecordsObserver();
		records.refetch();
	};
</script>

<main>
	<h1>Covid 19 - {countryName}</h1>
	{#await $countryList}
		<p>Loading countries....</p>
	{:then result}
		<Select
			on:select={(e) => {
				countryName = e.detail.label;
			}}
			items={result.data.getAllCountries.map((c) =>
				({
					"value": c,
					"label": c.replace(/_/g, " ")
				}))}
		/>
	{:catch error}
		<p>Error loading countries: {error}</p>
	{/await}

	{#await $records}
		<p>Loading records....</p>
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

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>