package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Leyendo con método 1")
	leoArchivo()

	fmt.Println("Leyendo con método 2")
	leoArchivo2()

	fmt.Println("Grabando archivo con método 1")
	graboArchivo()

	fmt.Println("Grabando archivo con método 2")
	graboArchivo2()
}

func leoArchivo() {
	archivo, err := ioutil.ReadFile("./archivo.txt")
	if err != nil {
		fmt.Println("Error reading")
	} else {
		fmt.Println(string(archivo))
	}
}

func leoArchivo2() {
	archivo, err := os.Open("./archivo.txt")
	if err != nil {
		fmt.Println("Error reading")
	} else {
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {
			registro := scanner.Text()
			fmt.Printf("Linea > " + registro + "\n")
		}
	}
	archivo.Close()
}

func graboArchivo() {
	archivo, err := os.Create("./nuevoArchivo.txt")
	if err != nil {
		fmt.Println("Error reading")
		return
	}
	fmt.Fprintln(archivo, "Esta es una linea nueva a grabar en un archivo que acabo de crear")
	archivo.Close()
}

func graboArchivo2() {
	fileName := "./nuevoArchivo.txt"
	if Append(fileName, "\n Esta es una segunda linea agregada al archivo ya creado") == false {
		fmt.Println("Error writing the second line")
	}
}

func Append(fileName string, text string) bool {
	archivo, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644) //las variables os.O_WRONLY es para dar permisos de lectura y escritura, mientras que os.O_APPEND es para inficar el modo de apertura y posterior modificación (agregado de información nueva hacia el final del archivo). En cuanto a 0644 es parte del manejo de permisos en Linux
	if err != nil {
		fmt.Println("Error opening file")
		return false
	}

	_, err = archivo.WriteString(text)
	if err != nil {
		fmt.Println("Error writing file")
		return false
	}

	return true
}
