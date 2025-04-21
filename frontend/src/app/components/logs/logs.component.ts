import { Component } from '@angular/core';
import { AppButtonComponent } from '../app-button/app-button.component';

@Component({
  selector: 'app-logs',
  imports: [AppButtonComponent],
  templateUrl: './logs.component.html',
  styleUrl: './logs.component.css',
})
export class LogsComponent {}
