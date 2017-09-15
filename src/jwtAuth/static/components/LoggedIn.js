import { Component } from 'react'
import  {Product } from './Product'


export class LoggedIn extends Component{
  constructor(props) {
  	super(props)
  	this.state = {
  		products: []
  	}
  }

  render() {
    return (
      <div className="col-lg-12">
        <span className="pull-right"><a onClick={this.logout}>Log out</a></span>
        <h2>Welcome to We R VR</h2>
        <p>Below you'll find the latest games that need feedback. Please provide honest feedback so developers can make the best games.</p>
        <div className="row">

        {this.state.products.map(function(product, i){
          return <Product key={i} product={product} />
        })}
        </div>
      </div>);
  }
}