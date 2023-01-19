import { Component, OnInit } from '@angular/core';
import { MenuItem } from 'primeng/api';

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
      label: 'Exercise',
      icon: 'pi pi-fw pi-sign-in',
      routerLink: '/exercise',
    },
  ];

  constructor() { }

  setNavbarItems = () => {
    
  }

  ngOnInit(): void {
    this.setNavbarItems();
  }

}
