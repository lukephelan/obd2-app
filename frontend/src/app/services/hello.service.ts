import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

interface HelloResponse {
  message: string;
}

@Injectable({
  providedIn: 'root',
})
export class HelloService {
  private readonly apiUrl = 'http://localhost:8080/api';

  constructor(private http: HttpClient) {}

  getHello(): Observable<HelloResponse> {
    return this.http.get<HelloResponse>(`${this.apiUrl}/hello`);
  }
}
