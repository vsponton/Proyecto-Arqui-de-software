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





/*import React, { useEffect, useState } from 'react';
import './Login.css';

function Login() {
  const [courses, setCourses] = useState([]);
  const [query, setQuery] = useState(''); // State for current input value
  const [search, setSearch] = useState(''); // State for confirmed search term

  useEffect(() => {
    console.log(`Fetching data from http://localhost:8080/courses/search?query=${search}`)
    // Fetch data from the API based on the search term
    fetch(`http://localhost:8080/courses/search?query=${search}`)
      .then(response => response.json())
      .then(data => setCourses(data.results))
      .catch(error => console.error('Error fetching courses:', error));
  }, [search]);

  const handleSearchChange = (e) => {
    console.log(`Current query: ${e.target.value}`)
    setQuery(e.target.value); // Update the query state as the user types
  };

  const handleSearchSubmit = (e) => {
    e.preventDefault(); // Prevent the default form submission
    setSearch(query); // Update the search state with the current query value
  };

  return (
    <div className="App">
      <div className="SearchBar">
        <form onSubmit={handleSearchSubmit}>
          <input
            type="text"
            placeholder="Search for courses..."
            value={query}
            onChange={handleSearchChange}
          />
        </form>
      </div>
      <div className="Courses">
        {courses != null ? (
          courses.map(course => (
            <div key={course.id} className="Course">
              <img src={course.image_url} alt={course.title} className="Course-image" />
              <div className="Course-details">
                <h1 className="Course-title">{course.title}</h1>
                <p className="Course-description">{course.description}</p>
                <p className="Course-category"><strong>{course.category}</strong></p>
              </div>
            </div>
          ))
        ) : (
          <p>Loading courses...</p>
        )}
      </div>
    </div>
  );
}

export default Login;
*/
