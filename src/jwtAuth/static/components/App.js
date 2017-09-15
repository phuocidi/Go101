import { Component, PropTypes } from 'react'
import { LoggedIn } from './LoggedIn'
import { Home } from './Home'

export class App extends Component {
  componentWillMount() {
  }

  render(){

    if (this.loggedIn) {
      return (<LoggedIn />);
    } else {
      return (<Home />);
    }
  }
}