import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { TableModule } from 'primeng/table';
import { SelectButtonModule } from 'primeng/selectbutton';
import { InputTextModule } from 'primeng/inputtext';
import { IconFieldModule } from 'primeng/iconfield';
import { InputIconModule } from 'primeng/inputicon';
import { Navbar } from '../../components/navbar/navbar';

@Component({
  selector: 'app-home',
  imports: [CommonModule, TableModule, SelectButtonModule, FormsModule, InputTextModule, IconFieldModule, InputIconModule, Navbar],
  templateUrl: './home.html',
  styleUrl: './home.css'
})
export class Home {
  // Opções de visualização
  viewOptions = [
    { label: 'Estoque', value: 'estoque', icon: 'pi pi-warehouse' },
    { label: 'Pedidos', value: 'pedidos', icon: 'pi pi-truck' }
  ];
  
  selectedView = 'estoque';

  products = [
    { id: 1, name: 'Notebook Dell', categoria: 'Eletrônicos', quantidade: 15 },
    { id: 2, name: 'Mouse Logitech', categoria: 'Acessórios', quantidade: 50 },
    { id: 3, name: 'Teclado Mecânico', categoria: 'Acessórios', quantidade: 25 },
    { id: 4, name: 'Monitor 24"', categoria: 'Eletrônicos', quantidade: 8 },
    { id: 5, name: 'Cabo USB-C', categoria: 'Cabos', quantidade: 100 }
  ];

  deliveries = [
    { 
      id: 1, 
      para: 'São Paulo - Escritório Central', 
      items: [
        { name: 'Notebook Dell', quantidade: 3 },
        { name: 'Mouse Logitech', quantidade: 5 }
      ]
    },
    { 
      id: 2, 
      para: 'Rio de Janeiro - Filial Sul', 
      items: [
        { name: 'Monitor 24"', quantidade: 2 },
        { name: 'Teclado Mecânico', quantidade: 4 }
      ]
    },
    { 
      id: 3, 
      para: 'Belo Horizonte - Centro de Distribuição', 
      items: [
        { name: 'Cabo USB-C', quantidade: 20 },
        { name: 'Mouse Logitech', quantidade: 10 }
      ]
    }
  ];

  get currentData(): any[] {
    return this.selectedView === 'estoque' ? this.products : this.deliveries;
  }
}
