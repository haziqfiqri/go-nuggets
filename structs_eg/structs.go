package structs_eg

import (
	"fmt"
	"time"
)

type User struct {
	name string
	age int
	isMember bool
	address string
	createdAt time.Time
}

type Player struct {
	name string
	online bool
	Match
}

type Match struct {
	id string
	score int
	date time.Time
	player []*Player
}

func (m *Match) UpdateMatch(v int) { 
	m.score = v
}

func Structs() {

	newUser := User{
		name: "John",
		age: 25,
		isMember: true,
		createdAt: time.Now(),
	}


	fmt.Println("newUser", newUser)
	fmt.Println("newUser.name", newUser.name)

	newPtr := &newUser
	fmt.Println("newPtr", newPtr)
	fmt.Println("newPtr.name", newPtr.name)

	newPtr.age = 30
	fmt.Println("newPtr.age", newPtr.age)

	fmt.Println("User", User{
		name:      "Haziq",
		age:       29,
		isMember:  true,
		createdAt: time.Now(),
	})

	// anonymous struct
	customResponse := struct {
		statusCode int
		message string
	} {
		statusCode: 422,
		message: "cheating",
	}

	fmt.Println("customResponse", customResponse)

	newMatch := Match{
		id: "1",
		date: time.Now(),
		player: []*Player{
			{
				name: "John",
				online: true,
			},
			{
				name: "Haziq",
				online: false,
			},
		},
	}

	

	fmt.Println("newMatch", newMatch)

	newMatch.UpdateMatch(10)
	fmt.Println("newMatch", newMatch)

	return
}