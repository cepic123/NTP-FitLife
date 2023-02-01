import { Component, OnInit } from '@angular/core';
import { CommentService } from '../comment/comment.service';
import { ComplaintService } from '../complaint/complaint.service';
import { Block } from '../models/block';
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

  userId?: number;
  workouts: Workout[] = [];
  displayWorkout: boolean = false;
  displayComments: boolean = false;
  workout: Workout = {
    name: "",
    description: "",
    sets: []
  }
  comments: Comment[] = [];
  blocks: Block[]  = [];
  blockIds: any[] = [];

  constructor(
    private workoutService: WorkoutService,
    private userWorkoutsService: UserWorkoutsService,
    private allWorkoutsService: AllWorkoutsService,
    private commentSerivce: CommentService,
    private complaintService: ComplaintService
  ) { }

  ngOnInit(): void {
    var userId = localStorage.getItem('userId')
    this.userId = userId ? parseInt(userId) : undefined
    this.getBlockedUsers();
  }

  getBlockedUsers() {
    this.complaintService.getBlockedUsers(this.userId).subscribe((data) => {
      this.blocks = data;
      this.blockIds = this.blocks.map((block) => block.block_subject_id)
      this.getAllWorkouts();
    })
  }

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
      })
    }
  }

  getAllWorkouts() {
    this.workoutService.getAllWorkouts().subscribe((data) => {
      this.workouts = data;
      for (let block of this.blockIds) {
        this.workouts = this.workouts.filter((el) => {return el.coachId != block})
      }
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
