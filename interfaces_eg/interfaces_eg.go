package interfaces_eg

import "fmt"

type Animal interface {
	Walk()
	Eat(food string)
}

type BaseAnimal struct {
	Name string
}

func (b *BaseAnimal) Eat(food string) {
	fmt.Printf("😋 %s is eating %s (generic eating logic)\n", b.Name, food)
}

type Dog struct {
	BaseAnimal
	Breed string
}

func (d *Dog) Walk() {
	fmt.Printf("🐕 %s the %s is running fast!\n", d.Name, d.Breed)
}

type Cat struct {
	BaseAnimal
}

func (c *Cat) Walk() {
	fmt.Printf("🐈 %s is walking stealthily...\n", c.Name)
}

func (c *Cat) Eat(food string) {
	fmt.Printf("🐈 %s is sniffing the %s suspiciously before eating.\n", c.Name, food)
}

func PerformDailyRoutine(a Animal, food string) {
	fmt.Printf("\n--- Routine for %T ---\n", a)
	a.Walk()
	a.Eat(food)
}

func Interfaces() {
	dog := &Dog{
		BaseAnimal: BaseAnimal{Name: "Buddy"},
		Breed:      "Retriever",
	}

	cat := &Cat{
		BaseAnimal: BaseAnimal{Name: "Whiskers"},
	}

	fmt.Println("Accessing Name directly via promotion:", dog.Name)

	PerformDailyRoutine(dog, "bone")
	PerformDailyRoutine(cat, "tuna")
}