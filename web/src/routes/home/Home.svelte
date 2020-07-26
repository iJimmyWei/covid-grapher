<script>
    import NavLink from "../../components/NavLink.svelte";
</script>

<main>
    <section class={"introduction"}>
        <h1>
            Covid 19
        </h1>
        <p>
            On February 11, 2020 the World Health Organization announced an official name for the disease that is causing the 2019 novel coronavirus outbreak, first identified in Wuhan China. The new name of this disease is coronavirus disease 2019, abbreviated as COVID-19. In COVID-19, ‘CO’ stands for ‘corona,’ ‘VI’ for ‘virus,’ and ‘D’ for disease. Formerly, this disease was referred to as “2019 novel coronavirus” or “2019-nCoV”.
        </p>
        <p>
            There are many types of human coronaviruses including some that commonly cause mild upper-respiratory tract illnesses. COVID-19 is a new disease, caused by a novel (or new) coronavirus that has not previously been seen in humans.
        </p>
        <p class={"rightAlign"}>
            - CDC (https://www.cdc.gov)
        </p>
    </section>
    <h2>
        Graph Visualization
    </h2>
    <section class={"buttons"}>
        <NavLink to="country"><button>Search by Country</button></NavLink>
        <NavLink to="region/asia"><button>Search by Region</button></NavLink>
        <NavLink to="worldwide"><button>View Worldwide Trend</button></NavLink>
    </section>
</main>

<style>
    .introduction {
        padding-bottom: 32px;
    }
    .rightAlign{
        text-align: right;
    }
    .buttons {
        display: flex;
        flex-direction: column;
        max-width: 240px;
        margin: 0 auto;
    }
</style>