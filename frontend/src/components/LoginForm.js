import React, { useState } from "react";
import Logo from "../assets/images/Logo Text.png";

function LoginForm({ LoginForm, error }) {
  const [details, setDetails] = useState({ name: "", email: "", password: "" });

  const submitHandler = (e) => {
    e.preventDefault();
    LoginForm(details);
  };
  return (
    <form onSubmit={submitHandler}>
      <div className="header">
        <img src={Logo} alt=""></img>
      </div>
      <div className="form-inner">
        <h1>LOG IN</h1>
        {error !== "" && <p className="error">{error}</p>}
        <div className="form-group">
          <label htmlFor="name">NAME : </label>
          <input type="text" name="name" id="name" onChange={(e) => setDetails({ ...details, name: e.target.value })} value={details.name} />
        </div>
        <div className="form-group">
          <label htmlFor="email">EMAIL : </label>
          <input type="email" name="email" id="email" onChange={(e) => setDetails({ ...details, email: e.target.value })} value={details.email} />
        </div>
        <div className="form-group">
          <label htmlFor="password">PASSWORD : </label>
          <input type="password" name="password" id="password" onChange={(e) => setDetails({ ...details, password: e.target.value })} value={details.password} />
        </div>
        <input type="submit" value="LOG IN" />
      </div>
    </form>
  );
}

export default LoginForm;
