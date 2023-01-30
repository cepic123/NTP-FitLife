import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from 'src/app/models/user';

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

  getAllUsers(): Observable<User[]> {
    return this.http.get<User[]>('/api/user');
  }

  deleteUser(userId: number): Observable<Object> {
    return this.http.delete<User[]>('/api/user/' + userId);
  }

  permanentlyDeleteUser(userId: number): Observable<Object> {
    return this.http.delete<Object>('/api/user/delete/' + userId);
  }

  restoreUser(userId: number): Observable<Object> {
    return this.http.put<Object>('/api/user/restore/' + userId, {});
  }
}
