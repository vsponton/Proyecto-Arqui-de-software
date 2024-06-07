import React, { useEffect, useState } from 'react';
import './Home.css';
import Cookies from "universal-cookie";
import { FaEmber } from 'react-icons/fa';

const Cookie = new Cookies();


/**
async function getUserByEmail(email){
    return await fetch('http://localhost:8080/user/' + email, {
    method: 'GET',
    
}).then(response => response.json())
}

async function getCourses(){
  return await fetch('http://localhost:8080/course')
    .then(response => response.json())
}

async function getCursoByUserId(userId){
  return await fetch('http://localhost:8080/course/' + userId, {
    method: "GET",
    
  }).then(response => response.json())
}
async function getCursoByCategory(category){
  return await fetch('http://localhost:8080/course/' + category, {
    method: "GET",
    
  }).then(response => response.json())
}

async function postCurso(curso){
  return await fetch('http://localhost:8080/course', {
    method: "POST",
    
    body: JSON.stringify(curso)
  }).then(response => response.json())
}

async function putCurso(curso){
  return await fetch('http://localhost:8080/course/' + curso.id_course, {
    method: "PUT",
    
    body: JSON.stringify(curso)
  }).then(response => response.json())
}

async function deleteCurso(curso){
  return await fetch('http://localhost:8080/course/' + curso.id_course, {
    method: "DELETE",
    
  }).then(response => response.json())
}

async function getCursoByDescription(description){
  return await fetch('http://localhost:8080/course/' + description, {
    method: "GET",
    
  }).then(response => response.json())
}
**/

const Home = () => {
  const [admin, setAdmin] = useState(false);
  const [isLogged, setIsLogged] = useState(false);

  const [needCourses, setNeedCourses] = useState(true);
  const [needAvailableCourses, setNeedAvailableCourses] = useState(true);
  
  const [courses, setCourses] = useState([
    /*{
      title: "Base de datos",
      description: "El curso ofrece una introducción completa al diseño, implementación y gestión de bases de datos..",
      category: "programacion",
      image_url: "https://blog.continentaluniversity.us/hubfs/que-es-sistema-gestion-base-datos-cuf.jpg",
      duration: "5 meses, 2 veces por semana 2 horas cada dia.",
      instructor: "Carlos Ceballos, Ingeniero en Sistemas ",
      requirements: "Acceso a una computadora con conexión a internet.",
    },
    {
      title: "Programación en C++",
      description: " Este curso ofrece una introducción completa a la programación en C++, con conceptos basicos y avanzados llevados a la practica.",
      category: "programacion",
      image_url: "https://img-c.udemycdn.com/course/750x422/5127236_8148.jpg",
      duration: "8 semanas, con un compromiso de 4-6 horas por semana.",
      instructor: "Florencia Ceballos, Ingeniera en Sistemas con más de 10 años de experiencia en desarrollo de software y enseñanza de programación.",
      requirements: "Acceso a una computadora con conexión a internet.",
    },
    {
      title: "Ciberseguridad",
      description: " Los estudiantes aprenderán sobre amenazas y vulnerabilidades, técnicas de protección de datos, criptografía y estrategias de defensa cibernética.",
      category: "programacion",
      image_url: "https://web-assets.esetstatic.com/tn/-x700/wls/2023/2023-09/cybersecurity.jpeg",
      duration: "10 semanas, con un compromiso de 4-6 horas por semana.",
      instructor: "Agostina Cristal, Ingeniera en Sistemas con más de 10 años de experiencia en desarrollo de software y enseñanza de programación.",
      requirements: "Conocimientos básicos de informática y redes y acceso a una computadora con conexión a internet.",
    }*/
  ]);

  useEffect(() => {
    fetch('http://localhost:8080/course')
      .then(response => response.json())
      .then(data => setCourses(data))
      .catch(error => console.error('Error fetching courses:', error));
  }, [courses]);

  const [availableCourses, setAvailableCourses] = useState([
    {
      title: "Programación en GO",
      description: "Este curso está diseñado para proporcionar una comprensión completa del lenguaje de programación GO. A través de una combinación de teoría y práctica, los estudiantes aprenderán los fundamentos de la programación en GO, incluyendo estructuras de control, funciones, clases, objetos, y manejo de memoria.",
      category: "programacion",
      image_url: "https://i.pinimg.com/564x/3d/d4/fd/3dd4fdcd69a2858b06bd01be9ea3c531.jpg",
      duration: "8 semanas, con un compromiso de 4-6 horas por semana.",
      instructor: "Flor Ceballos, Ingeniera en Sistemas con más de 10 años de experiencia en desarrollo de software y enseñanza de programación.",
      requirements: "Acceso a una computadora con conexión a internet.",
    }
  ]);

  if (Cookie.get("token") && !isLogged){
    setIsLogged(true)
  }

  /*
  if(!courses.length && needCourses){
    getCourses().then(response => setCourses(response))
    setNeedCourses(false)
  }
    */

  const toggleAdmin = () => {
    setAdmin(!admin);
  };

  const showHomeAdmin = () => {
    return (
      <div className="container">
        <div className="sidebar">
          <div className="admin">ADMINISTRADOR</div>
          <div className="menu-item">Cursos</div>
        </div>
        <div className="main-content">
          <div className="search-bar">
            <input type="text" placeholder="Buscar" />
          </div>
          <div className="courses">
            {courses.map((course, index) => (
              <div key={index} className="Course">
                <div className="course-item">
                  <div>
                    <img src={course.image_url} alt={course.title} className="Course-image" />
                    <span>{course.title}</span>
                  </div>
                  <div>
                    <p className="course-category">{course.category}</p>
                    <p className="course-duration">{course.duration}</p>
                    <p className="course-instructor">{course.instructor}</p>
                    <p className="course-requirements">{course.requirements}</p>
                    <div className="actions">
                      <button className="edit">✏️</button>
                      <button className="add">+</button>
                    </div>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>
        <div className="add-delete-buttons">
          <button className="add-course">+</button>
          <div className="add-course-text">
            <p>+ add new course</p>
          </div>
          <button className="delete-course">✖</button>
          <div className="delete-course-text">
            <p>x delete course</p>
          </div>
        </div>
      </div>
    );
  };

  const showHomeAlumno = () => {
    return (
      <div className="containerAlum">
        <div className="left-section">
          <div className="header">
            <div className="student">ALUMNO</div>
            <div className="search-bar">
              <input type="text" placeholder="Buscar" />
            </div>
          </div>
          <div className="courses">
            <div className="courses-title">Mis Cursos</div>
            {courses.map((course, index) => (
              <div key={index} className="Course">
                <div className="course-item">
                  <div>
                    <img src={course.image_url} alt={course.title} className="Course-image" />
                    <span>{course.title}</span>
                  </div>
                  <div>
                    <p className="course-category">Categoria: {course.category}</p>
                    <p className="course-duration">{course.duration}</p>
                    <p className="course-instructor">{course.instructor}</p>
                    <p className="course-requirements">{course.requirements}</p>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>
        <div className="right-section">
          <div className="available-courses">
            <div className="available-courses-title">Cursos disponibles</div>
            {availableCourses.map((course, index) => (
              <div key={index} className="Course">
                <div className="course-item">
                  <div>
                    <img src={course.image_url} alt={course.title} className="Course-image" />
                    <span>{course.title}</span>
                  </div>
                  <div>
                    <p className="course-duration">{course.duration}</p>
                  </div>
                </div>
              </div>
            ))}
          </div>
          <button className="alum-button">INSCRIBIRME</button>
        </div>
      </div>
    );
  };

  return (
    <div>
      <button onClick={toggleAdmin}>CAMBIAR VISTA</button>
      {admin ? showHomeAdmin() : showHomeAlumno()}
    </div>
  );
};

export default Home;
