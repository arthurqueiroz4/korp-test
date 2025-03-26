import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatButtonToggleModule } from '@angular/material/button-toggle';
import { MatTabsModule } from '@angular/material/tabs';
import { RouterModule } from '@angular/router';
import { ProductListComponent } from '../product-list/product-list.component';
import { ProductFormComponent } from '../product-form/product-form.component';
import { InvoiceListComponent } from '../invoice-list/invoice-list.component';
import { InvoiceFormComponent } from '../invoice-form/invoice-form.component';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [
    CommonModule,
    MatCardModule,
    MatButtonModule,
    MatButtonToggleModule,
    MatTabsModule,
    RouterModule,
    ProductListComponent,
    ProductFormComponent,
    InvoiceListComponent,
    InvoiceFormComponent
  ],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent {
  selectedTabIndex = 0;
  selectedOption: 'list' | 'create' = 'list';
  selectedInvoiceOption: 'list' | 'create' = 'list';

  onTabChange(index: number): void {
    this.selectedTabIndex = index;
  }

  onOptionChange(option: 'list' | 'create'): void {
    this.selectedOption = option;
  }

  onInvoiceOptionChange(option: 'list' | 'create'): void {
    this.selectedInvoiceOption = option;
  }

  onProductCreated(): void {
    this.selectedOption = 'list';
  }

  onInvoiceCreated(): void {
    this.selectedInvoiceOption = 'list';
  }
}
