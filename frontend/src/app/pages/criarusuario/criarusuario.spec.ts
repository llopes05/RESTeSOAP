import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Criarusuario } from './criarusuario';

describe('Criarusuario', () => {
  let component: Criarusuario;
  let fixture: ComponentFixture<Criarusuario>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Criarusuario]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Criarusuario);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
