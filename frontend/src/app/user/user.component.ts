import { Component, OnInit } from '@angular/core';
import { Role } from '../models/user';
import { UserService } from './services/user.service';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {

  username: string = ""
  email: string = ""
  password1: string = ""
  password2: string = ""

  roles: Role[] = [
    {name: 'User', code: 'user'},
    {name: 'Coach', code: 'coach'},
  ];
  role: Role = {
    name: 'User', 
    code: 'user'
  };

  constructor(private userService: UserService) { }

  ngOnInit(): void {

  }

  createUser() {
    this.userService.createUser(this.username, this.email, this.password1, this.role.code).subscribe((data) => {
      alert(data);
    })
  }

  validate() {
    return this.validateEmail(this.email) === null || this.username.length < 5 ||
      this.password1.length < 5 || this.password2.length < 5 || this.password1 !== this.password2
  }

  validateEmail(email: string) {
    return String(email)
    .toLowerCase()
    .match(
      /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
    );
  }
}
