<script lang="ts">
  import KillEvent from "../models/kill_event";

  let { rounds, teams, selectedEvent = $bindable() } = $props();

  function onclick(event) {
    const roundIndex = event.target.dataset.roundIndex;
    const eventIndex = event.target.dataset.eventIndex;
    selectedEvent = new KillEvent(rounds[roundIndex].killFeed[eventIndex]);
  }
</script>

{#each rounds as round, roundIndex}
  <div class="row">
    <div class="col">
      <div class="card">
        <div class="card-header">
          Round {round.roundNumber} | {round.winnerState.teamName(teams)} Wins |
          {round.teamScoreToString(teams)}
        </div>
        <div class="card-body">
          <ul>
            {#each round.killFeed as killEvent, eventIndex}
              <li>
                <pre
                  {onclick}
                  data-event-index={eventIndex}
                  data-round-index={roundIndex}>
                {JSON.stringify(killEvent)}
                </pre>
              </li>
            {/each}
          </ul>
        </div>
      </div>
    </div>
  </div>
{/each}
