<script lang="ts">
    import type { Query } from "./generated/graphql";

    export let data: Query;
</script>

<main>
    {#each data.recordsByCountryName as record (record.id)}
        <p>{record.dateRep} - {record.deaths} deaths</p>
    {/each}
</main>