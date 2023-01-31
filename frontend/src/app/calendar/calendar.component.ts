import { Component, OnInit } from '@angular/core';
import { CalendarOptions } from '@fullcalendar/core';
import dayGridPlugin from '@fullcalendar/daygrid';
import interactionPlugin from '@fullcalendar/interaction';
import { CalendarEntry } from '../models/calendarEntry';
import { Workout } from '../models/workout';
import { UserWorkoutsService } from '../user-workouts/services/user-workouts.service';
import { CalendarService } from './services/calendar.service';

@Component({
  selector: 'app-calendar',
  templateUrl: './calendar.component.html',
  styleUrls: ['./calendar.component.css']
})
export class CalendarComponent implements OnInit {

  displayEntry: boolean = false;
  calendarEntries: CalendarEntry[] = [];
  workout: Workout = {sets: []};
  workouts: Workout[] = [];
  displayWorkout: boolean = false;
  displayWorkoutPicker: boolean = false;
  workoutDate: string = '';
  calendarOptions: CalendarOptions = {
    initialView: 'dayGridMonth',
    plugins: [dayGridPlugin, interactionPlugin],
    dateClick: this.handleDateClick.bind(this),
  };

  constructor(private userWorkoutsService: UserWorkoutsService, private calendarService: CalendarService,
     private userWorkoutService: UserWorkoutsService) { }

  ngOnInit(): void {
    this.getUserWorkouts();
    this.getCalendarEntries();
  }

  removeEntry() {
    console.log(this.calendarEntries)
    for (let calendarEntry of this.calendarEntries) {
      if (calendarEntry.date === this.workoutDate) {
        this.calendarService.deleteCalendarEntry(calendarEntry.ID).subscribe((data) => {
          this.getCalendarEntries();
          this.displayEntry = false;
        })
      }
    }
  }

  handleDateClick(arg: any) {
    this.workoutDate = arg.dateStr;
    for (let calendarEntry of this.calendarEntries) {
      if (calendarEntry.date === arg.dateStr) {
        this.userWorkoutService.getWorkout(calendarEntry.workoutId).subscribe((data) => {
          this.workout = data;
          this.displayEntryDialog()
        })
        return
      }
    }
    this.displayWorkoutPickerDialog()
    console.log(arg)
  }

  getCalendarEntries() {
    var userId = localStorage.getItem("userId");
    if (userId) {
      this.calendarService.getCalendarEntries(parseInt(userId)).subscribe((data) => {
        this.calendarEntries = data;
        var events = []
        for (let calendarEntry of this.calendarEntries) {
          events.push({title: calendarEntry.workoutName, date: calendarEntry.date})
        }
        this.calendarOptions.events = events;
      })
    }
  }

  showDetailed(workoutId: number) {
    this.userWorkoutsService.getWorkout(workoutId).subscribe((data) => {
      this.workout = data;
      console.log(this.workout);
    });
    this.displayWorkoutDialog()
  }

  addToDate(workoutId: number, workoutName: string) {
    var userId = localStorage.getItem("userId");
    if (userId) {
      this.calendarService.createCalendarEntry(parseInt(userId), workoutId, this.workoutDate, workoutName).subscribe((data) => {
        this.getCalendarEntries();
        this.displayWorkoutPicker = false;
      })
    }
  }

  displayWorkoutPickerDialog() {
    this.displayWorkoutPicker = true;
  }

  displayWorkoutDialog() {
    this.displayWorkout = true;
  }

  displayEntryDialog() {
    this.displayEntry = true;
  }

  getUserWorkouts() {
    var userId = localStorage.getItem("userId");
    var userIdNum;
    if (userId) {
      userIdNum = parseInt(userId);
    }
    var userWorkoutIds: number[] = [];
    this.userWorkoutsService.getUserWorkoutReferences(userIdNum).subscribe((data) => {
      for (var userWorkout of data) {
        userWorkoutIds.push(userWorkout.workoutReferenceID);
      }
      if (userWorkoutIds.length > 0) {
        this.userWorkoutsService.getUserWorkouts(userWorkoutIds).subscribe((data) => {
          this.workouts = data;
        })
      }
    })
  }
}
