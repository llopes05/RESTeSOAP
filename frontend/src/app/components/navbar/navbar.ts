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
      label: 'Adicionar Livro',
      icon: 'pi pi-book',
      routerLink: ['/adicionar-livro']
    },
    {
      label: 'Criar Empréstimo',
      icon: 'pi pi-calendar',
      routerLink: ['/criar-emprestimo']
    },
    {
      label: 'Criar Usuário',
      icon: 'pi pi-user-plus',
      routerLink: ['/criarusuario']
    }
  ];
}
