import { Component, PropTypes } from 'react'
import { LoggedIn } from './LoggedIn'
import { Home } from './Home'
import auth0 from 'auth0-js'
import history from '../history';


let AUTH0_DOMAIN 		= process.env.AUTH0_DOMAIN
let AUTH0_CLIENT_ID		= process.env.AUTH0_CLIENT_ID

export class App extends Component {
  constructor  (props) {
  		super(props)
  }

  componentWillMount() {
  	//this.setupAjax()
    this.parseHash();
    this.setState();
  }

  render(){

    if (this.loggedIn) {
      return (<LoggedIn />);
    } else {
      return (<Home />);
    }
  }


  setupAjax() {
 	$.ajaxSetup({
 		beforeSend: (xhr) => {
 			if (localStorage.getItem('access_token')) {
 				xhr.setRequestHeader('Authorization', 'BEARER ' + localStorage.getItem('access_token') )
 			}
 		}
 	})
  }

  parseHash () {
  	this.auth0 = new auth0.WebAuth({
  		domain: AUTH0_DOMAIN,
  		clientID: AUTH0_CLIENT_ID
  	});

  	this.auth0.parseHash(window.location.hash, function (err, authResult) {
  		if (err) {
  			return console.log(err);
  		}

  		//console.log(authResult);
  		if (authResult && authResult.accessToken && authResult.idToken) {
  			let expiresAt = JSON.stringify((authResult.expiresIn * 1000) + new Date().getTime());

        localStorage.setItem('access_token', authResult.accessToken);
  			localStorage.setItem('id_token', authResult.idToken);
  			localStorage.setItem('profile', JSON.stringify(authResult.idTokenPayload))
  		  localStorage.setItem('expires_at', expiresAt)
      }
  	});
  }

  // Set user login state
  setState() {
  	var idToken = localStorage.getItem('id_token')
  	if(idToken) {
  		this.loggedIn = true;
  	}else {
  		this.loggedIn = false;
  	}
  }

}