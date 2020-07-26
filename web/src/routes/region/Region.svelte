
<script lang="ts">
	import { gql } from "apollo-boost";
	import { query, getClient } from "svelte-apollo";
	import type { Query } from "../../generated/graphql";
	import Records from "./Records.svelte";
	import { navigate } from "svelte-routing";
	import Select from 'svelte-select';

	const client = getClient();

	const REGIONS = gql`
	{
		getAllRegions
	}`;
	const regionList = query<Query>(client, { query: REGIONS });

	const updateRecordsObserver = (isInit?: boolean) => {
		const RECORD = gql`
		{
			getRecordsByRegion(region: "${region}"){
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

	$: if (region) {
		updateRecordsObserver();
		records.refetch();
	}
</script>

<main>
	<h1>Covid 19 - {region}</h1>
	{#await $regionList}
		<p>Loading regions....</p>
	{:then result}
		<Select
			containerStyles={
				"text-transform: capitalize"
			}
			selectedValue={region}
			on:select={(e) => {
				navigate("/region/" + e.detail.value.toLowerCase());
			}}
			items={result.data.getAllRegions}
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
