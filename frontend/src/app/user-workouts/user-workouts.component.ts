import { Component, OnInit } from '@angular/core';
import { Workout } from '../models/workout';
import { UserWorkoutsService } from './services/user-workouts.service';

@Component({
  selector: 'app-user-workouts',
  templateUrl: './user-workouts.component.html',
  styleUrls: ['./user-workouts.component.css']
})
export class UserWorkoutsComponent implements OnInit {

  comment: string = "";
  displayWorkout: boolean = false;
  displayCommentDialog: boolean = false;
  workouts: Workout[] = [];
  workout: Workout = {
    name: "",
    description: "",
    sets: []
  }

  constructor(private userWorkoutsService: UserWorkoutsService) { }

  ngOnInit(): void {
    this.getUserWorkouts();
  }

  createComment() {
    alert("Create Comment");
  }

  showDetailed(workoutId: number) {
    this.userWorkoutsService.getWorkout(workoutId).subscribe((data) => {
      this.workout = data;
      console.log(this.workout);
    });
    this.showWorkoutDialog();
  }

  getUserWorkouts() {
    var userId = localStorage.getItem("userId");
    var userIdNum;
    if (userId) {
      userIdNum = parseInt(userId);
    }
    var userWorkoutIds: number[] = [];
    this.userWorkoutsService.getUserWorkoutReferences(1).subscribe((data) => {
      for (var userWorkout of data) {
        userWorkoutIds.push(userWorkout.workoutReferenceID);
      }
      this.userWorkoutsService.getUserWorkouts(userWorkoutIds).subscribe((data) => {
        this.workouts = data;
      })
    })
  }

  showWorkoutDialog() {
    this.displayWorkout = true;
  }

  openCommentDialog() {
    this.displayCommentDialog = true;
  }
}
