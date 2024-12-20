export default class DemoInfo {
  mapName: string;
  serverName: string;
  durationInSec: number;
  durationInTick: number;
  durationInFrame: number;

  constructor(data: Object) {
    this.mapName = data.mapName;
    this.serverName = data.serverName;
    this.durationInSec = data.durationInSec;
    this.durationInTick = data.durationInTick;
    this.durationInFrame = data.durationInFrame;
  }
}
