<script lang="ts">
  import Spinner from "./Spinner.svelte";

  let { teams } = $props();
</script>

<div class="row">
  {#if teams && teams.length > 0}
    <div class="col">
      <div class="card">
        <div class="card-header">
          {teams[0].name}
        </div>
        <div class="card-body">
          <ul>
            {#each teams[0].players as player}
              <li>{player.name}</li>
            {/each}
          </ul>
        </div>
      </div>
    </div>

    <div class="col">
      <div class="card">
        <div class="card-header">
          {teams[1].name}
        </div>
        <div class="card-body">
          <ul>
            {#each teams[1].players as player}
              <li>{player.name}</li>
            {/each}
          </ul>
        </div>
      </div>
    </div>
  {:else}
    <Spinner />
  {/if}
</div>
