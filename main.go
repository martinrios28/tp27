package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"
)

func validar(user string, password string) bool {
	file, err := os.Open("usuarios.csv")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return false
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo CSV:", err)
		return false
	}

	for _, record := range records {
		if len(record) >= 2 && record[0] == user && record[1] == password {
			return true
		}
	}

	return false
}

func main() {
	var user, password string
	fmt.Print("Ingrese el usuario: ")
	fmt.Scanln(&user)
	fmt.Print("Ingrese la clave: ")
	fmt.Scanln(&password)

	if validar(user, password) {
		fmt.Println("Acceso concedido.")
	} else {
		fmt.Println("Acceso denegado.")
	}
}

func TestValidar(t *testing.T) {
	// Caso de prueba válido
	if !validar("bob", "50135875") {
		t.Error("Se esperaba que el usuario y la clave sean válidos")
	}

	// Caso de prueba inválido
	if validar("alice", "12345678") {
		t.Error("Se esperaba que el usuario y la clave sean inválidos")
	}

	// Caso de prueba con espacios en blanco
	if validar("eve ", "25436928") {
		t.Error("Se esperaba que el usuario y la clave sean inválidos debido a espacios en blanco")
	}
}
