import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Comment } from '../models/comment';

@Injectable({
  providedIn: 'root'
})
export class CommentService {

  constructor(private http: HttpClient) { }

  createComment(comment: Comment): Observable<Comment> {
    return this.http.post<Comment>('/api/comment', comment);
  }

  updateComment(comment: Comment): Observable<Comment> {
    return this.http.put<Comment>('/api/comment', comment);
  }
  
  getCommentByUserAndSubject(userId: number, workoutId: number, commentType: string): Observable<Comment> {
    return this.http.get<Comment>('/api/comment/' + userId + '/' + workoutId + '/' + commentType);
  }
}
