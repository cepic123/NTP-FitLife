import { Component, OnInit } from '@angular/core';
import { LoginService } from './services/login.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  username: string = ''
  password: string = ''

  constructor(
    private LoginService: LoginService,
  ) { }

  ngOnInit(): void {
  }
  
  login() {
    this.LoginService.login(this.username, this.password).subscribe((data) => {
      localStorage.setItem('x-jwt-token', data)
    })
  }
}

