import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Block } from '../models/block';
import { Complaint } from '../models/complaint';

@Injectable({
  providedIn: 'root'
})
export class ComplaintService {

  constructor(private http: HttpClient) { }

  createComplaint(complaint: Complaint): Observable<Complaint> {
    return this.http.post<Complaint>('/api/complaint', complaint);
  }

  getComplaints(subjectId: number): Observable<Complaint[]> {
    return this.http.get<Complaint[]>('/api/complaint/subject/' + subjectId);
  }

  blockUser(block: Block): Observable<Block[]> {
    return this.http.post<Block[]>('/api/block', block);
  }
}