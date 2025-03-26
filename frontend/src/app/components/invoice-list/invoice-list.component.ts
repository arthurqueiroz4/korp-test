import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatTableModule } from '@angular/material/table';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { RouterModule } from '@angular/router';
import { InvoiceService, Invoice } from '../../services/invoice.service';

@Component({
  selector: 'app-invoice-list',
  standalone: true,
  imports: [
    CommonModule,
    MatTableModule,
    MatCardModule,
    MatButtonModule,
    MatIconModule,
    RouterModule
  ],
  templateUrl: './invoice-list.component.html',
  styleUrl: './invoice-list.component.css'
})
export class InvoiceListComponent implements OnInit {
  invoices: Invoice[] = [];
  displayedColumns: string[] = ['id', 'numeration', 'status', 'productsCount'];
  isLoading = false;

  constructor(private invoiceService: InvoiceService) {}

  ngOnInit(): void {
    this.loadInvoices();
  }

  loadInvoices(): void {
    this.isLoading = true;
    this.invoiceService.getInvoices().subscribe({
      next: (response) => {
        this.invoices = response.content;
        this.isLoading = false;
      },
      error: (error) => {
        console.error('Erro ao carregar notas fiscais:', error);
        this.isLoading = false;
      }
    });
  }

  getStatusClass(status: string): string {
    return {
      'OPENED': 'status-opened',
      'PROCESSING': 'status-processing',
      'CLOSED': 'status-closed'
    }[status] || '';
  }

  getStatusText(status: string): string {
    return {
      'OPENED': 'Em Aberto',
      'PROCESSING': 'Processando',
      'CLOSED': 'Concluída'
    }[status] || status;
  }
} 