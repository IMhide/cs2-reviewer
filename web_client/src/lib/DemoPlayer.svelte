<script lang="ts">
  import { onMount } from "svelte";

  let { demoInfo, selectedEvent = $bindable() } = $props();
  let svg;

  let victim = $derived(selectedEvent.victim);
  let attacker = $derived(selectedEvent.attacker);
  let assister = $derived(selectedEvent.assister);
  let mapMetaData;

  onMount(() => {
    const BACKEND_URL = "http://localhost:4567/";
    const mapUrl = BACKEND_URL + "maps/" + demoInfo.mapName + ".png";
    const image = document.createElementNS(
      "http://www.w3.org/2000/svg",
      "image",
    );

    image.setAttribute("href", mapUrl);
    image.setAttribute("width", svg.getAttribute("width"));
    image.setAttribute("height", svg.getAttribute("height"));
    svg.appendChild(image);
    mapMetaData = metaData(demoInfo.mapName);
  });

  $effect(() => {
    if (selectedEvent === undefined || mapMetaData === undefined) return;
    svg.appendChild(
      createCircle(
        victim.positionX,
        victim.positionY,
        victim.positionZ,
        "red",
        mapMetaData,
      ),
    );
    svg.appendChild(
      createCircle(
        attacker.positionX,
        attacker.positionY,
        attacker.positionZ,
        "blue",
        mapMetaData,
      ),
    );
    if (assister.playerID != 0) {
      svg.appendChild(
        createCircle(
          assister.positionX,
          assister.positionY,
          assister.positionZ,
          "green",
          mapMetaData,
        ),
      );
    }
  });

  function createCircle(x, y, z, color, mapMetaData) {
    const circle = document.createElementNS(
      "http://www.w3.org/2000/svg",
      "circle",
    );
    let offsetY = mapMetaData.offsetY;
    let offsetX = mapMetaData.offsetX;
    const ratio = mapMetaData.ratio;

    if (mapMetaData.zSplit && z < mapMetaData.zSplit) {
      offsetX = mapMetaData.splitOffsetX;
      offsetY = mapMetaData.splitOffsetY;
    }

    const posX: number = ((parseFloat(offsetX) + x) / ratio).toFixed(5);
    const posY: number = ((parseFloat(offsetY) - y) / ratio).toFixed(5);

    console.log({ color: color, x: x, y: y, z: z, posX: posX, posY: posY });
    circle.setAttributeNS(null, "cx", posX);
    circle.setAttributeNS(null, "cy", posY);
    circle.setAttributeNS(null, "r", 5);
    circle.setAttributeNS(
      null,
      "style",
      "fill: " + color + "; stroke: white; stroke-width: 1px;",
    );
    return circle;
  }

  function metaData(map) {
    return [
      { map: "de_anubis", offsetX: 2830, offsetY: 3330, ratio: 5.25 },
      { map: "de_dust2", offsetX: 2475, offsetY: 3250, ratio: 4.4 },
      { map: "de_ancient", offsetX: 2590, offsetY: 1850, ratio: 4.3 },
      { map: "de_inferno", offsetX: 2090, offsetY: 3900, ratio: 4.9 },
      { map: "de_mirage", offsetX: 3250, offsetY: 1700, ratio: 5.0 },
      { map: "de_overpass", offsetX: 4830, offsetY: 1800, ratio: 5.2 },
      {
        map: "de_vertigo",
        offsetX: 3950,
        offsetY: 1300,
        ratio: 5.0,
        zSplit: 11680,
        splitOffsetX: 3950,
        splitOffsetY: 3490,
      },
      {
        map: "de_nuke",
        offsetX: 3300,
        offsetY: 1200,
        ratio: 7.0,
        zSplit: -480,
        splitOffsetX: 3300,
        splitOffsetY: 4400,
      },
      //   {map: 'de_train', offsetX: 2830, offsetY: 3330, ratio: 5.25},
    ].find(function (element) {
      if (map == element.map) {
        return element;
      }
    });
  }
</script>

<svg
  bind:this={svg}
  id="mapSvg"
  class="h-100"
  width="100%"
  height="100%"
  viewBox="0 0 1024 1024"
></svg>
