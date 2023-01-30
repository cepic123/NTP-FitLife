import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CoachRequestsComponent } from './coach-requests.component';

describe('CoachRequestsComponent', () => {
  let component: CoachRequestsComponent;
  let fixture: ComponentFixture<CoachRequestsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CoachRequestsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CoachRequestsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
