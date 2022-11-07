import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Book } from '../Book';

@Injectable({
  providedIn: 'root'
})
export class AppService {

  constructor(private _http: HttpClient) { }

  public delete(id: number): Observable<any> {
    return this._http.delete("http://localhost:8080/book/delete", { body: { Id: id } });
  }

  public get(): Observable<Book[]> {
    return this._http.get<Book[]>("http://localhost:8080/book/get");
  }

  public post(book: Book): Observable<any> {
    return this._http.post<Book>("http://localhost:8080/book/post", book)
  }

  public put(book: Book): Observable<any> {
    return this._http.put<Book>("http://localhost:8080/book/put", book)
  }
}
