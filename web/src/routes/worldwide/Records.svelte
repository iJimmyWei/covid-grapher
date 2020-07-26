<script lang="ts">
  import type { Query, Record } from "../../generated/graphql";
  import Line from "svelte-chartjs/src/Line.svelte";
  export let data: Pick<Query, "getRecords">;

  let allCountries: Record[][] = [];

  let country = [];

  // O(N) fastest way to group by country.. relies on data returned back in order of country and date
  data.getRecords.forEach((d, index) => {
    const isLast = data.getRecords.length - 1 === index;
    if (isLast) {
        country.push(d);

        country.reverse();
        allCountries.push(country);
        return;
    }

    const nextItem = data.getRecords[index + 1];
    const nextItemIsDiffCountry =
        nextItem.countriesAndTerritories !== d.countriesAndTerritories;
    if (nextItemIsDiffCountry) {
        country.push(d);

        country.reverse();
        allCountries.push(country);
        country = [];
        return;
    }

    country.push(d);
  });

  let options = {
    legend: {
      display: false,
    },
    hover: {
        animationDuration: 0,
    },
    tooltips: {
        enabled: true,
        mode: 'single',
        callbacks: {
            label: function(tooltipItems) {
                const index = tooltipItems.datasetIndex; 
                const country = allCountries[index][0].countriesAndTerritories.replace(/_/g, " ");

                return tooltipItems.yLabel + ' : ' + country;
            }
        }
    },
  };

  // Hide large outliers
  allCountries = allCountries.filter((c) => !c.find((d) => Number.parseFloat(d.cumulative_14d_per_10000) > 500 || Number.parseFloat(d.cumulative_14d_per_10000) < 0));

  let dataLine = {
    labels: allCountries[0].map((d) => d.dateRep),
    datasets: allCountries.map((c) => ({
        fill: false,
        lineTension: 0,
        backgroundColor: "rgba(225, 204,230, .3)",
        borderColor: "rgba(0, 0, 0, .36)",
        borderCapStyle: "square",
        borderDash: [],
        borderDashOffset: 0.0,
        borderWidth: 1,
        borderJoinStyle: "round",
        pointBorderWidth: 0,
        pointHoverRadius: 1,
        pointHoverBackgroundColor: "rgb(0, 0, 0)",
        pointHoverBorderColor: "rgb(205, 130,1 58)",
        pointHoverBorderWidth: 0,
        pointRadius: 0,
        pointHitRadius: 5,
        data: c.map((d) => d.cumulative_14d_per_10000),
    })),
  };
</script>

<main>
  <Line data={dataLine} {options} />
</main>
