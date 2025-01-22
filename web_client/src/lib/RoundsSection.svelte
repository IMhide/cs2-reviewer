<script lang="ts">
  let { rounds, teams } = $props();
</script>

{#each rounds as round}
  <div class="row">
    <div class="col">
      <div class="card">
        <div class="card-header">
          Round {round.roundNumber} | {round.winnerState.teamName(teams)} Wins |
          {round.teamScoreToString(teams)}
        </div>
        <div class="card-body">
          <ul>
            {#each round.killFeed as killEvent}
              <li>
                <pre>{JSON.stringify(killEvent)}</pre>
              </li>
            {/each}
          </ul>
        </div>
      </div>
    </div>
  </div>
{/each}
