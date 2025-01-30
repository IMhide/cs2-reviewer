<script lang="ts">
  import Spinner from "./Spinner.svelte";

  let { demoInfo } = $props();
</script>

<div class="row">
  <div class="col">
    {#if demoInfo === undefined}
      <Spinner />
    {:else}
      <h2>{demoInfo.mapName} - {demoInfo.serverName}</h2>
      <div class="row">
        <div class="col">
          <div class="card">
            <div class="card-body">
              <p>
                Duration in Seconds: {demoInfo.durationInSec}, Ticks: {demoInfo.durationInTick},
                Frames: {demoInfo.durationInFrame}
              </p>
            </div>
          </div>
        </div>
      </div>
    {/if}
  </div>
</div>
