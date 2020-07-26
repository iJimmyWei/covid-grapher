
<script lang="ts">
	import { gql } from "apollo-boost";
	import { query, getClient } from "svelte-apollo";
	import type { Query } from "../../generated/graphql";
	import Records from "./Records.svelte";

	const client = getClient();

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
