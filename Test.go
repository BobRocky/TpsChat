/*package main

import "fmt"

type Vehicle interface {
	move()
}

// структура "Автомобиль"
type Car struct{}

// структура "Самолет"
type Aircraft struct{}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}
func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

func main() {

	var tesla Vehicle = Car{}
	var boing Vehicle = Aircraft{}
	tesla.move()
	boing.move()
}
package main

import "fmt"

type Car struct{}
type Aircraft struct{}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}
func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

func driveCar(c Car) {
	c.move()
}
func driveAircraft(a Aircraft) {
	a.move()
}

func main() {

	var tesla Car = Car{}
	var boing Aircraft = Aircraft{}
	driveCar(tesla)
	driveAircraft(boing)
}

package main

import "fmt"

type Vehicle interface {
	move()
}

func drive(vehicle Vehicle) {
	vehicle.move()
}

type Car struct{}
type Aircraft struct{}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}
func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

func main() {

	tesla := Car{}
	boing := Aircraft{}
	drive(tesla)
	drive(boing)
}*/
