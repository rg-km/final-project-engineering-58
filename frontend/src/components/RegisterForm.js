import React, { useState } from "react";
import "../assets/css/RegisterForm.css"
import { Link } from "react-router-dom"
import axios from "axios"


const RegisterForm = () =>{
  const [username,setUsername] = useState("");
  const [email,setEmail] = useState("");
  const [password,setPassword] = useState("");

  const onChangeUsername = (e) => {
    const value = e.target.value
    setUsername(value)
  }

  const onChangeEmail = (e) => {
    const value = e.target.value
    setEmail(value)
  }

  const onChangePassword = (e) => {
    const value = e.target.value
    setPassword(value)
  }

  const KlikDaftar = () => {
    const data = {
      username: username,
      email : email,
      password : password
    }
    axios.post('http://localhost:9500/api/auth/register', data)
    .then(result =>{
      if(result){
        if(result.data){

        }
      }
    })
  }

  return (
    <div class="">
      <div class="mask py-5 d-flex align-items-center h-100 gradient-custom-3">
        <div class="container h-100">
          <div class="row d-flex justify-content-center align-items-center h-100">
            <div class="col-12 col-md-9 col-lg-7 col-xl-6">
              <div class="card">
                <div class="card-body p-5">
                  <h2 class="text-uppercase text-center mb-5">Create an account</h2>

                  <form>

                 <div class="form-floating mb-3">
              <input type="email" class="form-control" id="floatingInput" placeholder="username" value={username} onChange={onChangeUsername}/>
              <label for="floatingInput">User name</label>
              </div>
                <div class="form-floating mb-3">
              <input type="email" class="form-control" id="floatingInput" placeholder="name@example.com" value={email} onChange={onChangeEmail}/>
              <label for="floatingInput">Email address</label>
              </div>
              <div class="form-floating">
              <input type="password" class="form-control" id="floatingPassword" placeholder="Password" value={password} onChange={onChangePassword}/>
              <label for="floatingPassword">Password</label>
              </div>
              <br></br>

                <div class="form-check d-flex justify-content-center mb-5">
                  <input class="form-check-input me-2" type="checkbox" value="" id="form2Example3cg" />
                  <label class="form-check-label" for="form2Example3g">
                    I agree all statements in <a href="#!" class="text-body"><u>Terms of service</u></a>
                  </label>
                </div>

                <div class="d-flex justify-content-center">
                  <Link to="/kelas-saya">
                    <button type="button" class="btn btn-success btn-block btn-lg gradient-custom-4 text-body" onClick={KlikDaftar}>
                      Register
                    </button>
                  </Link>
                </div>

                <p class="text-center text-muted mt-5 mb-0">Have already an account? <Link to="/login"
                    class="fw-bold text-body"><u>Login here</u></Link></p>
              </form>
            </div>
          </div>
        </div>
        </div>
        </div>
      </div>
    </div>
  )
}

export default RegisterForm
