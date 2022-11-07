import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { FormGroup, FormControl, Validators } from '@angular/forms';


@Component({
  selector: 'app-dialog',
  templateUrl: './dialog.component.html',
  styleUrls: ['./dialog.component.scss']
})
export class DialogComponent implements OnInit {

  public formGroup = new FormGroup({
    Id: new FormControl<number|null>(this.data.Id ?? null),
    Title: new FormControl<string|null>(this.data.Title ?? null),
    Author: new FormControl<string|null>(this.data.Author ?? null),
    Year: new FormControl<number|null>(this.data.Year ?? null),
    Publisher: new FormControl<string|null>(this.data.Publisher ?? null),
    Language: new FormControl<string|null>(this.data.Language ?? null),
    Edition: new FormControl<number|null>(this.data.Edition ?? null),
    Pages: new FormControl<number|null>(this.data.Pages ?? null),
    Isbn: new FormControl<string|null>(this.data.Isbn ?? null),
    action: new FormControl<string>(this.data.action),
  });

  constructor(
    public dialogRef: MatDialogRef<DialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: any,
  ) { }

  ngOnInit(): void {}

  onNoClick(): void {
    this.dialogRef.close();
  }
}
