import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { InputTextModule } from 'primeng/inputtext';
import { ButtonModule } from 'primeng/button';
import { CardModule } from 'primeng/card';
import { InputNumberModule } from 'primeng/inputnumber';
import { FloatLabelModule } from 'primeng/floatlabel';
import { IconFieldModule } from 'primeng/iconfield';
import { InputIconModule } from 'primeng/inputicon';
import { MessageModule } from 'primeng/message';
import { Navbar } from '../../components/navbar/navbar';

@Component({
  selector: 'app-adicionarlivro',
  standalone: true, // Adicionado standalone: true
  imports: [
    CommonModule, 
    FormsModule, 
    InputTextModule, 
    ButtonModule, 
    CardModule, 
    InputNumberModule, 
    FloatLabelModule,
    IconFieldModule,
    InputIconModule,
    MessageModule,
    HttpClientModule,
    Navbar
  ],
  templateUrl: './adicionarlivro.html',
  styleUrl: './adicionarlivro.css'
})
export class AdicionarLivro {
  private gatewayUrl = 'http://localhost:8000/api';
  
  livro = {
    titulo: '',
    autor: '',
    ano: new Date().getFullYear(),
    disponivel: true
  };

  loading = false;
  mensagem = '';
  sucesso = false;

  constructor(private http: HttpClient) {}

  // ✅ NOVO MÉTODO PARA LIMPAR O FORMULÁRIO
  limparFormulario() {
    this.livro = {
      titulo: '',
      autor: '',
      ano: new Date().getFullYear(),
      disponivel: true
    };
  }

  adicionarLivro() {
    if (this.livro.titulo && this.livro.autor && this.livro.ano) {
      this.loading = true;
      this.mensagem = '';
      
      this.http.post(`${this.gatewayUrl}/livros/`, this.livro).subscribe({
        next: (response: any) => {
          console.log('Livro adicionado:', response);
          this.sucesso = true;
          this.mensagem = 'Livro adicionado com sucesso!';
          
          // Usar o novo método para limpar
          this.limparFormulario();
          
          this.loading = false;
          
          setTimeout(() => {
            this.mensagem = '';
            this.sucesso = false;
          }, 3000);
        },
        error: (error) => {
          console.error('Erro ao adicionar livro:', error);
          this.sucesso = false;
          this.mensagem = 'Erro ao adicionar livro. Tente novamente.';
          this.loading = false;
          
          setTimeout(() => {
            this.mensagem = '';
          }, 5000);
        }
      });
    } else {
      this.sucesso = false;
      this.mensagem = 'Por favor, preencha todos os campos obrigatórios.';
      
      setTimeout(() => {
        this.mensagem = '';
      }, 3000);
    }
  }
}