<script lang="ts">
  import DemoPlayer from "./lib/DemoPlayer.svelte";
  import DemoInfoSection from "./lib/DemoInfoSection.svelte";
  import TeamsSection from "./lib/TeamsSection.svelte";
  import RoundsSection from "./lib/RoundsSection.svelte";
  import DemoInfo from "./models/demo_info";
  import Team from "./models/team";
  import Round from "./models/round";

  const host = "http://localhost:4567";

  let informations = $state();
  let demoInfo = $state(new DemoInfo({}));
  let teams = $state([]);
  let rounds = $state([]);
  let selectedEvent = $state();

  $effect(() => {
    fetch(host + "/test")
      .then((res) => res.json())
      .then((data) => {
        informations = data;
        demoInfo = new DemoInfo(data.demoInfo);
        teams = data.teams.map((team) => Team.fromPayload(team));
        rounds = data.rounds.map((round) => new Round(round));
      })
      .catch((err) => {
        console.log(err);
      });
  });

  $inspect(teams);
</script>

<div class="row h-100">
  <div class="col-8 h-100 text-center">
    {#key demoInfo}
      <DemoPlayer {demoInfo} bind:selectedEvent />
    {/key}
  </div>

  <div class="col-4 overflow-scroll vh-100 pb-5">
    <DemoInfoSection {demoInfo} />
    <TeamsSection {teams} />
    <RoundsSection {rounds} {teams} bind:selectedEvent />
  </div>
</div>
