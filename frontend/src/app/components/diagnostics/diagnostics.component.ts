import { Component } from '@angular/core';
import { OBD2Service } from '../../services/obd2.service';

@Component({
  selector: 'app-diagnostics',
  imports: [],
  templateUrl: './diagnostics.component.html',
  styleUrl: './diagnostics.component.css',
})
export class DiagnosticsComponent {
  constructor(public obd: OBD2Service) {}
}
