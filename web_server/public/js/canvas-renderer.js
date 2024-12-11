window.addEventListener('load', function () {
  console.log('canvas-renderer.js');
  const svg = document.getElementById('mapSvg');
  const image = this.document.createElementNS('http://www.w3.org/2000/svg', 'image');
  const mapName = svg.getAttribute('data-map');

  image.setAttribute('href', 'maps/' + mapName + '.png');
  image.setAttribute('width', svg.getAttribute('width'));
  image.setAttribute('height', svg.getAttribute('height'));
  svg.appendChild(image);

  interactiveCanvas(svg);
})


function interactiveCanvas(svg) {
  const elements = document.getElementsByClassName("killEvent");
  let clickedKillEvent = null;

  Array.from(elements).forEach(function(element) {
    element.addEventListener('click', function(event) {
      if (clickedKillEvent) {
        clickedKillEvent.style.backgroundColor = 'transparent';
      }
      clickedKillEvent = event.target;
      clickedKillEvent.style.backgroundColor = 'blue';
      data = JSON.parse(clickedKillEvent.innerText);
      let victim = data.victim;
      let attacker = data.attacker;
      let assister = data.assister;
      svg.appendChild(createCircle(victim.positionX, victim.positionY, 'red'));
      console.log(victim);
      svg.appendChild(createCircle(attacker.positionX, attacker.positionY, 'blue'));
      console.log(attacker);
      if (assister.attackerID != 0) { 
      svg.appendChild(createCircle(assister.positionX, assister.positionY, 'green'));
        console.log(assister);
      }
      console.log(data);
    });
  });
}

function createCircle(x,y, color) {
  const circle= this.document.createElementNS('http://www.w3.org/2000/svg', 'circle');
  offsetY = 3330;
  offsetX = 2830;
  posX = (offsetX + x) / 5.25;
  posY = (offsetY - y) / 5.25;

  circle.setAttributeNS(null, 'cx', posX);
  circle.setAttributeNS(null, 'cy', posY);
  circle.setAttributeNS(null, 'r', 5);
  circle.setAttributeNS(null, 'style', 'fill: ' + color + '; stroke: white; stroke-width: 1px;' );
  return circle;
}
