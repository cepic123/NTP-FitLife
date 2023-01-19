import { Component, OnInit } from '@angular/core';
import { ExerciseService } from '../exercise/services/exercise.service';
import { CreateExerciseDTO, Exercise } from '../models/exercise';
import { Workout, Set } from '../models/workout';
import { WorkoutService } from './services/workout.service';

@Component({
  selector: 'app-workout',
  templateUrl: './workout.component.html',
  styleUrls: ['./workout.component.css']
})
export class WorkoutComponent implements OnInit {

  exercise: CreateExerciseDTO = {
    name: "",
    description: "",
    img: ""
  };

  workout: Workout = {
    name: "",
    description: "",
    sets: []
  }

  workouts: Workout[] = []

  exercises: Exercise[] = []

  constructor(
    private workoutService: WorkoutService,
    private exerciseService: ExerciseService,
    ) { }

  ngOnInit(): void {
  }

  getAllExercises() {
    this.exerciseService.getAllExercises().subscribe((data) => {
      this.exercises = data;
    });
  }

  createWorkoutJSON() {
    console.log(JSON.stringify(this.workout))
  }

  addSet() {
    let set: Set = {}
    this.workout.sets?.push(set)
  }
}
