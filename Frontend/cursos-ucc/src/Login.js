import React, { useState } from 'react'
import './Login.css';
import { FaUser, FaLock, FaEnvelope } from 'react-icons/fa';
import Cookies from "universal-cookie";

const Cookie = new Cookies();

function goto(path){
  window.location = window.location.origin + path
}

async function login(username, password) {
  return await fetch('http://localhost:8080/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
	body: JSON.stringify({"email": username, "password": password})
  })
    .then(response => {
      if (response.status === 400 || response.status === 401) {
        return { "user_id": -1 }
      }
      return response.json()
    })
    .then(response => {
      Cookie.set("token", response.token, { path: '/' })
      Cookie.set("email", username)
      goto("/")
    })
    
}

async function postRegister(firstname, lastname, email, password) {
  return await fetch('http://localhost:8080/register', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
	body: JSON.stringify({"firstname": firstname, "lastname": lastname, "email": email, "password": password})
  })
    .then(response => {
      if (response.status === 400 || response.status === 401) {
        return { "user_id": -1 }
      }
      return response.json()
    })
    .then(response => {
      Cookie.set("token", response.token, { path: '/' })
    })
}

const Login = () => {
  const [register, setRegister] = useState(false);
  const [errorMessages, setErrorMessages] = useState({});
  const error = "Invalid email or password";

  const registerLink = () => {
    setRegister(true);
  };

  const loginLink = () => {
    setRegister(false);
  };

  const handleLoginSubmit = async (event) => {
    event.preventDefault();
    const { email, password } = event.target.elements;
    const userData = await login(email.value, password.value);
    if (Cookie.get("token") > -1) {
      window.location = window.location.origin + "/";
    } else {
      setErrorMessages({ name: "default", message: error });
    }
  };

  const handleRegisterSubmit = async (event) => {
    event.preventDefault();
    const { firstname, lastname, email, password } = event.target.elements;
    const userData = await postRegister(firstname.value, lastname.value, email.value, password.value);
    if (Cookie.get("user_id") > -1) {
      window.location = window.location.origin + "/";
    } else {
      setErrorMessages({ name: "default", message: error });
    }
  };

  const renderErrorMessage = (name) =>
    name === errorMessages.name && (
      <div className="error">{errorMessages.message}</div>
    );

  const showRegister = () => {
    return (
      <div className="form-box register">
        <form onSubmit={handleRegisterSubmit}>
          <h1>Registration</h1>
          <div className="input-box">
            <input type="text" name="firstname" placeholder="Firstname" required />
            <FaUser className='icon' />
          </div>
          <div className="input-box">
            <input type="text" name="lastname" placeholder="Lastname" required />
            <FaUser className='icon' />
          </div>
          <div className="input-box">
            <input type="email" name="email" placeholder="Email" required />
            <FaEnvelope className='icon' />
          </div>
          <div className="input-box">
            <input type="password" name="password" placeholder="Password" required />
            <FaLock className='icon' />
          </div>
          {renderErrorMessage("default")}
          <button type="submit">Register</button>
          <div className="register-link">
            <p>Already have an account? <a href="#" onClick={loginLink}>Login</a></p>
          </div>
        </form>
      </div>
    )
  }

  const showLogin = () => {
    return (
      <div className="form-box login">
        <form onSubmit={handleLoginSubmit}>
          <h1>Login</h1>
          <div className="input-box">
            <input type="email" name="email" placeholder="Email" required />
            <FaEnvelope className='icon' />
          </div>
          <div className="input-box">
            <input type="password" name="password" placeholder="Password" required />
            <FaLock className='icon' />
          </div>
          {renderErrorMessage("default")}
          <button type="submit">Login</button>
          <div className="register-link">
            <p>Don't have an account? <a href="#" onClick={registerLink}>Register</a></p>
          </div>
        </form>
      </div>
    )
  }

  return (
    <div>
      {register ? showRegister() : showLogin()}
    </div>
  );
};

export default Login;
