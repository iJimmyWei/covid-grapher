<script lang="ts">
    import type { Query } from "./types";

    export let data: Query;
</script>

<main>
    {#each data.recordsByCountryCode as record (record.id)}
        <p>{record.dateRep} - {record.deaths} deaths</p>
    {/each}
</main>