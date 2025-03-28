import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Product } from './product.service';

export interface InvoiceProduct {
  id?: number;
  name: string;
  quantity: number;
}

export interface Invoice {
  id?: number;
  numeration: string;
  status: "OPENED" | "CLOSED" | "PROCESSING" | "FAILED";
  products: InvoiceProduct[];
}

export interface InvoiceResponse {
  content: Invoice[];
  page: number;
  size: number;
  total: number;
}

@Injectable({
  providedIn: 'root'
})
export class InvoiceService {
  private apiUrl = 'http://localhost:9000/api';

  constructor(private http: HttpClient) { }

  getInvoices(page: number = 0, size: number = 10): Observable<InvoiceResponse> {
    return this.http.get<InvoiceResponse>(`${this.apiUrl}/invoices?page=${page}&size=${size}`);
  }

  createInvoice(invoice: Invoice): Observable<Invoice> {
    return this.http.post<Invoice>(`${this.apiUrl}/invoices`, invoice);
  }

  enqueueInvoice(id: number): Observable<void> {
    return this.http.post<void>(`${this.apiUrl}/invoices/enqueue/${id}`, {});
  }
} 