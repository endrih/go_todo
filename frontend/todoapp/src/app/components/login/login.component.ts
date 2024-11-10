import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent implements OnInit {

  public id_token: string | null = null;
  constructor(private route: ActivatedRoute){
  }

  ngOnInit(){
    this.id_token = this.route.snapshot.queryParamMap.get('id_token');
  }
}
