package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()

	//log.Println("Soy un log")
	// loggear "Hola soy un log" usando la biblioteca log
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")

	// validar que la config este cargada correctamente
	if globals.ClientConfig == nil { // nil indica que no apunta ninguna dirección de memoria
		log.Fatalf("No se pudo cargar la configuración")
	}
	log.Printf("Config cargada: %+v", globals.ClientConfig)
	// loggeamos el valor de la config
	//log.Println(globals.ClientConfig.Mensaje)

	// leer de la consola
	//reader := bufio.NewReader(os.Stdin)
	//text, _ := reader.ReadString('\n')
	//log.Print(text)

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config
	utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje)

	// leer de la consola el mensaje
	//utils.LeerConsola()

	// generamos un paquete y lo enviamos al servidor
	utils.GenerarYEnviarPaquete()
}
