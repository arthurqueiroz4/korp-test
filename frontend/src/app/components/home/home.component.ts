import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatButtonToggleModule } from '@angular/material/button-toggle';
import { RouterModule } from '@angular/router';
import { ProductListComponent } from '../product-list/product-list.component';
import { ProductFormComponent } from '../product-form/product-form.component';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [
    CommonModule,
    MatCardModule,
    MatButtonModule,
    MatButtonToggleModule,
    RouterModule,
    ProductListComponent,
    ProductFormComponent
  ],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent {
  selectedOption: 'list' | 'create' = 'list';

  onOptionChange(option: 'list' | 'create'): void {
    this.selectedOption = option;
  }
}
