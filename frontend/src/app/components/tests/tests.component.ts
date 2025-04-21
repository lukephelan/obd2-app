import { Component } from '@angular/core';
import { AppButtonComponent } from '../app-button/app-button.component';

@Component({
  selector: 'app-tests',
  imports: [AppButtonComponent],
  templateUrl: './tests.component.html',
  styleUrl: './tests.component.css',
})
export class TestsComponent {}
