package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Usuarios struct {
	Usuarios []Usuario `json:"usuarios"`
}

type Usuario struct {
	Nome   string `json:"Nome"`
	Tipo   string `json:"Tipo"`
	Idade  int    `json:"Idade"`
	Social Social `json:"Social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {

	jsonFile, err := os.Open("json/usuarios.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened usuarios.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var usuarios Usuarios
	err = json.Unmarshal(byteValue, &usuarios)
	if err != nil {
		return
	}

	for i := 0; i < len(usuarios.Usuarios); i++ {
		fmt.Println("Nome : " + usuarios.Usuarios[i].Nome)
		fmt.Println("Tipo : " + usuarios.Usuarios[i].Tipo)
		fmt.Println("Idade : " + strconv.Itoa(usuarios.Usuarios[i].Idade))
		fmt.Println("Facebook : " + usuarios.Usuarios[i].Social.Facebook)
		fmt.Println("Twitter : " + usuarios.Usuarios[i].Social.Twitter)
	}

}
