import { NgClass } from '@angular/common';
import { Component, computed, input } from '@angular/core';

@Component({
  selector: 'app-app-button',
  imports: [NgClass],
  templateUrl: './app-button.component.html',
  styleUrl: './app-button.component.css',
})
export class AppButtonComponent {
  color = input<string>('blue');

  buttonClass = computed(() => {
    return {
      'bg-blue-600': this.color() === 'blue',
      'bg-red-600': this.color() === 'red',
    };
  });
}
