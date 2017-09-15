import auth0 from 'auth0-js';

export default class Auth {
  auth0 = new auth0.WebAuth({
    domain: 'phuocidi.auth0.com',
    clientID: 'bgLMLxaW2c4DKEBGDO0zYh2LXIQSXiSy',
    redirectUri: 'http://localhost:3000',
    audience: 'https://phuocidi.auth0.com/userinfo',
    responseType: 'token id_token',
    scope: 'openid'
  });

  login() {
    this.auth0.authorize();
  }
}
