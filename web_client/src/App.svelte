<script lang="ts">
  import DemoInfo from './models/demo_info'
  import DemoInfoSection from './lib/DemoInfoSection.svelte'
  import Team from './models/team'
  import Round from './models/round'


  const host = 'http://localhost:4567'

  let informations = $state()
  let demoInfo = $state(new DemoInfo({}))
  let teams = $state([])
  let rounds = $state([])

  $effect(() => {
    fetch(host + '/test')
      .then(res => res.json())
      .then(data => {
        informations = data
        demoInfo = new DemoInfo(data.demoInfo)
        teams = data.teams.map(team => new Team(team))
        rounds = data.rounds.map(round => new Round(round))
      }).catch(err => {
        console.log(err)
      })
  })
  $inspect(informations)
  $inspect(rounds)
</script>

<main>
  <div class="row h-100">
    <div class="col-8 h-100 text-center">
      <svg id="mapSvg" class="h-100" width="100%" height="100%" viewBox="0 0 1024 1024" data-map={demoInfo.mapName}></svg>
    </div>

    <div class="col-4 overflow-scroll vh-100 pb-5">
      <DemoInfoSection {demoInfo} />
    </div>
</main>
