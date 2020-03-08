package main

import (
	"fmt"
	"github.com/rafaelarodrigues90/go-apirest/config"
	"github.com/rafaelarodrigues90/go-apirest/src/modules/profile/model"
	"github.com/rafaelarodrigues90/go-apirest/src/modules/profile/repository"
)

func main() {
	// fmt.Println("Go MongoDB")

	db, err := config.GetMongoDB()

	if err != nil {
		fmt.Println(err)
	}

	// create database
	userRepository := repository.NewUserRepositoryMongo(db, "usuarios")

	
	saveUser(userRepository)
	// updateUser(userRepository)
	// deleteUser(userRepository)
	// getUser("U1", userRepository)
	// getUsers(userRepository)
}

// CREATE
func saveUser(userRepository repository.UserRepository)  {
	var p model.User
	// p.ID = "U3"
	p.FirstName = "Rafaela"
	p.LastName = "Rodrigues"
	p.Email = "rafaela@rafaela.com.br"

	err := userRepository.Save(&p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Usuário cadastrado.")
	}
}

// UPDATE
func updateUser(userRepository repository.UserRepository)  {
	var p model.User
	// p.ID = "U1"
	p.FirstName = "Rafaela"
	p.LastName = "Rodrigues"
	p.Email = "rafaela_sr@rafaela.com"

	err := userRepository.Update("U1", &p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Usuário atualizado.")
	}
}

// DELETE
func deleteUser(userRepository repository.UserRepository) {
	err := userRepository.Delete("U1")
	
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Usuário deletado.")
	}
}

// READ ONE
func getUser(id string, userRepository repository.UserRepository)  {
	user, err := userRepository.FindByID(id)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ID: " + user.ID.String())
		fmt.Println("Nome: " + user.FirstName)
		fmt.Println("Sobrenome: " + user.LastName)
		fmt.Println("Email: " + user.Email)
	}
}

// READ ALL
func getUsers(userRepository repository.UserRepository) {
	users, err := userRepository.FindAll()

	if err != nil {
		fmt.Println(err)
	}

	for _, user := range users {
		fmt.Println()
		fmt.Println("ID: " + user.ID.String())
		fmt.Println("Nome: " + user.FirstName)
		fmt.Println("Sobrenome: " + user.LastName)
		fmt.Println("Email: " + user.Email)
	}
}