<script lang="ts">
  import Spinner from "./lib/Spinner.svelte";
  import DemoPlayer from "./lib/DemoPlayer.svelte";
  import DemoInfoSection from "./lib/DemoInfoSection.svelte";
  import TeamsSection from "./lib/TeamsSection.svelte";
  import RoundsSection from "./lib/RoundsSection.svelte";
  import DemoInfo from "./models/demo_info";
  import Team from "./models/team";
  import Round from "./models/round";

  const host = "http://localhost:4567";

  let informations = $state();
  let demoInfo = $state();
  let teams = $state([]);
  let rounds = $state([]);
  let selectedEvent = $state();

  $effect(() => {
    fetch(host + "/test")
      .then((res) => res.json())
      .then((data) => {
        demoInfo = DemoInfo.parse(data.demoInfo);
        informations = data;
        teams = data.teams.map((team: any) => Team.parse(team));
        rounds = data.rounds.map((round: any) => {
          console.log(round);
          return Round.parse(round);
        });
      })
      .catch((err) => {
        console.log(err);
      });
  });
</script>

<div class="row h-100">
  <div class="col-8 h-100 text-center">
    {#if demoInfo === undefined}
      <Spinner />
    {:else}
      <DemoPlayer {demoInfo} {selectedEvent} />
    {/if}
  </div>

  <div class="col-4 overflow-scroll vh-100 pb-5">
    <DemoInfoSection {demoInfo} />
    <TeamsSection {teams} />
    <RoundsSection {rounds} {teams} bind:selectedEvent />
  </div>
</div>
