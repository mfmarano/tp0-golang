package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()

	globals.ClientConfig = utils.IniciarConfiguracion("config.json")

	if globals.ClientConfig == nil {
		log.Fatalf("No se pudo cargar la configuración")
	}

	utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje)

	utils.GenerarYEnviarPaquete()
}
