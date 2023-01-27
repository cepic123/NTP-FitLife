import { Component, OnInit } from '@angular/core';
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
      label: 'User',
      icon: 'pi pi-fw pi-sign-in',
      routerLink: '/user',
    },
    {
      label: 'Workout',
      icon: 'pi pi-fw pi-sign-in',
      routerLink: '/workout',
    },
    {
      label: 'My Workouts',
      icon: 'pi pi-fw pi-sign-in',
      routerLink: '/user-workouts',
    },
    {
      label: 'Exercise',
      icon: 'pi pi-fw pi-sign-in',
      routerLink: '/exercise',
    },
    {
      label: 'Logout',
      icon: 'pi pi-fw pi-sign-in',
      routerLink: '/logout',
    }
  ];

  // items: MenuItem[] = [
  //   {
  //     label: 'Login',
  //     icon: 'pi pi-fw pi-sign-in',
  //     routerLink: '/login',
  //   },
  //   {
  //     label: 'Register',
  //     icon: 'pi pi-fw pi-sign-in',
  //     routerLink: '/user',
  //   }
  // ];

  constructor(
    private loginService: LoginService,
    private logoutService: LogoutService
  ) {
    // this.loginService.loginEvent.subscribe(() => {
    //   this.setNavbarItems();
    // });
    // this.logoutService.logoutEvent.subscribe(() => {
    //   this.setLogoutItems();
    // });
  }

  setNavbarItems = () => {
    var role = localStorage.getItem('role');
    if (role === null || role === '') {
      this.setLogoutItems();
    } else {
      this.items = [
        {
          label: 'Workout',
          icon: 'pi pi-fw pi-sign-in',
          routerLink: '/workout',
          visible: role === 'coach',
        },
        {
          label: 'My Workouts',
          icon: 'pi pi-fw pi-sign-in',
          routerLink: '/user-workouts',
          visible: role === 'user',
        },
        {
          label: 'Exercise',
          icon: 'pi pi-fw pi-sign-in',
          routerLink: '/exercise',
          visible: role === 'coach',
        },
        {
          label: 'Logout',
          icon: 'pi pi-fw pi-sign-in',
          routerLink: '/logout',
          visible: role === 'coach' || role === 'user',
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
    // this.setNavbarItems();
  }

}
