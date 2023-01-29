import { Component, OnInit } from '@angular/core';
import { ExerciseService } from '../exercise/services/exercise.service';
import { CreateExerciseDTO, Exercise } from '../models/exercise';
import { Workout, Set, Rep } from '../models/workout';
import { UserWorkoutsService } from '../user-workouts/services/user-workouts.service';
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
  displayWorkout: boolean = false;

  workout: Workout = {
    name: "",
    description: "",
    sets: []
  }

  exerciseToAdd?: Exercise;

  exercises: Exercise[] = [];

  constructor(
    private workoutService: WorkoutService,
    private exerciseService: ExerciseService,
    private userWorkoutsService: UserWorkoutsService,
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

  validateWorkout() {
    if (this.workout.name && this.workout.name.length > 5 
      && this.workout.description && this.workout.description.length > 5) return false;
    return true;
  }

  creaeteWorkout() {
    this.workoutService.createWorkout(this.workout).subscribe((data) => {
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
