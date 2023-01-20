import { Component, OnInit } from '@angular/core';
import { ExerciseService } from '../exercise/services/exercise.service';
import { CreateExerciseDTO, Exercise } from '../models/exercise';
import { Workout, Set, Rep } from '../models/workout';
import { WorkoutService } from './services/workout.service';

@Component({
  selector: 'app-workout',
  templateUrl: './workout.component.html',
  styleUrls: ['./workout.component.css']
})
export class WorkoutComponent implements OnInit {

  setNum: number = 0;
  newRepNoReps: number = 0;

  display: boolean = false;

  workout: Workout = {
    name: "",
    description: "",
    sets: []
  }

  exerciseToAdd?: Exercise;

  workouts: Workout[] = [];

  exercises: Exercise[] = [];

  constructor(
    private workoutService: WorkoutService,
    private exerciseService: ExerciseService,
    ) { }

  ngOnInit(): void {
    this.getAllExercises()
  }

  getAllExercises() {
    this.exerciseService.getAllExercises().subscribe((data) => {
      this.exercises = data;
    });
  }

  createWorkoutJSON() {
    console.log(JSON.stringify(this.workout))
  }

  creaeteWorkout() {
    this.workoutService.createWorkout(this.workout).subscribe((data) => {
      alert("Here");
    })
  }
  validateRep() {
    if (this.newRepNoReps > 0 && this.exerciseToAdd) return false;
    return true;
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
    this.setNum = 0;
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
