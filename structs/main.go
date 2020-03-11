package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal      uint16 //min 0 max 65535
	brake_pedal    uint16
	steering_wheel int16 //-32k to 32k
	top_speed_kmh  float64
}

// Value receiver method
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}
func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax / kmh_multiple)
}

//pointer receiver
/* In simple terms, value receiver makes a copy of the type and pass it to the function. The function stack now holds an equal object but at a different location on memory. Pointer receiver passes the address of a type to the function. The function stack has a reference to the original object */
func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

func main() {
	car1 := car{gas_pedal: 22341, brake_pedal: 0, steering_wheel: 12561, top_speed_kmh: 225.0}

	car2 := car{41046, 0, 25040, 300.16}

	fmt.Println(car1.top_speed_kmh)
	fmt.Println(car2.gas_pedal)
	fmt.Println(car1.kmh())
	fmt.Println(car2.mph())

	car1.new_top_speed(500)
	fmt.Println(car1.kmh())
	fmt.Println(car1.mph())
	//methods in go
	//
}
