import { Component, OnInit } from '@angular/core';
import { AppService } from './service/app.service';
import { Book } from './Book'
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { DialogComponent } from './dialog/dialog.component';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit{
  readonly title: string = 'crud';

  readonly displayedColumns: string[] =
    ['Id', 'Title', 'Author', 'Edition', 'Year', 'Publisher', 'Language', 'Pages', 'Isbn', 'Options'];
  dataSource: Book[] = [];

  constructor(
    private _service: AppService,
    public dialog: MatDialog
  ) {}

  ngOnInit(): void {
    this.get();
  }

  public delete(book: Book): void {
    this._service.delete(book.Id).subscribe(() => this.get());
  }

  private get(): void {
    this._service.get().subscribe(data => this.dataSource = data);
  }

  private post(book: Book) {
    this._service.post(book).subscribe(() => this.get());
  }

  private put(book: Book) {
    this._service.put(book).subscribe(() => this.get());
  }

  public openDialog(book?: Book) {
    const dialogRef = this.dialog.open(DialogComponent, {
      width: '700px',
      data: book ? {...book, action: 'edit'} : {action: 'add'},
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log(result);
      if(result){
        console.log(result.action);
        result.action == 'add' ? this.post(result) : this.put(result);
      }
    });
  }
}
