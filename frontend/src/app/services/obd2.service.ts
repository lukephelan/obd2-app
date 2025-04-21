import { Injectable, signal } from '@angular/core';

@Injectable({ providedIn: 'root' })
export class OBD2Service {
  rpm = signal(0);
  speed = signal(0);
  voltage = signal(0);
  coolantTemp = signal(0);
  throttle = signal(0);

  constructor() {
    // simulate updates
    setInterval(() => {
      this.rpm.set(Math.floor(700 + Math.random() * 3000));
      this.speed.set(Math.floor(Math.random() * 120));
      this.voltage.set(12 + Math.random());
      this.coolantTemp.set(85 + Math.random() * 15);
      this.throttle.set(Math.floor(Math.random() * 100));
    }, 1000);
  }
}
