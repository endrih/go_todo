import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  public IdToken: string | null = null;
  constructor() { }
}
