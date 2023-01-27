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

const routes: Routes = [{
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
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}
