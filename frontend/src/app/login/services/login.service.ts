import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable, Output } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { LoginResponseDTO } from 'src/app/models/loginDTO';
import { EventEmitter } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  @Output() loginEvent: EventEmitter<any> = new EventEmitter();

  constructor(private http: HttpClient) { }

  login(username: string, password: string): Observable<LoginResponseDTO> {
    return this.http.post<LoginResponseDTO>('/api/login', {
      username: username,
      password: password
    });
  }

  emitLoginEvent() {
    this.loginEvent.emit();
  }
}
