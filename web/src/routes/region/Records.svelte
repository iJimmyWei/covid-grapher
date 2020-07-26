<script lang="ts">
    import type { Query } from "../../generated/graphql";
    import Line from "svelte-chartjs/src/Line.svelte"
    export let data: Pick<Query, "getRecordsByRegion">;

    const sortedData = data.getRecordsByRegion
        .filter((c) => c.cases >= 0 && c.deaths >= 0); // Strip away outliers

    let dataLine = {
    labels: data.getRecordsByRegion.map((d) => d.dateRep),
    datasets: [
      {
        label: "Number of Daily Deaths",
        fill: false,
        lineTension: 0.1,
        backgroundColor: "rgba(225, 204,230, .3)",
        borderColor: "rgb(205, 130, 158)",
        borderCapStyle: "butt",
        borderDash: [],
        borderDashOffset: 0.0,
        borderJoinStyle: "miter",
        pointBorderColor: "rgb(205, 130,1 58)",
        pointBackgroundColor: "rgb(255, 255, 255)",
        pointBorderWidth: 5,
        pointHoverRadius: 1,
        pointHoverBackgroundColor: "rgb(0, 0, 0)",
        pointHoverBorderColor: "rgb(205, 130,1 58)",
        pointHoverBorderWidth: 0.5,
        pointRadius: 1,
        pointHitRadius: 10,
        data: sortedData.map((d) => d.deaths)
      },
      {
        label: "Number of Daily Cases",
        fill: false,
        lineTension: 0.1,
        backgroundColor: "rgba(71, 225, 167, 0.5)",
        borderColor: "rgb(71, 225, 167)",
        borderCapStyle: "butt",
        borderDash: [],
        borderDashOffset: 0.0,
        borderJoinStyle: "miter",
        pointBorderColor: "rgb(205, 130,1 58)",
        pointBackgroundColor: "rgb(255, 255, 255)",
        pointBorderWidth: 5,
        pointHoverRadius: 1,
        pointHoverBackgroundColor: "rgb(0, 0, 0)",
        pointHoverBorderColor: "rgb(205, 130,1 58)",
        pointHoverBorderWidth: 0.5,
        pointRadius: 1,
        pointHitRadius: 10,
        data: sortedData.map((d) => d.cases)
      },
    ]
  };
</script>

<main>
    <Line data={dataLine} />
</main>