import { Component, OnInit } from '@angular/core';
import { User } from '../models/user';
import { UserService } from '../user/services/user.service';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {

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
}
