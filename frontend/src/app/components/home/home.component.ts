import { Component, ViewChild } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatButtonToggleModule } from '@angular/material/button-toggle';
import { MatTabsModule } from '@angular/material/tabs';
import { MatIconModule } from '@angular/material/icon';
import { MatTooltipModule } from '@angular/material/tooltip';
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
    MatIconModule,
    MatTooltipModule,
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
  @ViewChild(ProductListComponent) productList?: ProductListComponent;
  @ViewChild(InvoiceListComponent) invoiceList?: InvoiceListComponent;

  selectedTabIndex = 0;
  selectedOption: 'list' | 'create' = 'list';
  selectedInvoiceOption: 'list' | 'create' = 'list';
  isRefreshing = false;

  refreshTables(): void {
    this.isRefreshing = true;
    
    if (this.selectedOption === 'list') {
      this.productList?.loadProducts();
    }
    if (this.selectedInvoiceOption === 'list') {
      this.invoiceList?.loadInvoices();
    }

    setTimeout(() => {
      this.isRefreshing = false;
    }, 800);
  }

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
