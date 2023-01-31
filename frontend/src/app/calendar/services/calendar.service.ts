import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { CalendarEntry } from 'src/app/models/calendarEntry';

@Injectable({
  providedIn: 'root'
})
export class CalendarService {

  constructor(private http: HttpClient) { }

  createCalendarEntry(userId: number, workoutId: number, date: string, workoutName: string): Observable<CalendarEntry> {
    return this.http.post<CalendarEntry>('/api/workout/calendar/' + userId + '/' + workoutId + '/' + date + '/' + workoutName, {});
  }

  getCalendarEntries(userId: number): Observable<CalendarEntry[]> {
    return this.http.get<CalendarEntry[]>('/api/workout/calendar/' + userId);
  }

  deleteCalendarEntry(entryId: number): Observable<CalendarEntry> {
    return this.http.delete<CalendarEntry>('/api/workout/calendar/' + entryId);
  }
}