import { Routes } from '@angular/router';
import { Home } from './pages/home/home';
import { AdicionarLivro } from './pages/adicionarlivro/adicionarlivro';
import { CriarEmprestimo } from './pages/criaremprestimo/criaremprestimo';

export const routes: Routes = [
  { path: '', component: Home },
  { path: 'adicionar-livro', component: AdicionarLivro },
  { path: 'criar-emprestimo', component: CriarEmprestimo }
];