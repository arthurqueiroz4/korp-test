import { Component, Output, EventEmitter } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule, FormBuilder, FormGroup, FormArray, Validators } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatSelectModule } from '@angular/material/select';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { Router } from '@angular/router';
import { InvoiceService, Invoice, InvoiceProduct } from '../../services/invoice.service';
import { ProductService, Product, ProductResponse } from '../../services/product.service';
import { Observable, Subject } from 'rxjs';
import { debounceTime, distinctUntilChanged, switchMap, map } from 'rxjs/operators';

@Component({
  selector: 'app-invoice-form',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatIconModule,
    MatSelectModule,
    MatAutocompleteModule
  ],
  templateUrl: './invoice-form.component.html',
  styleUrl: './invoice-form.component.css'
})
export class InvoiceFormComponent {
  @Output() invoiceCreated = new EventEmitter<void>();
  
  invoiceForm: FormGroup;
  isLoading = false;
  filteredProducts: Observable<Product[]> = new Observable();
  private searchTerms = new Subject<string>();
  productBalances: { [key: number]: number } = {};

  constructor(
    private fb: FormBuilder,
    private invoiceService: InvoiceService,
    private productService: ProductService,
    private router: Router
  ) {
    this.invoiceForm = this.fb.group({
      numeration: ['', [Validators.required]],
      products: this.fb.array([])
    });

    this.setupProductSearch();
  }

  private setupProductSearch(): void {
    this.filteredProducts = this.searchTerms.pipe(
      debounceTime(300),
      distinctUntilChanged(),
      switchMap(term => this.productService.searchProductsByName(term)),
      map(response => response.content)
    );
  }

  searchProducts(event: Event): void {
    const input = event.target as HTMLInputElement;
    this.searchTerms.next(input.value);
  }

  displayProductName = (product: Product | null): string => {
    return product ? product.name : '';
  }

  get products() {
    return this.invoiceForm.get('products') as FormArray;
  }

  addProduct() {
    const productForm = this.fb.group({
      id: ['', Validators.required],
      name: ['', Validators.required],
      quantity: [1, [Validators.required, Validators.min(1)]],
      selectedProduct: [null]
    });

    this.products.push(productForm);
  }

  removeProduct(index: number) {
    this.products.removeAt(index);
  }

  onProductSelected(product: Product, index: number) {
    const productForm = this.products.at(index);
    this.productBalances[product.id!] = product.balance;
    
    productForm.patchValue({
      id: product.id,
      name: product.name,
      selectedProduct: product
    });

    productForm.get('quantity')?.setValidators([
      Validators.required,
      Validators.min(1),
      Validators.max(product.balance)
    ]);
    
    productForm.get('quantity')?.updateValueAndValidity();
  }

  getMaxQuantity(index: number): number {
    const productId = this.products.at(index).get('id')?.value;
    return this.productBalances[productId] || 0;
  }

  onSubmit(): void {
    if (this.invoiceForm.valid) {
      this.isLoading = true;
      const invoice: Invoice = {
        numeration: this.invoiceForm.value.numeration,
        status: 'OPENED',
        products: this.invoiceForm.value.products.map((p: any) => ({
          id: p.id,
          name: p.name,
          quantity: p.quantity
        }))
      };

      this.invoiceService.createInvoice(invoice).subscribe({
        next: () => {
          this.invoiceCreated.emit();
        },
        error: (error) => {
          console.error('Erro ao criar nota fiscal:', error);
          this.isLoading = false;
        }
      });
    }
  }
} 