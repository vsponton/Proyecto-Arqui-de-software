import React, { useState } from 'react'
import './Home.css';
import Cookies from "universal-cookie";

const Cookie = new Cookies();
function goto(path){
    window.location = window.location.origin + path
  }
const Home = () => {
  
    const[admin, setAdmin] =useState(false);
    const[courses, setCourses] = useState([
        {
            "title":"Programación en C++",
            "description": "Este curso está diseñado para proporcionar una comprensión completa del lenguaje de programación C++. A través de una combinación de teoría y práctica, los estudiantes aprenderán los fundamentos de la programación en C++, incluyendo estructuras de control, funciones, clases, objetos, y manejo de memoria.",
            "category": "programacion",
            "image_url":"https://i.pinimg.com/564x/3d/d4/fd/3dd4fdcd69a2858b06bd01be9ea3c531.jpg",
            "duration": "8 semanas, con un compromiso de 4-6 horas por semana.",
            "instructor": "Agustin Ceballos, Ingeniero en Sistemas con más de 10 años de experiencia en desarrollo de software y enseñanza de programación.",
            "requirements": "•	Acceso a una computadora con conexión a internet.",
    
        },{
            "title":"Programación en GO",
            "description": "Este curso está diseñado para proporcionar una comprensión completa del lenguaje de programación C++. A través de una combinación de teoría y práctica, los estudiantes aprenderán los fundamentos de la programación en C++, incluyendo estructuras de control, funciones, clases, objetos, y manejo de memoria.",
            "category": "programacion",
            "image_url":"https://i.pinimg.com/564x/3d/d4/fd/3dd4fdcd69a2858b06bd01be9ea3c531.jpg",
            "duration": "8 semanas, con un compromiso de 4-6 horas por semana.",
            "instructor": "Flor Ceballos, Ingeniero en Sistemas con más de 10 años de experiencia en desarrollo de software y enseñanza de programación.",
            "requirements": "•	Acceso a una computadora con conexión a internet.",
    
        }
    ])
    const[availableCourses, setAvailableCourses] = useState([
        {
            "title":"Programación en GO",
            "description": "Este curso está diseñado para proporcionar una comprensión completa del lenguaje de programación C++. A través de una combinación de teoría y práctica, los estudiantes aprenderán los fundamentos de la programación en C++, incluyendo estructuras de control, funciones, clases, objetos, y manejo de memoria.",
            "category": "programacion",
            "image_url":"https://i.pinimg.com/564x/3d/d4/fd/3dd4fdcd69a2858b06bd01be9ea3c531.jpg",
            "duration": "8 semanas, con un compromiso de 4-6 horas por semana.",
            "instructor": "Flor Ceballos, Ingeniero en Sistemas con más de 10 años de experiencia en desarrollo de software y enseñanza de programación.",
            "requirements": "•	Acceso a una computadora con conexión a internet.",
    
        }
    ])
    const toggleAdmin = () => {
        let prev = Cookie.get("admin")
        prev == "true" ? Cookie.set("admin", "false") : Cookie.set("admin", "true");
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
                    <input type="text" placeholder="Buscar" />
                </div>
                <div class="courses">


                    {courses != null ? (
                            courses.map(course => (
                                    <div key={course.id} className='Course'>
                                        <div class="course-item">
                                            <div>
                                                <img src={course.image_url} alt ={course.title} className='Course-image'/> 
                                                <span>Programación en C++</span>
                                            </div>
                                            <div>
                                                <p className='course-category'>{course.category}</p>
                                                <p className='course-duration'>{course.duration}</p>
                                                <p className='course-instructor'>{course.instructor}</p>
                                                <p className='course-requirements'>{course.requirements}</p>
                                                <div class="actions">
                                                    <button class="edit">✏️</button>
                                                    <button class="add">+</button>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                )
                            )
                        ):(
                        <p> Loading Courses...</p>
                    )}
                    
                    
                </div>
            </div>
            <div class="add-delete-buttons">
                <button class="add-course">+</button>
                <div class="add-course-text">
                    <p>+ add new course</p>
                </div>
                <button class="delete-course">✖</button>
                <div class="delete-course-text">
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
                    <div class="student">ALUMNO</div>
                    <div class="search-bar">
                    <input type="text" placeholder="Buscar" />
                </div>
                </div>
                <div className="courses">
                    <div className="courses-title">Mis Cursos</div>
                    

                    {courses != null ? (
                            courses.map(course => (
                                    <div key={course.id} className='Course'>
                                        <div class="course-item">
                                            <div>
                                                <img src={course.image_url} alt ={course.title} className='Course-image'/> 
                                                <span>Programación en C++</span>
                                            </div>
                                            <div>
                                                <p className='course-category'>Categoria: {course.category}</p>
                                                <p className='course-duration'>{course.duration}</p>
                                                <p className='course-instructor'>{course.instructor}</p>
                                                <p className='course-requirements'>{course.requirements}</p>
                                            </div>
                                        </div>
                                    </div>
                                )
                            )
                        ):(
                        <p> Loading Courses...</p>
                    )}
                </div>
            </div>
            <div className="right-section">
                <div className="available-courses">
                    <div className="available-courses-title">Cursos disponibles</div>
                    {availableCourses != null ? (
                            availableCourses.map(course => (
                                    <div key={course.id} className='Course'>
                                        <div class="course-item">
                                            <div>
                                                <img src={course.image_url} alt ={course.title} className='Course-image'/> 
                                                <span>Programación en C++</span>
                                            </div>
                                            <div>
                                                <p className='course-duration'>{course.duration}</p>
                                            </div>
                                        </div>
                                    </div>
                                )
                            )
                        ):(
                        <p> Loading Courses...</p>
                    )}
                </div>
                <button className="alum-button">INSCRIBIRME</button>
            </div>
        </div>
                
    )
  }

  return(
    <div>
    <button onClick={toggleAdmin}>CAMBIAR VISTA</button>
      {
        Cookie.get("token") !== undefined ? (
        Cookie.get("admin") == "true" ? showHomeAdmin() : showHomeAlumno()) : goto("/login")
      }
    </div>
  );
};
export default Home;
