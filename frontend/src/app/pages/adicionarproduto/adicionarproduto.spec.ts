import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Adicionarproduto } from './adicionarproduto';

describe('Adicionarproduto', () => {
  let component: Adicionarproduto;
  let fixture: ComponentFixture<Adicionarproduto>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Adicionarproduto]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Adicionarproduto);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
