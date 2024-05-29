package app

import (
	log "github.com/sirupsen/logrus"
	usersController "mvc-go/controllers/users"
)

func mapUrls() {

	// Sensores Mapping
	router.POST("/sensor", sensorController.SensorInsert)
	router.PUT("/sensor/:id/activar", sensorController.ActivarSensor)
	router.PUT("/sensor/:id/pausar", sensorController.PausarSensor)
	router.POST("/login", usersController.Login)

	router.GET("/vendedor/:id", vendedorController.GetVendedorById)

	router.GET("/comprador/:id", compradorController.GetCompradorById)

	// Bario Mapping
	router.GET("/barrio/:id", barrioController.GetBarrioDetalleById)
	router.GET("/barrio", barrioController.GetBarrios)

	// Medidcion Mapping
	router.POST("/medicion", medicionController.RegistrarMedicion)

	log.Info("Finishing mappings configurations")
}
