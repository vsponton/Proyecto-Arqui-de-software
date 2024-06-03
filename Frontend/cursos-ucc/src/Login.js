import React, { useState } from 'react'
import './Login.css';
import { FaUser,FaLock,FaEnvelope } from "react-icons/fa";

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
            <input type = "text" placeholder='Usename' requiered />
            <FaUser className='icon' />
          </div>
          <div className = "input-box">
            <input type = "email" placeholder='Email' requiered />
            <FaEnvelope className='icon' />
          </div>
          <div className = "input-box">
            <input type = "password"
            placeholder ='Password' requiered />
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





