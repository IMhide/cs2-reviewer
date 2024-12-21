<script lang="ts">
  let {teams} = $props()
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
    <div class="spinner-border" role="status">
      <span class="visually-hidden">Loading...</span>
    </div>
  {/if}
</div>

