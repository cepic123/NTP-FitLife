import { Component, OnInit } from '@angular/core';
import { Workout } from '../models/workout';
import { UserWorkoutsService } from '../user-workouts/services/user-workouts.service';
import { WorkoutService } from '../workout/services/workout.service';
import { AllWorkoutsService } from './services/all-workouts.service';

@Component({
  selector: 'app-all-workouts',
  templateUrl: './all-workouts.component.html',
  styleUrls: ['./all-workouts.component.css']
})
export class AllWorkoutsComponent implements OnInit {

  workouts: Workout[] = [];
  displayWorkout: boolean = false;

  workout: Workout = {
    name: "",
    description: "",
    sets: []
  }

  constructor(
    private workoutService: WorkoutService,
    private userWorkoutsService: UserWorkoutsService,
    private allWorkoutsService: AllWorkoutsService
  ) { }

  addWorkoutToUser(workoutId: number) {
    var userId = localStorage.getItem("userId");
    if (userId) {
      this.allWorkoutsService.AddWorkoutToUser(parseInt(userId), workoutId).subscribe((data) => {
        alert(data);
      })
    }
  }

  ngOnInit(): void {
    this.getAllWorkouts()
  }

  getAllWorkouts() {
    this.workoutService.getAllWorkouts().subscribe((data) => {
      this.workouts = data;
    })
  }

  showDetailed(workoutId: number) {
    this.userWorkoutsService.getWorkout(workoutId).subscribe((data) => {
      this.workout = data;
      console.log(this.workout);
    });
    this.showWorkoutDialog();
  }

  showWorkoutDialog() {
    this.displayWorkout = true;
  }
}
