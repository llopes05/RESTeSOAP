import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { Navbar } from '../../components/navbar/navbar';

@Component({
  selector: 'app-criaremprestimo',
  imports: [CommonModule, FormsModule, HttpClientModule, Navbar],
  templateUrl: './criaremprestimo.html',
  styleUrl: './criaremprestimo.css'
})
export class CriarEmprestimo {
  private gatewayUrl = 'http://localhost:8000/api';
  
  // Dados para o formulário
  livros: any[] = [];
  usuarios: any[] = [];
  
  emprestimo = {
    livro_id: '',
    usuario_id: '',
    data_fim: ''
  };

  loading = false;
  mensagem = '';
  sucesso = false;

  constructor(private http: HttpClient) {
    this.carregarLivros();
    this.carregarUsuarios();
    this.definirDataPadrao();
  }

  definirDataPadrao() {
    // Definir data de devolução padrão para 14 dias a partir de hoje
    const dataFim = new Date();
    dataFim.setDate(dataFim.getDate() + 14);
    this.emprestimo.data_fim = dataFim.toISOString().split('T')[0];
  }

  carregarLivros() {
    this.http.get(`${this.gatewayUrl}/livros`).subscribe({
      next: (data: any) => {
        this.livros = data.filter((livro: any) => livro.disponivel);
      },
      error: (error) => {
        console.error('Erro ao carregar livros:', error);
        // Dados de exemplo para teste
        this.livros = [
          { ID: 1, titulo: 'Dom Casmurro', autor: 'Machado de Assis', disponivel: true },
          { ID: 3, titulo: '1984', autor: 'George Orwell', disponivel: true }
        ];
      }
    });
  }

  carregarUsuarios() {
    this.http.get(`${this.gatewayUrl}/usuarios`).subscribe({
      next: (data: any) => {
        this.usuarios = data;
      },
      error: (error) => {
        console.error('Erro ao carregar usuários:', error);
        // Dados de exemplo para teste
        this.usuarios = [
          { ID: 1, nome: 'João Silva', email: 'joao@email.com' },
          { ID: 2, nome: 'Maria Santos', email: 'maria@email.com' }
        ];
      }
    });
  }

  criarEmprestimo() {
    if (!this.emprestimo.livro_id || !this.emprestimo.usuario_id || !this.emprestimo.data_fim) {
      this.sucesso = false;
      this.mensagem = 'Por favor, preencha todos os campos obrigatórios.';
      
      setTimeout(() => {
        this.mensagem = '';
      }, 3000);
      return;
    }

    this.loading = true;
    this.mensagem = '';

    this.http.post(`${this.gatewayUrl}/emprestimos`, this.emprestimo).subscribe({
      next: (response: any) => {
        console.log('Empréstimo criado:', response);
        this.sucesso = true;
        this.mensagem = 'Empréstimo criado com sucesso!';
        
        // Limpar formulário
        this.emprestimo = {
          livro_id: '',
          usuario_id: '',
          data_fim: ''
        };
        this.definirDataPadrao();
        
        this.loading = false;
        
        // Recarregar livros para atualizar disponibilidade
        this.carregarLivros();
        
        // Limpar mensagem após 3 segundos
        setTimeout(() => {
          this.mensagem = '';
          this.sucesso = false;
        }, 3000);
      },
      error: (error: any) => {
        console.error('Erro ao criar empréstimo:', error);
        this.sucesso = false;
        this.mensagem = 'Erro ao criar empréstimo. Tente novamente.';
        this.loading = false;
        
        // Limpar mensagem de erro após 5 segundos
        setTimeout(() => {
          this.mensagem = '';
        }, 5000);
      }
    });
  }

  getLivroSelecionado() {
    return this.livros.find(livro => livro.ID == this.emprestimo.livro_id);
  }

  getUsuarioSelecionado() {
    return this.usuarios.find(usuario => usuario.ID == this.emprestimo.usuario_id);
  }
}
