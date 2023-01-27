import { Injectable, Output } from '@angular/core';
import { EventEmitter } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class LogoutService {

  @Output() logoutEvent: EventEmitter<any> = new EventEmitter();

  constructor() { }

  logoutUser() {
    localStorage.clear()
    this.logoutEvent.emit();
  }
}
