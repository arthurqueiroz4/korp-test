import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

export interface Product {
  id?: number;
  name: string;
  description: string;
  balance: number;
}

export interface ProductResponse {
  content: Product[];
  page: number;
  size: number;
}

@Injectable({
  providedIn: 'root'
})
export class ProductService {
  private apiUrl = 'http://localhost:8000/api';

  constructor(private http: HttpClient) { }

  getProducts(page: number = 0, size: number = 10): Observable<ProductResponse> {
    return this.http.get<ProductResponse>(`${this.apiUrl}/products?page=${page}&size=${size}`);
  }

  createProduct(product: Product): Observable<Product> {
    return this.http.post<Product>(`${this.apiUrl}/products`, product);
  }
}
