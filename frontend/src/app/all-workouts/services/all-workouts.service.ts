import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Exercise } from 'src/app/models/exercise';

@Injectable({
  providedIn: 'root'
})
export class AllWorkoutsService {

  constructor(private http: HttpClient) { }

  addWorkoutToUser(userId: number, workoutId: number): Observable<string>{
    return this.http.post<string>('/api/user/' + userId + '/' + workoutId, {});
  }
}
