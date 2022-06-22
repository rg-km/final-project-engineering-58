import axios from "axios";
import React, { Component } from "react";
import { Link } from "react-router-dom";

export default class Login extends Component {
  handleSubmit = (e) => {
    e.preventDefault();

    const data = {
      email: this.email,
      password: this.password,
    };
    axios
      .post(`jwt.SigningMethodHS256, claims`, data)
      .then((res) => {
        console.log(res);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  render() {
    return (
      <div className="App">
        <div className="title">LOGIN</div>
        <div className="mb-3 row">
          <label for="staticEmail" className="col-sm-2 col-form-label">
            Email
          </label>
          <div className="col-sm-10">
            <input type="email" className="form-control" id="staticEmail" onChange={(e) => (this.email = e.target.value)} />
          </div>
        </div>
        <div className="mb-3 row">
          <label for="inputPassword" className="col-sm-2 col-form-label">
            Password
          </label>
          <div className="col-sm-10">
            <input type="password" className="form-control" id="inputPassword" onChange={(e) => (this.password = e.target.value)} />
          </div>
        </div>

        <div className="forgot">
          <a className="link-primary" href="#">
            Forgot Password ?
          </a>
        </div>

        <div className="d-grid gap-2 col-2 mx-auto">
          <button type="button" className="btn btn-outline-dark">
            LOG IN
          </button>
        </div>
        <div className="noaccount">
          Don't have an account ?<br></br>
          <Link to={`/signup`} className="link-primary">
            Sign Up
          </Link>
        </div>
      </div>
    );
  }
}
