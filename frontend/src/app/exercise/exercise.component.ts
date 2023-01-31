import { Component, OnInit } from '@angular/core';
import { CreateExerciseDTO, Exercise } from '../models/exercise';
import { ExerciseService } from './services/exercise.service';

@Component({
  selector: 'app-exercise',
  templateUrl: './exercise.component.html',
  styleUrls: ['./exercise.component.css']
})
export class ExerciseComponent implements OnInit {

  userId?: number;
  exercise: CreateExerciseDTO = {
    name: "",
    description: "",
    img: ""
  };

  exercises: Exercise[] = []

  constructor(private exerciseService: ExerciseService) { }

  ngOnInit(): void {
    var userId = localStorage.getItem('userId')
    this.userId = userId ? parseInt(userId) : undefined
    this.getAllExercises()
  }

  createExercise() {
    this.exercise.coachId = this.userId;
    this.exerciseService.createExercise(this.exercise).subscribe((data) => {
      this.getAllExercises();
    })
  }

  getAllExercises() {
    this.exerciseService.getAllExercises(this.userId).subscribe((data) => {
      this.exercises = data;
    });
  }

  validateExercise() {
    return !(this.exercise?.name !== "" && this.exercise?.description !== "")
  }

}
