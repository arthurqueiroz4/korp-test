import { Component, Output, EventEmitter } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { Router, RouterModule } from '@angular/router';
import { ProductService } from '../../services/product.service';

@Component({
  selector: 'app-product-form',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatCardModule,
    RouterModule
  ],
  templateUrl: './product-form.component.html',
  styleUrl: './product-form.component.css'
})
export class ProductFormComponent {
  @Output() productCreated = new EventEmitter<void>();
  
  productForm: FormGroup;
  isLoading = false;

  constructor(
    private fb: FormBuilder,
    private productService: ProductService,
    private router: Router
  ) {
    this.productForm = this.fb.group({
      name: ['', [Validators.required, Validators.maxLength(128)]],
      description: ['', [Validators.required, Validators.maxLength(256)]],
      balance: [0, [Validators.required, Validators.min(0)]]
    });
  }

  onSubmit(): void {
    if (this.productForm.valid) {
      this.isLoading = true;
      this.productService.createProduct(this.productForm.value).subscribe({
        next: () => {
          this.productCreated.emit();
        },
        error: (error) => {
          console.error('Erro ao criar produto:', error);
          this.isLoading = false;
        }
      });
    }
  }
}
