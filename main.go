package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

func main() {
	// Definir el flag para el path del archivo TOML
	tomlPath := flag.String("input", "", "Ruta del archivo TOML de entrada")
	envPath := flag.String("output", ".env", "Ruta del archivo .env de salida (opcional, por defecto es .env)")
	flag.Parse()

	// Validar que se haya proporcionado el path del archivo TOML
	if *tomlPath == "" {
		fmt.Println("Error: Debes proporcionar la ruta del archivo TOML.")
		fmt.Println("Uso: toml_to_env -input <ruta_del_archivo_toml> [-output <ruta_del_archivo_env>]")
		os.Exit(1)
	}

	// Leer el archivo TOML
	tomlData, err := os.ReadFile(*tomlPath)
	if err != nil {
		fmt.Println("Error al leer el archivo TOML:", err)
		os.Exit(1)
	}

	// Parsear el archivo TOML
	var data map[string]interface{}
	err = toml.Unmarshal(tomlData, &data)
	if err != nil {
		fmt.Println("Error al parsear el archivo TOML:", err)
		os.Exit(1)
	}

	// Crear el archivo .env
	envFile, err := os.Create(*envPath)
	if err != nil {
		fmt.Println("Error al crear el archivo .env:", err)
		os.Exit(1)
	}
	defer envFile.Close()

	// Función recursiva para manejar estructuras anidadas
	var writeEnv func(prefix string, data map[string]interface{})
	writeEnv = func(prefix string, data map[string]interface{}) {
		for key, value := range data {
			fullKey := prefix + "_" + key
			switch v := value.(type) {
			case map[string]interface{}:
				// Si el valor es un mapa, llamamos recursivamente a la función
				writeEnv(fullKey, v)
			default:
				// Si no es un mapa, escribimos la clave y el valor en el archivo .env
				envKey := strings.ToUpper(fullKey)
				envValue := fmt.Sprintf("%v", v)
				_, err := envFile.WriteString(fmt.Sprintf("%s=%s\n", envKey, envValue))
				if err != nil {
					fmt.Println("Error al escribir en el archivo .env:", err)
					os.Exit(1)
				}
			}
		}
	}

	// Escribir las claves y valores en el archivo .env
	for section, values := range data {
		switch v := values.(type) {
		case map[string]interface{}:
			writeEnv(strings.ToUpper(section), v)
		default:
			// Si no es un mapa, escribimos directamente
			envKey := strings.ToUpper(section)
			envValue := fmt.Sprintf("%v", v)
			_, err := envFile.WriteString(fmt.Sprintf("%s=%s\n", envKey, envValue))
			if err != nil {
				fmt.Println("Error al escribir en el archivo .env:", err)
				os.Exit(1)
			}
		}
	}

	fmt.Printf("Archivo .env generado exitosamente en: %s\n", *envPath)
}
