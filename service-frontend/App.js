
import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import UserForm from './UserForm';
import LoginForm from './LoginForm';

class App extends Component {
  render() {
      return (
          <Router>
            <div className="App">
        <div className="navbar">
          <h2 className="center ">Fleet-Management</h2>
      <div>
        <ul>
          <li>
            <Link to="/signup">Signup</Link>
          </li>
          <li>
            <Link to="/login">Login</Link>
          </li>
        </ul>

        <hr />

        <Route path="/signup" component={UserForm} />
        <Route path="/login" component={LoginForm} />
      </div>
      </div>
      </div>
    </Router>
      );
  }
  }
export default App;


