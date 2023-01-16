import { Component, OnInit } from '@angular/core';
import { CreateUserService } from './services/create-user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  username: string = ''
  password: string = ''

  constructor(
    private createUserService: CreateUserService,
  ) { }

  ngOnInit(): void {
  }
  
  login() {
    this.createUserService.login(this.username, this.password).subscribe((data) => {
      alert(data);
    })
  }
}

