
import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { Navbar } from '../../components/navbar/navbar';
import { CardModule } from 'primeng/card';
import { InputTextModule } from 'primeng/inputtext';
import { ButtonModule } from 'primeng/button';
import { MessageModule } from 'primeng/message';

@Component({
  selector: 'app-criarusuario',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    HttpClientModule,
    Navbar,
    CardModule,
    InputTextModule,
    ButtonModule,
    MessageModule
  ],
  templateUrl: './criarusuario.html',
  styleUrl: './criarusuario.css'
})
export class Criarusuario {
  private gatewayUrl = 'http://localhost:8000/api';

  usuario = {
    nome: '',
    email: ''
  };

  loading = false;
  mensagem = '';
  sucesso = false;

  constructor(private http: HttpClient) {}

  criarUsuario() {
    if (!this.usuario.nome || !this.usuario.email) {
      this.sucesso = false;
      this.mensagem = 'Por favor, preencha todos os campos obrigatórios.';
      setTimeout(() => { this.mensagem = ''; }, 3000);
      return;
    }
    this.loading = true;
    this.mensagem = '';
    this.http.post(`${this.gatewayUrl}/usuarios/`, this.usuario).subscribe({
      next: (response: any) => {
        this.sucesso = true;
        this.mensagem = 'Usuário criado com sucesso!';
        this.usuario = { nome: '', email: '' };
        this.loading = false;
        setTimeout(() => {
          this.mensagem = '';
          this.sucesso = false;
        }, 3000);
      },
      error: (error: any) => {
        console.error('Erro ao criar usuário:', error);
        this.sucesso = false;
        this.mensagem = 'Erro ao criar usuário. Tente novamente.';
        this.loading = false;
        setTimeout(() => { this.mensagem = ''; }, 5000);
      }
    });
  }
}
