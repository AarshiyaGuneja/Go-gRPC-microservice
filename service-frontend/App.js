
import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import UserForm from './UserForm';
import SignIn from './Signin';

class App extends Component {
  render() {
      return (
          <Router>
          <h2 className="center ">Fleet-Management</h2>
      <div>
        <ul>
          <li>
            <Link to="/signup">SignUp</Link>
          </li>
          <li>
            <Link to="/login">LogIn</Link>
          </li>
        </ul>
        <hr />

        <Route path="/signup" component={UserForm} />
        <Route path="/login" component={SignIn} />
      </div>
    </Router>
      );
  }
  }
export default App;



