import { Injectable, signal } from '@angular/core';

export type Adapter = 'USB' | 'Bluetooth' | 'Wi-Fi';
export type Units = 'metric' | 'imperial';

@Injectable({ providedIn: 'root' })
export class SettingsService {
  private adapter = signal<Adapter>('USB');
  private pollingRate = signal(1000);
  private units = signal<Units>('metric');
  private connected = signal(false);

  readonly adapter$ = this.adapter.asReadonly();
  readonly pollingRate$ = this.pollingRate.asReadonly();
  readonly units$ = this.units.asReadonly();
  readonly connected$ = this.connected.asReadonly();

  setPollingRate(rate: number) {
    this.pollingRate.set(rate);
  }

  setUnits(units: Units) {
    this.units.set(units);
  }

  setConnected(connected: boolean) {
    this.connected.set(connected);
  }

  setAdapter(adapter: Adapter) {
    this.adapter.set(adapter);
  }
}
