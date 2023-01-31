import { Component, OnInit } from '@angular/core';
import { CommentService } from '../comment/comment.service';
import { Comment } from '../models/comment';
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
  displayComments: boolean = false;
  workout: Workout = {
    name: "",
    description: "",
    sets: []
  }
  comments: Comment[] = [];

  constructor(
    private workoutService: WorkoutService,
    private userWorkoutsService: UserWorkoutsService,
    private allWorkoutsService: AllWorkoutsService,
    private commentSerivce: CommentService
  ) { }

  showComments(workoutId: number) {
    this.commentSerivce.getSubjectComments(workoutId, "WORKOUT").subscribe((data) => {
      this.comments = data;
      this.displayComments = true;
    })
  }

  addWorkoutToUser(workoutId: number) {
    var userId = localStorage.getItem("userId");
    if (userId) {
      this.allWorkoutsService.addWorkoutToUser(parseInt(userId), workoutId).subscribe((data) => {
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
