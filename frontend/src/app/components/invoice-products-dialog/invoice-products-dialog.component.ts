import { Component, Inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatDialogModule, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatTableModule } from '@angular/material/table';
import { Invoice, InvoiceProduct } from '../../services/invoice.service';

@Component({
  selector: 'app-invoice-products-dialog',
  standalone: true,
  imports: [CommonModule, MatDialogModule, MatTableModule],
  template: `
    <h2 mat-dialog-title>Produtos da Nota Fiscal</h2>
    <mat-dialog-content>
      <table mat-table [dataSource]="data.products">
        <ng-container matColumnDef="id">
          <th mat-header-cell *matHeaderCellDef>ID</th>
          <td mat-cell *matCellDef="let product">{{product.id}}</td>
        </ng-container>

        <ng-container matColumnDef="name">
          <th mat-header-cell *matHeaderCellDef>Nome</th>
          <td mat-cell *matCellDef="let product">{{product.name}}</td>
        </ng-container>

        <ng-container matColumnDef="quantity">
          <th mat-header-cell *matHeaderCellDef>Quantidade</th>
          <td mat-cell *matCellDef="let product">{{product.quantity}}</td>
        </ng-container>

        <tr mat-header-row *matHeaderRowDef="displayedColumns"></tr>
        <tr mat-row *matRowDef="let row; columns: displayedColumns;"></tr>
      </table>
    </mat-dialog-content>
    <mat-dialog-actions align="end">
      <button mat-button mat-dialog-close>Fechar</button>
    </mat-dialog-actions>
  `
})
export class InvoiceProductsDialogComponent {
  displayedColumns: string[] = ['id', 'name', 'quantity'];

  constructor(
    public dialogRef: MatDialogRef<InvoiceProductsDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: Invoice
  ) {}
} 