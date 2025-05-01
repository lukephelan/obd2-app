import { Component } from '@angular/core';
import { AppLayoutComponent } from './components/app-layout/app-layout.component';
import { HelloService } from './services/hello.service';
import { HttpClientModule } from '@angular/common/http';

@Component({
  selector: 'app-root',
  imports: [AppLayoutComponent, HttpClientModule],
  providers: [HelloService],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css',
})
export class AppComponent {
  message: string = '';

  constructor(private helloService: HelloService) {}

  ngOnInit() {
    this.helloService.getHello().subscribe(response => {
      this.message = response.message;
      console.log(response);
    });
  }
}
