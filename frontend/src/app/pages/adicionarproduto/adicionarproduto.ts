import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { InputTextModule } from 'primeng/inputtext';
import { SelectModule } from 'primeng/select';
import { ButtonModule } from 'primeng/button';
import { CardModule } from 'primeng/card';
import { InputNumberModule } from 'primeng/inputnumber';
import { FloatLabelModule } from 'primeng/floatlabel';
import { IconFieldModule } from 'primeng/iconfield';
import { InputIconModule } from 'primeng/inputicon';
import { Navbar } from '../../components/navbar/navbar';

@Component({
  selector: 'app-adicionarproduto',
  imports: [
    CommonModule, 
    FormsModule, 
    InputTextModule, 
    SelectModule, 
    ButtonModule, 
    CardModule, 
    InputNumberModule, 
    FloatLabelModule,
    IconFieldModule,
    InputIconModule,
    Navbar
  ],
  templateUrl: './adicionarproduto-test.html',
  styleUrl: './adicionarproduto.css'
})
export class Adicionarproduto {
  produto = {
    name: '',
    quantidade: 0,
    categoria: ''
  };

  categorias = [
    { label: 'Eletrônicos', value: 'Eletrônicos' },
    { label: 'Acessórios', value: 'Acessórios' },
    { label: 'Cabos', value: 'Cabos' },
    { label: 'Periféricos', value: 'Periféricos' },
    { label: 'Hardware', value: 'Hardware' },
    { label: 'Software', value: 'Software' },
    { label: 'Móveis', value: 'Móveis' },
    { label: 'Escritório', value: 'Escritório' }
  ];

  adicionarProduto() {
    if (this.produto.name && this.produto.quantidade >= 0 && this.produto.categoria) {
      console.log('Produto adicionado:', this.produto);
      
      this.produto = {
        name: '',
        quantidade: 0,
        categoria: ''
      };
      
      alert('Produto adicionado com sucesso!');
    } else {
      alert('Por favor, preencha todos os campos obrigatórios.');
    }
  }

  // integração futura com API:
  /*
  async adicionarProduto() {
    if (this.produto.name && this.produto.quantidade >= 0 && this.produto.categoria) {
      try {
        const response = await fetch('/api/produtos', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(this.produto)
        });

        if (response.ok) {
          const resultado = await response.json();
          console.log('Produto salvo no backend:', resultado);
          
          this.produto = { name: '', quantidade: 0, categoria: '' };
          alert('Produto adicionado com sucesso!');
        } else {
          throw new Error('Erro ao salvar produto');
        }
      } catch (error) {
        console.error('Erro:', error);
        alert('Erro ao adicionar produto. Tente novamente.');
      }
    } else {
      alert('Por favor, preencha todos os campos obrigatórios.');
    }
  }
  */
}
