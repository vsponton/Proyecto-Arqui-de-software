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
<<<<<<< HEAD

async function getAvailableCourses(){
  return await fetch('http://localhost:8080/course/user/available', {
    method: "POST",
    body: JSON.stringify({"token":Cookie.get("token")})
  }).then(response => response.json())
}

async function getRegisteredCourses(){
  return await fetch('http://localhost:8080/course/user/registered', {
    method: "POST",
    body: JSON.stringify({"token":Cookie.get("token")})
  }).then(response => response.json())
}

async function registerToCourse(id){
  return await fetch("http://localhost:8080/course/register", {
    method: "POST",
    body: JSON.stringify({
      token: Cookie.get("token"),
      course_id: id,
    })
    
  }).then(response => response.json())
}

function goto(path){
  window.location = window.location.origin + path
}

=======
**/
>>>>>>> b8ca57bfe323424126dfecabcf2b9fa4fc94d9f5

const Home = () => {
  const [admin, setAdmin] = useState(false);
  const [isLogged, setIsLogged] = useState(false);

  const [needCourses, setNeedCourses] = useState(true);
  const [needAvailableCourses, setNeedAvailableCourses] = useState(true);
  const [needRegisteredCourses, setNeedRegisteredCourses] = useState(true);
  
  const [courses, setCourses] = useState([]);
  const [registeredCourses, setRegisteredCourses] = useState([]);
  const [availableCourses, setAvailableCourses] = useState([]);

<<<<<<< HEAD
=======
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
>>>>>>> b8ca57bfe323424126dfecabcf2b9fa4fc94d9f5


  /*
  if(!courses.length && needCourses){
    getCourses().then(response => setCourses(response))
    setNeedCourses(false)
  }
<<<<<<< HEAD
  if(!availableCourses.length && needAvailableCourses){
    getAvailableCourses().then(response => {
      if (response) {
        setAvailableCourses(response)
      }
    })
    setNeedAvailableCourses(false)
  }
  if(!registeredCourses.length && needRegisteredCourses){
    getRegisteredCourses().then(response => {
      if (response) {
        setRegisteredCourses(response)
      }
    })
    setNeedRegisteredCourses(false)
  }
=======
    */

>>>>>>> b8ca57bfe323424126dfecabcf2b9fa4fc94d9f5
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
            {courses ? courses.map((course, index) => (
              <div key={index} className="Course" onClick={() => goto("/courses/" + course.id_course)}>
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
            )) : <p> Loading... </p>}
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
            {registeredCourses ? registeredCourses.map((course, index) => (
              <div key={index} className="Course" onClick={() => goto("/courses/" + course.id_course)}>
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
            )) : <p>Loading courses...</p>}
          </div>
        </div>
        <div className="right-section">
          <div className="available-courses">
            <div className="available-courses-title">Cursos disponibles</div>
            {availableCourses ? availableCourses.map((course, index) => (
              <div key={index} className="Course" onClick={() => goto("/courses/" + course.id_course)}>
                <div className="course-item">
                  <div>
                    <img src={course.image_url} alt={course.title} className="Course-image" />
                    <span>{course.title}</span>
                  </div>
                  <div>
                    <p className="course-duration">{course.duration}</p>
                  </div>
                  <button className="alum-button" onClick={() => registerToCourse(course.id_course)}>INSCRIBIRME</button>
                </div>
              </div>
            )) : <p>Loading Courses...</p>}
          </div>
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
