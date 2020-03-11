package main
import "fmt"

func main() {
	grades := make(map[string]float32)

	grades["Timmy"] = 42
	grades["Jess"] = 88
	grades["Mike"] = 67

	fmt.Println(grades)

	TimsGrade := grades["Timmy"]
	fmt.Println(TimsGrade)

	delete(grades,"Timmy")
	fmt.Println(grades)

	for k,v := range grades {
		fmt.Println(k,":",v)
	}
}