import { Component, OnInit } from '@angular/core';
import { Workout } from '../models/workout';
import { Comment } from '../models/comment';
import { UserWorkoutsService } from './services/user-workouts.service';
import { CommentService } from '../comment/comment.service';

@Component({
  selector: 'app-user-workouts',
  templateUrl: './user-workouts.component.html',
  styleUrls: ['./user-workouts.component.css']
})
export class UserWorkoutsComponent implements OnInit {

  comment: Comment = {
    comment: ""
  };
  selectedWorkout: number = 0;
  displayWorkout: boolean = false;
  displayCommentDialog: boolean = false;
  workouts: Workout[] = [];
  workout: Workout = {
    name: "",
    description: "",
    sets: []
  }

  constructor(private userWorkoutsService: UserWorkoutsService,
    private commentService: CommentService) { }

  ngOnInit(): void {
    this.getUserWorkouts();
  }

  createComment(commentType: string) {
    var userId = localStorage.getItem("userId");
    this.comment.userID = userId ?  parseInt(userId) : 0; 
    var username = localStorage.getItem("username");
    this.comment.username = username ? username : undefined; 
    this.comment.subjectID = this.selectedWorkout;
    this.comment.commentType = commentType;
    this.commentService.createComment(this.comment).subscribe((data) => {
      console.log(data);
    })
  }

  validateComment() {
    return this.comment.comment.length < 10
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

  openCommentDialog(workoutId: number) {
    this.selectedWorkout = workoutId;
    this.commentService.getCommentByUserAndSubject(1, workoutId, "WORKOUT").subscribe((data) => {
      this.comment = data;
    })
    this.displayCommentDialog = true;
  }
}
