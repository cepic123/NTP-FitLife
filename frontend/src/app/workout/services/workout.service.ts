import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Workout } from '../../models/workout';

@Injectable({
  providedIn: 'root'
})
export class WorkoutService {

  constructor(private http: HttpClient) { }

  createWorkout(workout?: Workout): Observable<Workout>{
    return this.http.post<Workout>('/api/workout', workout);
  }
 
  getAllWorkouts(): Observable<Workout[]> {
    return this.http.post<Workout[]>('/api/userWorkouts', {
      workoutIds: null
      }
    );
  }
}
