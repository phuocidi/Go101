import { Component } from 'react'
import  {Product } from './Product'
import history from '../history';


export class LoggedIn extends Component{
  constructor(props) {
  	super(props)
    this.logout = this.logout.bind(this)
  	this.state = {
  		products: []
  	}
  }
  componentDidMount() {
    this.serverRequest = $.get('http://localhost:3000/products', function (result) {
      this.setState( {
        products: results,
      });
    }.bind(this));
  }



  render() {
    return (
      <div className="col-lg-12">
      <a onClick={this.logout} className="btn btn-info btn-lg">
          <span className="glyphicon glyphicon-log-out"></span> Log out
      </a>
        <h2>Welcome to We R VR</h2>
        <p>Below you'll find the latest games that need feedback. Please provide honest feedback so developers can make the best games.</p>
        <div className="row">

        {this.state.products.map(function(product, i){
          return <Product key={i} product={product} />
        })}
        </div>
      </div>);
  }

  logout () {
    localStorage.removeItem('id_token');
    localStorage.removeItem('access_token');
    localStorage.removeItem('profile');
    history.replace();
  }
}