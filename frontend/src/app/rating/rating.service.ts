import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Rating } from '../models/rating';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class RatingService {

  constructor(private http: HttpClient) { }

  createRating(rating: Rating): Observable<Rating> {
    return this.http.post<Rating>('/api/rating', rating);
  }

  getRating(subjectId: number, ratingType: string): Observable<Number> {
    return this.http.get<Number>('/api/rating/' + subjectId + '/' + ratingType);
  }

  updateRating(rating: Rating): Observable<Rating> {
    return this.http.put<Rating>('/api/rating', rating);
  }
  
  getRatingByUserAndSubject(userId: number, subjectId: number, ratingType: string): Observable<Rating> {
    return this.http.get<Rating>('/api/rating/' + userId + '/' + subjectId + '/' + ratingType);
  }
}
