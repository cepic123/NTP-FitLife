import { Component, OnInit } from '@angular/core';
import { User } from '../models/user';
import { UserService } from '../user/services/user.service';

@Component({
  selector: 'app-coach-requests',
  templateUrl: './coach-requests.component.html',
  styleUrls: ['./coach-requests.component.css']
})
export class CoachRequestsComponent implements OnInit {

  users: User[] = [];
  constructor(private userService: UserService) { }

  ngOnInit(): void {
    this.getAllUsers()
  }

  getAllUsers() {
    this.userService.getAllUsers().subscribe((data) => {
      this.users = data;
    })
    console.log(this.users)
  }

  permanentlyDeleteUser(userId: number) {
    this.userService.permanentlyDeleteUser(userId).subscribe((data) => {
      this.getAllUsers();
    })
  }

  restoreUser(userId: number) {
    this.userService.restoreUser(userId).subscribe((data) => {
      this.getAllUsers();
    })
  }

}
