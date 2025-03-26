import { Component, Output, EventEmitter } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule, FormBuilder, FormGroup, FormArray, Validators } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatSelectModule } from '@angular/material/select';
import { Router } from '@angular/router';
import { InvoiceService, Invoice, InvoiceProduct } from '../../services/invoice.service';

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
    MatSelectModule
  ],
  templateUrl: './invoice-form.component.html',
  styleUrl: './invoice-form.component.css'
})
export class InvoiceFormComponent {
  @Output() invoiceCreated = new EventEmitter<void>();
  
  invoiceForm: FormGroup;
  isLoading = false;

  constructor(
    private fb: FormBuilder,
    private invoiceService: InvoiceService,
    private router: Router
  ) {
    this.invoiceForm = this.fb.group({
      numeration: ['', [Validators.required]],
      products: this.fb.array([])
    });
  }

  get products() {
    return this.invoiceForm.get('products') as FormArray;
  }

  addProduct() {
    const productForm = this.fb.group({
      id: ['', Validators.required],
      quantity: [1, [Validators.required, Validators.min(1)]]
    });

    this.products.push(productForm);
  }

  removeProduct(index: number) {
    this.products.removeAt(index);
  }

  onSubmit(): void {
    if (this.invoiceForm.valid) {
      this.isLoading = true;
      const invoice: Invoice = {
        numeration: this.invoiceForm.value.numeration,
        status: 'OPENED',
        products: this.invoiceForm.value.products
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