import React, { useEffect, useState} from 'react'
import './Courses.css';


function getCursoById(id){

    if(id === 1){
        return {
            "title":"Programación en C++",
            "description": "Este curso está diseñado para proporcionar una comprensión completa del lenguaje de programación C++. A través de una combinación de teoría y práctica, los estudiantes aprenderán los fundamentos de la programación en C++, incluyendo estructuras de control, funciones, clases, objetos, y manejo de memoria.",
            "category": "programacion",
            "image_url":"https://i.pinimg.com/564x/3d/d4/fd/3dd4fdcd69a2858b06bd01be9ea3c531.jpg",
            "duration": "8 semanas, con un compromiso de 4-6 horas por semana.",
            "instructor": "Agustin Ceballos, Ingeniero en Sistemas con más de 10 años de experiencia en desarrollo de software y enseñanza de programación.",
            "requirements": "•	Acceso a una computadora con conexión a internet.",

        }
    }
    return {
        "title":"Programación en GO",
        "description": "Este curso está diseñado para proporcionar una comprensión completa del lenguaje de programación C++. A través de una combinación de teoría y práctica, los estudiantes aprenderán los fundamentos de la programación en C++, incluyendo estructuras de control, funciones, clases, objetos, y manejo de memoria.",
        "category": "programacion",
        "image_url":"https://i.pinimg.com/564x/3d/d4/fd/3dd4fdcd69a2858b06bd01be9ea3c531.jpg",
        "duration": "8 semanas, con un compromiso de 4-6 horas por semana.",
        "instructor": "Flor Ceballos, Ingeniero en Sistemas con más de 10 años de experiencia en desarrollo de software y enseñanza de programación.",
        "requirements": "•	Acceso a una computadora con conexión a internet.",

    }

}

const Courses = () => {
    const[course, setCourse] =useState({});

    const loadCourse = () =>{
        let curso = getCursoById(5);
        setCourse(curso);
    }

    const showCourses= () =>{
        return(
            <div>
                <div class="course-title">
                    {course.title}
                </div>
                <div class="course-info">
                    <div>
                        <h4>Descripcion:</h4>
                        <p class="course-description">{course.description}</p>
                    </div>
                    <div>
                        
                        <p class="course-instructor">{course.instructor}</p>
                    </div>
                    <div>
                        
                        <p class="course-duration">{course.duration}</p>
                    </div>
                    <div>
                        <div class="course-requirements-container">
                            <h4>Requisitos</h4>
                            <p class="course-requirements">{course.requirements}</p>
                        </div>
                        <div class="course-image">
                            <img src={course.image_url} alt={course.title}></img>
                        </div>
                    </div>
                </div>
            </div>
        )
    }

    return (
        <div>
            <button onClick={loadCourse}>CARGAR</button>
            {showCourses()}
        </div>
    )
}

export default Courses;