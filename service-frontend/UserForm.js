import React, { Component } from 'react';
import {connect} from 'react-redux';

class UserForm extends Component {
    handleSubmit = (e) => {
        e.preventDefault();
          const firstName = this.getFirstName.value;
          const lastName = this.getLastName.value;
          const email =  this.getEmail.value;
          const password =  this.getPassword.value;
          const data = {
            id: new Date(),
            firstName,
            lastName,
            email,
            password
        }
        console.log(data)

        this.props.dispatch({
          type:'ADD_USER',
      data});
    this.getFirstName.value = '';
    this.getLastName.value = '';
    this.getEmail.value = '';
    this.getPassword.value = '';
    
    }

    render() {
        return (

        <div className="post-container">
          <h1 className="post_heading">Signup</h1>
          <form className="form" onSubmit={this.handleSubmit} >
            <input required type="text" ref={(input) => this.getFirstName = input}
                placeholder="Enter First Name" /><br /><br />
            <input required type="text" ref={(input) => this.getLastName = input}
                placeholder="Enter Last Name" /><br /><br />
            <input required type="text" ref={(input) => this.getEmail = input}
                cols="28" placeholder="Enter User Email" /><br /><br />
            <input required type="text" ref={(input) => this.getPassword = input}
                placeholder="Enter Password" /><br /><br />
            <button>Submit</button>
          </form>
        </div>
        );
        }
        }
        
        const mapStateToProps = (state) => {
          return {
              users: state
          }
        }
        
        export default connect(mapStateToProps)(UserForm);