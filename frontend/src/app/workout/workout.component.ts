import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AllWorkoutsService } from '../all-workouts/services/all-workouts.service';
import { ExerciseService } from '../exercise/services/exercise.service';
import { Exercise } from '../models/exercise';
import { Workout, Set, Rep } from '../models/workout';
import { WorkoutService } from './services/workout.service';

@Component({
  selector: 'app-workout',
  templateUrl: './workout.component.html',
  styleUrls: ['./workout.component.css']
})
export class WorkoutComponent implements OnInit {

  userId?: number;
  setNum: number = 0;
  newRepNoReps: number = 0;
  breakLngth: number = 0;
  noSets: number = 0;
  display: boolean = false;
  displayWorkout: boolean = false;

  workout: Workout = {
    name: "",
    description: "",
    sets: []
  }
  comments: Comment[] = [];
  exerciseToAdd?: Exercise;

  exercises: Exercise[] = [];

  constructor(
    private workoutService: WorkoutService,
    private exerciseService: ExerciseService,
    private allWorkoutService: AllWorkoutsService,
    private router: Router
    ) { }

  ngOnInit(): void {
    var userId = localStorage.getItem('userId')
    this.userId = userId ? parseInt(userId) : undefined
    this.getAllExercises()
  }

  getAllExercises() {
    this.exerciseService.getAllExercises(this.userId).subscribe((data) => {
      this.exercises = data;
    });
  }

  validateWorkout() {
    if (this.workout.name && this.workout.name.length > 5 
      && this.workout.description && this.workout.description.length > 5) return false;
    return true;
  }

  creaeteWorkout() {
    this.workout.coachId = this.userId;
    this.workoutService.createWorkout(this.workout).subscribe((data) => {
      var userId = localStorage.getItem("userId");
      if (userId && data.ID) {
        this.allWorkoutService.addWorkoutToUser(parseInt(userId), data.ID).subscribe((data) => {
          this.router.navigate(['user-workouts']);
        })
      }
    })
  }

  validateRep() {
    if (this.newRepNoReps > 0 && this.exerciseToAdd) return false;
    return true;
  }
  
  showWorkoutDialog() {
    this.displayWorkout = true;
  }

  addRep() {
    let rep: Rep = {
      exercise: this.exerciseToAdd,
      noReps: this.newRepNoReps,
      orderNum: this.workout.sets[this.setNum].reps?.length
    }
    this.workout.sets[this.setNum].reps?.push(rep);
    this.exerciseToAdd = undefined;
    this.newRepNoReps = 0;
  }

  addSet() {
    let set: Set = {
      orderNum: this.workout.sets.length,
      reps: []
    }
    this.workout.sets?.push(set)
  }

  showDialog(setNum: number) {
    this.setNum = setNum;
    this.display = true;
  }
}
