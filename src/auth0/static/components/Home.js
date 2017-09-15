import { Component, PropTypes } from 'react'
import auth0 from 'auth0-js'
// import Auth from '../Auth.js';
import history from '../history';

export class Home extends Component{
  constructor(props) {
    super(props)
    
    this.authenticate = this.authenticate.bind(this)

    this.AUTH0_API_AUDIENCE  = process.env.AUTH0_API_AUDIENCE
    this.AUTH0_DOMAIN        = process.env.AUTH0_DOMAIN
    this.AUTH0_CLIENT_ID     = process.env.AUTH0_CLIENT_ID
    this.AUTH0_CALLBACK_URL  = process.env.AUTH0_CALLBACK_URL
  }
    authenticate() {
    this.webAuth = new auth0.WebAuth({
      domain: this.AUTH0_DOMAIN,
      clientID: this.AUTH0_CLIENT_ID,
      redirectUri: this.AUTH0_CALLBACK_URL,
      audience: this.AUTH0_API_AUDIENCE,
      responseType: 'token id_token',
      scope: 'openid'
    });  

    this.webAuth.authorize();


 //   this.auth = new Auth();
 //   auth.login();
  }

  render() {
    return (
    <div className="container">
      <div className="col-xs-12 jumbotron text-center">
        <h1>We R VR TRAN</h1>
        <p>Provide valuable feedback to VR experience developers.</p>
        <a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>

      </div>
    </div>);
  }


}
