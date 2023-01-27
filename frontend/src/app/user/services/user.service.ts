import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { RowToggler } from 'primeng/table';
import { Observable, throwError } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http: HttpClient) { }

  createUser(username: string, email: string, password: string, role: string): Observable<string> {
    return this.http.post<string>('/api/user', {
      username: username,
      password: password,
      email: email,
      role: role,
    });
  }
}
