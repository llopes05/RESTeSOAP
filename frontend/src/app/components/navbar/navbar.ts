import { Component } from '@angular/core';
import { RouterModule } from '@angular/router';
import { MenubarModule } from 'primeng/menubar';
import { ButtonModule } from 'primeng/button';
import { MenuItem } from 'primeng/api';

@Component({
  selector: 'app-navbar',
  imports: [MenubarModule, ButtonModule, RouterModule],
  templateUrl: './navbar.html',
  styleUrl: './navbar.css'
})
export class Navbar {
  items: MenuItem[] = [
    {
      label: 'Início',
      icon: 'pi pi-home',
      routerLink: ['/']
    },
    {
      label: 'Criar pedido',
      icon: 'pi pi-truck',
      routerLink: ['/criar-pedido']
    },
    {
      label: 'Adicionar produto',
      icon: 'pi pi-plus',
      routerLink: ['/adicionar-produto']
    }
    
  ];
}
