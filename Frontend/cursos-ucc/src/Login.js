import React, { useState } from 'react'
import './Login.css';
import { FaUser,FaLock,FaEnvelope } from 'react-icons/fa';
import Cookies from "universal-cookie";

const Cookie = new Cookies();

async function login(username, password) {
  return await fetch('http://localhost:8090/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
	body: JSON.stringify({"username": username, "password":password})
  })
    .then(response => {
      if (response.status == 400 || response.status == 401)
      {
        return {"user_id": -1}
      }
      return response.json()
    })
    .then(response => {
      Cookie.set("user_id", response.user_id, {path: '/'})
      Cookie.set("username", username, {path: '/login'})
    })
 }
const Login = () => {
  const[register, setRegister] =useState(false);

  const registerLink = () => {
    setRegister(true);
  };
  const loginLink = () => {
      setRegister(false);
  };

  const showRegister = () => {
    return (
      <div className = "form-box register">
        <form action = "">
          <h1>Registration</h1>
          <div className = "input-box">
            <input type = "text" placeholder='Firstname' required />
            <FaUser className='icon' />
          </div>
          <div className = "input-box">
            <input type = "text" placeholder='Lastname' required />
            <FaUser className='icon' />
          </div>
          <div className = "input-box">
            <input type = "email" placeholder='Email' required />
            <FaEnvelope className='icon' />
          </div>
          <div className = "input-box">
            <input type = "password"
            placeholder ='Password' required />
            <FaLock className='icon'/>
          </div>
          <button type = "submit">Register</button>
          <div className="register-link">
            <p>Already have an account? <a href="#" onClick={loginLink}>Login</a></p>
          </div>
        </form>
      </div>
    )
  }

    const showLogin = () => {
      return (
        <div className = "form-box login">
        <form action = "">
          <h1>Login</h1>
          <div className = "input-box">
            <input type = "email" placeholder='Email' requiered />
            <FaEnvelope className='icon' />
          </div>
          <div className = "input-box">
            <input type = "password"
            placeholder ='Password' requiered />
            <FaLock className='icon'/>
          </div>
          <button type = "submit">Login</button>
          <div className="register-link">
            <p>Don't have an account? <a href="#" onClick={registerLink}>Register</a></p>
          </div>
        </form>
      </div>
      )
    }
  

  return(
    <div>
      {register ? showRegister() : showLogin()}
    </div>
  );
};
export default Login;





