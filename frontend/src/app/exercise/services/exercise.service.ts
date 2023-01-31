import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { CreateExerciseDTO, Exercise } from '../../models/exercise';

@Injectable({
  providedIn: 'root'
})
export class ExerciseService {

  constructor(private http: HttpClient) { }

  createExercise(createExerciseDTO?: CreateExerciseDTO): Observable<Exercise>{
    return this.http.post<Exercise>('/api/exercise', createExerciseDTO);
  }

  getAllExercises(userId?: number): Observable<Exercise[]> {
    return this.http.get<Exercise[]>('/api/exercise/' + userId);
  }
}
