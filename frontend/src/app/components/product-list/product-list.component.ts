import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatTableModule } from '@angular/material/table';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { RouterModule } from '@angular/router';
import { ProductService, Product } from '../../services/product.service';

@Component({
  selector: 'app-product-list',
  standalone: true,
  imports: [
    CommonModule,
    MatTableModule,
    MatCardModule,
    MatButtonModule,
    RouterModule
  ],
  templateUrl: './product-list.component.html',
  styleUrl: './product-list.component.css'
})
export class ProductListComponent implements OnInit {
  products: Product[] = [];
  displayedColumns: string[] = ['id', 'name', 'description', 'balance'];
  isLoading = false;

  constructor(private productService: ProductService) {}

  ngOnInit(): void {
    this.loadProducts();
  }

  loadProducts(): void {
    this.isLoading = true;
    this.productService.getProducts().subscribe({
      next: (response) => {
        setTimeout(() => {
          this.products = response.content;
          this.isLoading = false;
        }, 500);
      },
      error: (error) => {
        console.error('Erro ao carregar produtos:', error);
        this.isLoading = false;
      }
    });
  }
}
