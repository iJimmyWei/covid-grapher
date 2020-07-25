<script lang="ts">
    import type { Query } from "./generated/graphql";
    import Line from "svelte-chartjs/src/Line.svelte"
    export let data: Pick<Query, "recordsByCountryName">;

    const sortedData = data.recordsByCountryName.map((e) => e).sort((a, b) => a > b ? 1 : 0).reverse();

    let dataLine = {
    labels: sortedData.map((d) => d.dateRep),
    datasets: [
      {
        label: "Number of Deaths",
        fill: true,
        lineTension: 0.3,
        backgroundColor: "rgba(225, 204,230, .3)",
        borderColor: "rgb(205, 130, 158)",
        borderCapStyle: "butt",
        borderDash: [],
        borderDashOffset: 0.0,
        borderJoinStyle: "miter",
        pointBorderColor: "rgb(205, 130,1 58)",
        pointBackgroundColor: "rgb(255, 255, 255)",
        pointBorderWidth: 10,
        pointHoverRadius: 1,
        pointHoverBackgroundColor: "rgb(0, 0, 0)",
        pointHoverBorderColor: "rgb(205, 130,1 58)",
        pointHoverBorderWidth: 1,
        pointRadius: 1,
        pointHitRadius: 10,
        data: sortedData.map((d) => d.deaths)
      },
    ]
  };
</script>

<main>
    <!-- {#each data.recordsByCountryName as record (record.id)}
        <p>{record.dateRep} - {record.deaths} deaths</p>
    {/each} -->

    <Line data={dataLine} />
</main>