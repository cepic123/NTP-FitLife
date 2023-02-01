import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MenuItem } from 'primeng/api';
import { LoginService } from '../login/services/login.service';
import { LogoutService } from '../logout/services/logout.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {
  
  items: MenuItem[] = [
    {
      label: 'Login',
      icon: 'pi pi-fw pi-sign-in',
      routerLink: '/login',
    },
    {
      label: 'Register',
      icon: 'pi pi-fw pi-sign-in',
      routerLink: '/user',
    }
  ];

  constructor(
    private loginService: LoginService,
    private logoutService: LogoutService,
    private router: Router
  ) {
    this.loginService.loginEvent.subscribe(() => {
      this.setNavbarItems();
    });
    this.logoutService.logoutEvent.subscribe(() => {
      this.setLogoutItems();
    });
  }
  
  setNavbarItems = () => {
    var role = localStorage.getItem('role');
    if (role === 'admin') {
      this.router.navigate(['coach-requests']);
    } else {
      this.router.navigate(['calendar']);
    }
    if (role === null || role === '') {
      this.setLogoutItems();
    } else {
      this.items = [
        {
          label: 'Workout',
          icon: 'pi pi-fw pi-sign-in',
          routerLink: '/workout',
          visible: role === 'coach'
        },
        {
          label: 'My Workouts',
          icon: 'pi pi-fw pi-sign-in',
          routerLink: '/user-workouts',
        },
        {
          label: 'All Workouts',
          icon: 'pi pi-fw pi-sign-in',
          routerLink: '/all-workouts',
          visible: role !== 'coach'
        },
        {
          label: 'Exercise',
          icon: 'pi pi-fw pi-sign-in',
          routerLink: '/exercise',
          visible: role === 'coach'
        },
        {
          label: 'View Users',
          icon: 'pi pi-fw pi-sign-in',
          routerLink: '/all-users',
        },
        {
          label: 'Coach Requests',
          icon: 'pi pi-fw pi-sign-in',
          routerLink: '/coach-requests',
          visible: role === 'admin'
        },
        {
          label: 'Calendar',
          icon: 'pi pi-fw pi-sign-in',
          routerLink: '/calendar',
        },
        {
          label: 'Logout',
          icon: 'pi pi-fw pi-sign-in',
          routerLink: '/logout',
        }
      ]
    }
  }

  setLogoutItems = () => {
    this.items = [
      {
        label: 'Login',
        icon: 'pi pi-fw pi-sign-in',
        routerLink: '/login',
      },
      {
        label: 'Register',
        icon: 'pi pi-fw pi-sign-in',
        routerLink: '/user',
      }
    ]
  }
  
  ngOnInit(): void {
    this.setNavbarItems();
  }

}
