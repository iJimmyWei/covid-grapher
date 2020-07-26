
<script lang="ts">
	import { gql } from "apollo-boost";
	import { query, getClient } from "svelte-apollo";
	import type { Query } from "../../generated/graphql";
	import Records from "./Records.svelte";
	import Select from 'svelte-select';
	import { navigate } from "svelte-routing";

	const client = getClient();
	
	const COUNTRIES = gql`
	{
		getAllCountries
	}`;
	const countryList = query<Query>(client, { query: COUNTRIES });

	const updateRecordsObserver = (isInit?: boolean) => {
		const RECORD = gql`
		{
			getRecords(countryName: "${countryId}"){
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
		
	export let countryId = "Japan"
	let countryName = countryId.replace(/_/g, " ")
	let records = updateRecordsObserver(true);

	$: if (countryId) {
		updateRecordsObserver();
		records.refetch();

		countryName = countryId.replace(/_/g, " ")
	};
</script>

<main>
	<h1>Covid 19 - {countryName}</h1>
	{#await $countryList}
		<p>Loading countries....</p>
	{:then result}
		<Select
			containerStyles={
				"text-transform: capitalize"
			}
			selectedValue={{
				"value": countryId,
				"label": countryName
			}}
			on:select={(e) => {
				navigate("/country/" + e.detail.value.toLowerCase());
			}}
			items={result.data.getAllCountries.map((c) =>
					({
						"value": c,
						"label": c.replace(/_/g, " ")
					})
				)
			}
		></Select>
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
