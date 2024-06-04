// backend/logging/logging.go
package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	// Puedes configurar el logger aquí
	Log.SetFormatter(&logrus.JSONFormatter{}) // cambiar el formato a JSON si lo preferimos
	Log.SetOutput(os.Stdout)                  // Enviar los logs a la salida estándar
	Log.SetLevel(logrus.DebugLevel)           // Ajusta el nivel de log a Debug
}
