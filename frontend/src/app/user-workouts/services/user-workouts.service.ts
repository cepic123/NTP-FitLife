import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { UserWorkout } from '../../models/userWorkout';
import { Workout } from '../../models/workout';

@Injectable({
  providedIn: 'root'
})
export class UserWorkoutsService {

  constructor(private http: HttpClient) { }
  
  getUserWorkoutReferences(userId?: number): Observable<UserWorkout[]> {
    return this.http.get<UserWorkout[]>('/api/userWorkoutRefs/' + userId);
  }

  getUserWorkouts(workoutIds?: number[]): Observable<Workout[]> {
    return this.http.post<Workout[]>('/api/userWorkouts', {
      workoutIds: workoutIds
      }
    );
  }

  getWorkout(workoutId?: number): Observable<Workout> {
    return this.http.get<Workout>('/api/workout/' + workoutId);
  }
  
  deleteWorkout(workoutId?: number): Observable<Workout> {
    return this.http.delete<Workout>('/api/workout/' + workoutId);
  }

  removeFromUser(userId: number, workoutId: number): Observable<string> {
    return this.http.delete<string>('/api/user/' + userId + '/' + workoutId);
  }
}
