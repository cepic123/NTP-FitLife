import { TestBed } from '@angular/core/testing';

import { AllWorkoutsService } from './all-workouts.service';

describe('AllWorkoutsService', () => {
  let service: AllWorkoutsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AllWorkoutsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
