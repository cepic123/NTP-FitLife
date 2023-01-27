import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LogoutService } from './services/logout.service';

@Component({
  selector: 'app-logout',
  templateUrl: './logout.component.html',
  styleUrls: ['./logout.component.css']
})
export class LogoutComponent implements OnInit {

  constructor(private logoutService: LogoutService, private router: Router) { }

  ngOnInit(): void {
    this.logoutService.logoutUser();
    this.router.navigate(['login']);
  }

}
