import jQuery from 'jquery'
import React from 'react'
import { Component, PropTypes } from 'react'
import { render } from 'react-dom'
import { App } from './components/App'


window.$ = window.jQuery = jQuery;
window.React = React


render(<App />, document.getElementById("app"));
