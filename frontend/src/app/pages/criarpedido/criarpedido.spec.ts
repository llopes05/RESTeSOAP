import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Criarpedido } from './criarpedido';

describe('Criarpedido', () => {
  let component: Criarpedido;
  let fixture: ComponentFixture<Criarpedido>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Criarpedido]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Criarpedido);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
