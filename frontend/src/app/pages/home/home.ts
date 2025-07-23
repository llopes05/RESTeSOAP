import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { TableModule } from 'primeng/table';
import { SelectButtonModule } from 'primeng/selectbutton';
import { InputTextModule } from 'primeng/inputtext';
import { IconFieldModule } from 'primeng/iconfield';
import { InputIconModule } from 'primeng/inputicon';
import { ButtonModule } from 'primeng/button';
import { TagModule } from 'primeng/tag';
import { Navbar } from '../../components/navbar/navbar';

@Component({
  selector: 'app-home',
  imports: [
    CommonModule, 
    TableModule, 
    SelectButtonModule, 
    FormsModule, 
    InputTextModule, 
    IconFieldModule, 
    InputIconModule, 
    ButtonModule,
    TagModule,
    HttpClientModule,
    Navbar
  ],
  templateUrl: './home.html',
  styleUrl: './home.css'
})
export class Home {
  // URL do gateway Django (será criado)
  private gatewayUrl = 'http://localhost:8000/api';
  
  // Opções de visualização
  viewOptions = [
    { label: 'Livros', value: 'livros', icon: 'pi pi-book' },
    { label: 'Usuários', value: 'usuarios', icon: 'pi pi-users' },
    { label: 'Empréstimos', value: 'emprestimos', icon: 'pi pi-calendar' }
  ];
  
  selectedView = 'livros';
  loading = false;

  // Dados para exibição
  livros: any[] = [];
  usuarios: any[] = [];
  emprestimos: any[] = [];

  constructor(private http: HttpClient) {}

  get currentData() {
    switch (this.selectedView) {
      case 'livros':
        return this.livros;
      case 'usuarios':
        return this.usuarios;
      case 'emprestimos':
        return this.emprestimos;
      default:
        return [];
    }
  }

  carregarLivros() {
    this.loading = true;
    this.http.get(`${this.gatewayUrl}/livros`).subscribe({
      next: (data: any) => {
        this.livros = data;
        this.loading = false;
      },
      error: (error) => {
        console.error('Erro ao carregar livros:', error);
        this.loading = false;
        // Dados de exemplo para teste
        this.livros = [
          { ID: 1, titulo: 'Dom Casmurro', autor: 'Machado de Assis', ano: 1899, disponivel: true },
          { ID: 2, titulo: 'O Cortiço', autor: 'Aluísio Azevedo', ano: 1890, disponivel: false },
          { ID: 3, titulo: '1984', autor: 'George Orwell', ano: 1949, disponivel: true }
        ];
      }
    });
  }

  carregarUsuarios() {
    this.loading = true;
    this.http.get(`${this.gatewayUrl}/usuarios`).subscribe({
      next: (data: any) => {
        this.usuarios = data;
        this.loading = false;
      },
      error: (error) => {
        console.error('Erro ao carregar usuários:', error);
        this.loading = false;
        // Dados de exemplo para teste
        this.usuarios = [
          { ID: 1, nome: 'João Silva', email: 'joao@email.com' },
          { ID: 2, nome: 'Maria Santos', email: 'maria@email.com' },
          { ID: 3, nome: 'Pedro Oliveira', email: 'pedro@email.com' }
        ];
      }
    });
  }

  carregarEmprestimos() {
    this.loading = true;
    this.http.get(`${this.gatewayUrl}/emprestimos`).subscribe({
      next: (data: any) => {
        this.emprestimos = data;
        this.loading = false;
      },
      error: (error) => {
        console.error('Erro ao carregar empréstimos:', error);
        this.loading = false;
        // Dados de exemplo para teste
        this.emprestimos = [
          { 
            ID: 1, 
            livro: { titulo: 'Dom Casmurro' }, 
            usuario: { nome: 'João Silva' },
            data_inicio: '2025-01-15T00:00:00Z',
            data_fim: '2025-01-29T00:00:00Z',
            devolvido: false
          },
          { 
            ID: 2, 
            livro: { titulo: 'O Cortiço' }, 
            usuario: { nome: 'Maria Santos' },
            data_inicio: '2025-01-10T00:00:00Z',
            data_fim: '2025-01-24T00:00:00Z',
            devolvido: true
          }
        ];
      }
    });
  }

  formatDate(dateString: string): string {
    if (!dateString) return 'N/A';
    const date = new Date(dateString);
    return date.toLocaleDateString('pt-BR');
  }
}
