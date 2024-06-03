import React, { useState } from 'react'
import './Home.css';
import { FaUser,FaLock,FaPencil } from "react-icons/fa";

const Home = () => {
  
    const[admin, setAdmin] =useState(false);

    const toggleAdmin = () => {
        setAdmin(!admin);
    }
    
    const showHomeAdmin = () => {
        return (
            //if ... es administrador
            <div class="container">
            <div class="sidebar">
                <div class="admin">ADMINISTRADOR</div>
                <div class="menu-item">Cursos</div>
            </div>
            <div class="main-content">
                <div class="search-bar">
                    <input type="text" placeholder="Search" />
                </div>
                <div class="courses">
                    <div class="course-item">
                        <span>Programación en C++</span>
                        <div class="actions">
                            <button class="edit">✏️</button>
                            <button class="add">+</button>
                        </div>
                    </div>
                    <div class="course-item">
                        <span>Desarrollo de Software</span>
                        <div class="actions">
                            <button class="edit">✏️</button>
                            <button class="add">+</button>
                        </div>
                    </div>
                    <div class="course-item">
                        <span>Base de datos</span>
                        <div class="actions">
                            <button class="edit">✏️</button>
                            <button class="add">+</button>
                        </div>
                    </div>
                    <div class="course-item">
                        <span>Fundamentos de programación</span>
                        <div class="actions">
                            <button class="edit">✏️</button>
                            <button class="add">+</button>
                        </div>
                    </div>
                    <div class="course-item">
                        <span>Ciberseguridad</span>
                        <div class="actions">
                            <button class="edit">✏️</button>
                            <button class="add">+</button>
                        </div>
                    </div>
                    <div class="course-item">
                        <span>Desarrollo Backend</span>
                        <div class="actions">
                            <button class="edit">✏️</button>
                            <button class="add">+</button>
                        </div>
                    </div>
                    <div class="course-item">
                        <span>Desarrollo Frontend</span>
                        <div class="actions">
                            <button class="edit">✏️</button>
                            <button class="add">+</button>
                        </div>
                    </div>
                </div>
            </div>
            <div class="add-delete-buttons">
                <button class="add-course">+</button>
                <div class="add-course-text">
                    <p>+ add new course</p>
                </div>
                <button class="delete-course">✖</button>
                <div class="description">
                    <p>x delete course</p>
                </div>
            </div>
        </div>
        //else ... alumno
        
            
           
       
        )
    }

  const showHomeAlumno = () => {
    return (
        <div className="containerAlum">
            <div className="left-section">
                <div className="header">
                    <input type="text" placeholder="ALUMNO" />
                    <div class="search-bar">
                    <input type="text" placeholder="Search" />
                </div>
                </div>
                <div className="courses">
                    <div className="courses-title">Cursos</div>
                    <div className="course">
                        <span>programación en C++</span>
                        <button>+</button>
                    </div>
                    <div className="course">
                        <span>Desarrollo de Software</span>
                        <button>+</button>
                    </div>
                    <div className="course">
                        <span>Base de datos</span>
                        <button>+</button>
                    </div>
                    <div className="course">
                        <span>Fundamentos de programación</span>
                        <button>+</button>
                    </div>
                    <div className="course">
                        <span>Ciberseguridad</span>
                        <button>+</button>
                    </div>
                </div>
            </div>
            <div className="right-section">
                <div className="available-courses">
                    <div className="available-courses-title">Cursos disponibles</div>
                    <div className="available-course">
                        <input type="text" value="machine-learning" readOnly />
                    </div>
                    <div className="available-course">
                        <input type="text" value="desarrollo web" readOnly />
                    </div>
                    <div className="available-course">
                        <input type="text" value="JavaScript desde cero" readOnly />
                    </div>
                </div>
                <button className="alum-button">INSCRIBIRME</button>
            </div>
        </div>
                
    )
  }

  return(
    <div>
    <button onClick={toggleAdmin}>CAMBIAR VISTA</button>
      {admin ? showHomeAdmin() : showHomeAlumno()}
    </div>
  );
};
export default Home;
