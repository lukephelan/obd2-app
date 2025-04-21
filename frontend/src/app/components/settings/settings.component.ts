import { Component } from '@angular/core';
import { AppButtonComponent } from '../app-button/app-button.component';
import { SettingsService } from '../../services/settings.service';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-settings',
  imports: [AppButtonComponent, FormsModule, CommonModule],
  templateUrl: './settings.component.html',
  styleUrl: './settings.component.css',
})
export class SettingsComponent {
  constructor(public settings: SettingsService) {}
}
