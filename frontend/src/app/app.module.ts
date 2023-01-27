import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { UserComponent } from './user/user.component';
import { AppRoutingModule } from './router/app-routing.module';
import { MenubarModule } from 'primeng/menubar';
import { NavbarComponent } from './navbar/navbar.component';
import { LoginComponent } from './login/login.component';
import { FormsModule } from '@angular/forms';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { InterceptorService } from './interceptor/interceptor.service';
import { InputTextModule}  from 'primeng/inputtext';
import { ButtonModule } from 'primeng/button';
import { WorkoutComponent } from './workout/workout.component';
import { TableModule } from 'primeng/table';
import { ExerciseComponent } from './exercise/exercise.component';
import { DialogModule } from 'primeng/dialog';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { UserWorkoutsComponent } from './user-workouts/user-workouts.component';
import { InputTextareaModule } from 'primeng/inputtextarea';
import { LogoutComponent } from './logout/logout.component';
import { DropdownModule } from 'primeng/dropdown';

@NgModule({
  declarations: [
    AppComponent,
    UserComponent,
    NavbarComponent,
    LoginComponent,
    WorkoutComponent,
    ExerciseComponent,
    UserWorkoutsComponent,
    LogoutComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    MenubarModule,
    FormsModule,
    HttpClientModule,
    InputTextModule,
    ButtonModule,
    TableModule,
    DialogModule,
    BrowserAnimationsModule,
    InputTextareaModule,
    DropdownModule
  ],
  providers: [{
    provide: HTTP_INTERCEPTORS,
    useClass: InterceptorService,
    multi: true
   },],
  bootstrap: [AppComponent]
})
export class AppModule { }
