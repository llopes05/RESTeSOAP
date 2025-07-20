import { Routes } from '@angular/router';
import { Home } from './pages/home/home';
import { Adicionarproduto } from './pages/adicionarproduto/adicionarproduto';
import { Criarpedido } from './pages/criarpedido/criarpedido';

export const routes: Routes = [
  { path: '', component: Home },
  { path: 'adicionar-produto', component: Adicionarproduto },
  { path: 'criar-pedido', component: Criarpedido }
];