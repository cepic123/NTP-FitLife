import { Component, OnInit } from '@angular/core';
import { ComplaintService } from '../complaint/complaint.service';
import { Block } from '../models/block';
import { Complaint } from '../models/complaint';
import { User } from '../models/user';
import { UserService } from '../user/services/user.service';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {

  username: string = ""
  userId?: number = 0
  role?: string = "";
  users: User[] = [];
  complaint: Complaint = {
    complaint_text: "",
    subject_name: "",
    user_name: ""
  };
  complaints: Complaint[] = [];
  displayComplaints: boolean = false;

  displayComplaintInput: boolean = false;

  constructor(private userService: UserService,
    private complaintService: ComplaintService) { }

  validateComplaint() {
    return !(this.complaint.complaint_text.length > 5);
  }

  saveComplaint() {
    this.complaintService.createComplaint(this.complaint).subscribe((data) => {
      this.displayComplaintInput = false;
    })
  }
  
  blockUser(userId: number, username: string) {
    var block: Block = {
      user_id: this.userId,
      block_subject_id: userId,
      user_name: this.username,
      subject_name: username,
    }
    this.complaintService.blockUser(block).subscribe((data) => {
      alert(data);
    })
  }
  
  getComplaints(userId: number) {
    this.complaintService.getComplaints(userId).subscribe((data) => {
      this.complaints = data;
      this.openCommentDialog();
    })
  }

  ngOnInit(): void {
    var role = localStorage.getItem("role");
    this.role = role ? role : undefined;

    var userId = localStorage.getItem("userId");
    this.userId = userId ? parseInt(userId) : undefined;
    
    var username = localStorage.getItem("username");
    this.username = username ? username : "";
    this.getAllUsers()
  }

  getAllUsers() {
    this.userService.getAllUsers().subscribe((data) => {
      this.users = data;
    })
    console.log(this.users)
  }

  deleteUser(userId: number) {
    this.userService.deleteUser(userId).subscribe((data) => {
      this.getAllUsers();
    })
  }

  restoreUser(userId: number) {
    this.userService.restoreUser(userId).subscribe((data) => {
      this.getAllUsers();
    })
  }

  openCommentDialog() {
    this.displayComplaints = true;
  }

  openCommentInputDialog(userId: number, username: string) {
    this.complaint.complaint_subject_id = userId;
    this.complaint.user_name = this.username;
    this.complaint.user_id = userId;
    this.displayComplaintInput = true;
  }
}
