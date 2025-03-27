import { Component, OnInit  } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatTableModule } from '@angular/material/table';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatPaginatorModule, PageEvent } from '@angular/material/paginator';
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
    MatPaginatorModule,
    RouterModule
  ],
  templateUrl: './invoice-list.component.html',
  styleUrl: './invoice-list.component.css'
})
export class InvoiceListComponent implements OnInit {
  invoices: Invoice[] = [];
  displayedColumns: string[] = ['id', 'numeration', 'status', 'productsCount', 'actions'];
  isLoading = false;

  // Paginação
  totalItems = 0;
  pageSize = 10;
  currentPage = 0;
  pageSizeOptions = [5, 10, 25, 50];

  constructor(
    private invoiceService: InvoiceService,
  ) {}

  ngOnInit(): void {
    this.loadInvoices();
  }

  loadInvoices(): void {
    this.isLoading = true;
    this.invoiceService.getInvoices(this.currentPage, this.pageSize).subscribe({
      next: (response) => {
        this.invoices = response.content;
        this.totalItems = response.total;
        this.isLoading = false;
      },
      error: (error) => {
        console.error('Erro ao carregar notas fiscais:', error);
        this.isLoading = false;
      }
    });
  }

  onPageChange(event: PageEvent): void {
    this.currentPage = event.pageIndex;
    this.pageSize = event.pageSize;
    this.loadInvoices();
  }

  getStatusClass(status: string): string {
    return {
      'OPENED': 'status-opened',
      'PROCESSING': 'status-processing',
      'CLOSED': 'status-closed',
      'FAILED': 'status-failed'
    }[status] || '';
  }

  getStatusText(status: string): string {
    return {
      'OPENED': 'Em Aberto',
      'PROCESSING': 'Em Processamento',
      'CLOSED': 'Concluída',
      'FAILED': 'Falhou'
    }[status] || status;
  }

  canProcess(invoice: Invoice): boolean {
    return invoice.status === 'OPENED' || invoice.status === 'FAILED';
  }

  processInvoice(invoice: Invoice): void {
    if (invoice.id) {
      this.invoiceService.enqueueInvoice(invoice.id).subscribe({
        next: () => {
          invoice.status = 'PROCESSING';
        },
        error: (error) => {
          console.error('Erro ao processar nota fiscal:', error);
        }
      });
    }
  }
} 