window.addEventListener("load", function () {
  console.log("canvas-renderer.js");
  const svg = document.getElementById("mapSvg");
  const image = this.document.createElementNS(
    "http://www.w3.org/2000/svg",
    "image",
  );
  const mapName = svg.getAttribute("data-map");

  image.setAttribute("href", "maps/" + mapName + ".png");
  image.setAttribute("width", svg.getAttribute("width"));
  image.setAttribute("height", svg.getAttribute("height"));
  svg.appendChild(image);

  interactiveCanvas(svg, metaData(mapName));
});

function interactiveCanvas(svg, mapMetaData) {
  console.log(mapMetaData);
  const elements = document.getElementsByClassName("killEvent");
  let clickedKillEvent = null;

  Array.from(elements).forEach(function (element) {
    element.addEventListener("click", function (event) {
      if (clickedKillEvent) {
        clickedKillEvent.style.backgroundColor = "transparent";
      }
      clickedKillEvent = event.target;
      clickedKillEvent.style.backgroundColor = "lightblue";
      data = JSON.parse(clickedKillEvent.innerText);
      console.log(data);
      let victim = data.victim;
      let attacker = data.attacker;
      let assister = data.assister;
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
      console.log(attacker);
      if (assister.attackerID != 0) {
        svg.appendChild(
          createCircle(
            assister.positionX,
            assister.positionY,
            assister.positionZ,
            "green",
            mapMetaData,
          ),
        );
        console.log(assister);
      }
      console.log(data);
    });
  });
}

function createCircle(x, y, z, color, mapMetaData) {
  const circle = this.document.createElementNS(
    "http://www.w3.org/2000/svg",
    "circle",
  );
  offsetY = mapMetaData.offsetY;
  offsetX = mapMetaData.offsetX;
  ratio = mapMetaData.ratio;

  if (mapMetaData.zSplit && z < mapMetaData.zSplit) {
    offsetX = mapMetaData.splitOffsetX;
    offsetY = mapMetaData.splitOffsetY;
  }

  posX = (offsetX + x) / ratio;
  posY = (offsetY - y) / ratio;

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
