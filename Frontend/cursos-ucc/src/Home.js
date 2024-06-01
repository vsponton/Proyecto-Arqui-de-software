import React, { useState } from 'react'
import './Home.css';
import { FaUser,FaLock,FaPencil } from "react-icons/fa";

const Home = () => {
  
    const[register, setRegister] =useState(false);


    const showHome = () => {
        return (
            <div class="container">
            <div class="sidebar">
                <div class="admin">ADMINISTRADOR</div>
                <div class="menu-item">Cursos</div>
            </div>
            <div class="main-content">
                <div class="search-bar">
                    <input type="text" placeholder="Buscar" />
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
            
           
       
        )
    }

  

  return(
    <div>
      {showHome()}
    </div>
  );
};
export default Home;

