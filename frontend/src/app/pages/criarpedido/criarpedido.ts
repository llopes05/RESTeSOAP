import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Navbar } from '../../components/navbar/navbar';

@Component({
  selector: 'app-criarpedido',
  imports: [CommonModule, FormsModule, Navbar],
  templateUrl: './criarpedido.html',
  styleUrl: './criarpedido.css'
})
export class Criarpedido {
  // produtos simulados (futuramente da API)
  produtosEstoque = [
    { id: 1, name: 'Notebook Dell', categoria: 'Eletrônicos', quantidade: 15 },
    { id: 2, name: 'Mouse Logitech', categoria: 'Acessórios', quantidade: 50 },
    { id: 3, name: 'Teclado Mecânico', categoria: 'Acessórios', quantidade: 25 },
    { id: 4, name: 'Monitor 24"', categoria: 'Eletrônicos', quantidade: 8 },
    { id: 5, name: 'Cabo USB-C', categoria: 'Cabos', quantidade: 100 }
  ];

  pedido = {
    destino: '',
    itens: [] as Array<{produto: any, quantidadeSolicitada: number}>
  };

  produtoSelecionado: any = null;
  quantidadeParaAdicionar: number = 1;

  adicionarProdutoAoPedido() {
    if (this.produtoSelecionado && this.quantidadeParaAdicionar > 0) {
      if (this.quantidadeParaAdicionar > this.produtoSelecionado.quantidade) {
        alert(`Quantidade solicitada (${this.quantidadeParaAdicionar}) excede o estoque disponível (${this.produtoSelecionado.quantidade})`);
        return;
      }

      const itemExistente = this.pedido.itens.find(item => item.produto.id === this.produtoSelecionado.id);
      
      if (itemExistente) {
        const novaQuantidade = itemExistente.quantidadeSolicitada + this.quantidadeParaAdicionar;
        if (novaQuantidade > this.produtoSelecionado.quantidade) {
          alert(`Quantidade total (${novaQuantidade}) excederia o estoque disponível (${this.produtoSelecionado.quantidade})`);
          return;
        }
        itemExistente.quantidadeSolicitada = novaQuantidade;
      } else {
        this.pedido.itens.push({
          produto: this.produtoSelecionado,
          quantidadeSolicitada: this.quantidadeParaAdicionar
        });
      }

      this.produtoSelecionado = null;
      this.quantidadeParaAdicionar = 1;
    }
  }

  removerProdutoDoPedido(index: number) {
    this.pedido.itens.splice(index, 1);
  }

  criarPedido() {
    if (!this.pedido.destino.trim()) {
      alert('Por favor, informe o destino do pedido');
      return;
    }

    if (this.pedido.itens.length === 0) {
      alert('Por favor, adicione pelo menos um produto ao pedido');
      return;
    }

    console.log('Pedido criado:', this.pedido);
    alert('Pedido criado com sucesso!');
    
    // integração futura com API:
    /*
    const dadosPedido = {
      destino: this.pedido.destino,
      itens: this.pedido.itens.map(item => ({
        produtoId: item.produto.id,
        quantidade: item.quantidadeSolicitada
      }))
    };

    fetch('/api/pedidos', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(dadosPedido)
    })
    .then(response => response.json())
    .then(resultado => {
      alert(`Pedido #${resultado.id} criado com sucesso!`);
      this.pedido = { destino: '', itens: [] };
    })
    .catch(error => {
      console.error('Erro:', error);
      alert('Erro ao criar pedido. Tente novamente.');
    });
    */
    
    this.pedido = { destino: '', itens: [] };
  }

  calcularTotalItens(): number {
    return this.pedido.itens.reduce((total, item) => total + item.quantidadeSolicitada, 0);
  }
}
