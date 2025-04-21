import { Routes } from '@angular/router';
import { LiveDataComponent } from './components/live-data/live-data.component';
import { DiagnosticsComponent } from './components/diagnostics/diagnostics.component';
import { VehicleInfoComponent } from './components/vehicle-info/vehicle-info.component';
import { TestsComponent } from './components/tests/tests.component';
import { LogsComponent } from './components/logs/logs.component';
import { SettingsComponent } from './components/settings/settings.component';
import { HomeComponent } from './components/home/home.component';

export const routes: Routes = [
  { path: '', redirectTo: 'home', pathMatch: 'full' },
  { path: 'home', component: HomeComponent },
  { path: 'live', component: LiveDataComponent },
  { path: 'diagnostics', component: DiagnosticsComponent },
  { path: 'vehicle', component: VehicleInfoComponent },
  { path: 'tests', component: TestsComponent },
  { path: 'logs', component: LogsComponent },
  { path: 'settings', component: SettingsComponent },
];
