import { TestBed } from '@angular/core/testing';

import { UserWorkoutsService } from './user-workouts.service';

describe('UserWorkoutsService', () => {
  let service: UserWorkoutsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(UserWorkoutsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
