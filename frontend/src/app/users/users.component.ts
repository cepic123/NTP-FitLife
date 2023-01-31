import { Component, OnInit } from '@angular/core';
import { ComplaintService } from '../complaint/complaint.service';
import { Complaint } from '../models/complaint';
import { User } from '../models/user';
import { UserService } from '../user/services/user.service';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {

  role?: string = "";
  users: User[] = [];
  complaint: Complaint = {
    complaint_text: "",
    subject_name: "",
    user_name: ""
  };
  complaints: Complaint[] = [];
  displayComplaints: boolean = false;

  constructor(private userService: UserService,
    private complaintService: ComplaintService) { }

  getComplaints(userId: number) {
    this.complaintService.getComplaints(userId).subscribe((data) => {
      this.complaints = data;
      this.openCommentDialog();
    })
  }

  ngOnInit(): void {
    var role = localStorage.getItem("role");
    this.role = role ? role : undefined;
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
}
