import { Component } from '@angular/core';
import { OBD2Service } from '../../services/obd2.service';
import { DecimalPipe } from '@angular/common';

@Component({
  selector: 'app-live-data',
  imports: [DecimalPipe],
  templateUrl: './live-data.component.html',
  styleUrl: './live-data.component.css',
})
export class LiveDataComponent {
  constructor(public obd: OBD2Service) {}
}
