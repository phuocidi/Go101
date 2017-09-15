import React from 'react'
import { Component, PropTypes } from 'react'
import { render } from 'react-dom'
import { App } from './components/App'
window.React = React


render(<App />, document.getElementById("app"));
