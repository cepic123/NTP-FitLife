import { Component, OnInit } from '@angular/core';
import { CreateExerciseDTO, Exercise } from '../models/exercise';
import { ExerciseService } from './services/exercise.service';

@Component({
  selector: 'app-exercise',
  templateUrl: './exercise.component.html',
  styleUrls: ['./exercise.component.css']
})
export class ExerciseComponent implements OnInit {

  exercise: CreateExerciseDTO = {
    name: "",
    description: "",
    img: ""
  };

  exercises: Exercise[] = []

  constructor(private exerciseService: ExerciseService) { }

  ngOnInit(): void {
    this.getAllExercises()
  }

  createExercise() {
    this.exerciseService.createExercise(this.exercise).subscribe((data) => {
      this.getAllExercises();
    })
  }

  getAllExercises() {
    this.exerciseService.getAllExercises().subscribe((data) => {
      this.exercises = data;
    });
  }

  validateExercise() {
    return !(this.exercise?.name !== "" && this.exercise?.description !== "")
  }

}
