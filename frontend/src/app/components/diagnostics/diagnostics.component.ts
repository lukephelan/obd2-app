import { Component } from '@angular/core';
import { OBD2Service } from '../../services/obd2.service';
import { AppButtonComponent } from '../app-button/app-button.component';

@Component({
  selector: 'app-diagnostics',
  imports: [AppButtonComponent],
  templateUrl: './diagnostics.component.html',
  styleUrl: './diagnostics.component.css',
})
export class DiagnosticsComponent {
  constructor(public obd: OBD2Service) {}
}
