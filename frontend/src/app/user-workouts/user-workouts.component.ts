import { Component, OnInit } from '@angular/core';
import { Workout } from '../models/workout';
import { Comment } from '../models/comment';
import { Rating } from '../models/rating';
import { UserWorkoutsService } from './services/user-workouts.service';
import { CommentService } from '../comment/comment.service';
import { RatingService } from '../rating/rating.service';
import { getSafePropertyAccessString } from '@angular/compiler';

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
  commentExists: boolean = false;
  rating: Rating = {
    rating: 0 
  }
  ratingExists: boolean = false;

  constructor(private userWorkoutsService: UserWorkoutsService,
    private commentService: CommentService,
    private ratingService: RatingService) { }

  ngOnInit(): void {
    this.getUserWorkouts();
  }

  removeFromUser(workoutId: number) {
    var userId = localStorage.getItem("userId");
    if (userId) {
      this.userWorkoutsService.removeFromUser(parseInt(userId), workoutId).subscribe((data) => {
        alert(data);
        this.getUserWorkouts();
      })
    }
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
    this.getComment(this.comment.userID, this.comment.subjectID);
  }

  createRating(ratingType: string) {
    var userId = localStorage.getItem("userId");
    this.rating.userID = userId ?  parseInt(userId) : 0; 
    var username = localStorage.getItem("username");
    this.rating.username = username ? username : undefined; 
    this.rating.subjectID = this.selectedWorkout;
    this.rating.ratingType = ratingType;
    this.ratingService.createRating(this.rating).subscribe((data) => {
      console.log(data);
    })
    this.getRating(this.rating.userID, this.rating.subjectID);
  }

  updateComment() {
    this.commentService.updateComment(this.comment).subscribe((data) => {
      console.log(data);
    })
  }

  updateRating() {
    this.ratingService.updateRating(this.rating).subscribe((data) => {
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
    this.userWorkoutsService.getUserWorkoutReferences(userIdNum).subscribe((data) => {
      for (var userWorkout of data) {
        userWorkoutIds.push(userWorkout.workoutReferenceID);
      }
      if (userWorkoutIds.length > 0) {
        this.userWorkoutsService.getUserWorkouts(userWorkoutIds).subscribe((data) => {
          this.workouts = data;
        })
      }
    })
  }

  showWorkoutDialog() {
    this.displayWorkout = true;
  }

  openCommentDialog(workoutId: number) {
    this.selectedWorkout = workoutId;
    var userId = localStorage.getItem("userId");
    if (userId) {
      this.getRating(parseInt(userId), workoutId);
      this.getComment(parseInt(userId), workoutId);
      this.displayCommentDialog = true;
    }
  }


  getRating(userId: number, workoutId: number) {
    this.ratingService.getRatingByUserAndSubject(userId, workoutId, "WORKOUT").subscribe((data) => {
      this.ratingExists = data.username !== "";
      this.rating = data ? data : {
        rating: 0
      };
      console.log(this.rating)
    });
  }

  getComment(userId: number, workoutId: number) {
    this.commentService.getCommentByUserAndSubject(userId, workoutId, "WORKOUT").subscribe((data) => {
      this.commentExists = data.comment !== "";
      this.comment = data ? data : {
        comment: ""
      };
    });
  }
}
