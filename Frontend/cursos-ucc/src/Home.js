import React, { useState } from 'react';
import './Home.css';

const Home = () => {
  const [admin, setAdmin] = useState(false);
  const [courses, setCourses] = useState([
    {
      title: "Programación en C++",
      description: "Este curso está diseñado para proporcionar una comprensión completa del lenguaje de programación C++. A través de una combinación de teoría y práctica, los estudiantes aprenderán los fundamentos de la programación en C++, incluyendo estructuras de control, funciones, clases, objetos, y manejo de memoria.",
      category: "programacion",
      image_url: "https://i.pinimg.com/564x/3d/d4/fd/3dd4fdcd69a2858b06bd01be9ea3c531.jpg",
      duration: "8 semanas, con un compromiso de 4-6 horas por semana.",
      instructor: "Agustin Ceballos, Ingeniero en Sistemas con más de 10 años de experiencia en desarrollo de software y enseñanza de programación.",
      requirements: "Acceso a una computadora con conexión a internet.",
    },
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
