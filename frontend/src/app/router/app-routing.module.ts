import {
  NgModule
} from '@angular/core';
import {
  RouterModule,
  Routes
} from '@angular/router';
import {
  UserComponent
} from '../user/user.component';
import {
  LoginComponent
} from '../login/login.component';
import {
  WorkoutComponent
} from '../workout/workout.component';
import {
  ExerciseComponent
} from '../exercise/exercise.component';
import {
  UserWorkoutsComponent
} from '../user-workouts/user-workouts.component';
import { AuthGuard } from './auth.guard';
import { LogoutComponent } from '../logout/logout.component';
import { AllWorkoutsComponent } from '../all-workouts/all-workouts.component';
import { UsersComponent } from '../users/users.component';
import { CoachRequestsComponent } from '../coach-requests/coach-requests.component';
import { CalendarComponent } from '../calendar/calendar.component';

const routes: Routes = [
  {
    path: 'user',
    component: UserComponent,
  },
  {
    path: 'login',
    component: LoginComponent
  },
  {
    path: 'logout',
    component: LogoutComponent,
    // canActivate: [AuthGuard],
    // data: {
    //   expectedRoles: ['user','coach','admin'],
    // },
  },
  {
    path: 'workout',
    component: WorkoutComponent,
    // canActivate: [AuthGuard],
    // data: {
    //   expectedRoles: ['coach'],
    // },
  },
  {
    path: 'exercise',
    component: ExerciseComponent,
    // canActivate: [AuthGuard],
    // data: {
    //   expectedRoles: ['coach'],
    // },
  },
  {
    path: 'user-workouts',
    component: UserWorkoutsComponent
  },
  {
    path: 'all-workouts',
    component: AllWorkoutsComponent
  },
  {
    path: 'all-users',
    component: UsersComponent
  },
  {
    path: 'coach-requests',
    component: CoachRequestsComponent
  },
  {
    path: 'calendar',
    component: CalendarComponent
  },
];
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}
