package main

import "fmt"

func main() {
	fruits := make([]string, 1)
	fruits[0] = "Apple"

	rujak := []string{"Pineapple", "Cucumber", "Mango", "Watermelon"}

	nn := copy(fruits, rujak)

	sliceOfRujak := rujak[1:3]
	sliceOfRujak2 := rujak[:3]

	fmt.Printf("Fruits: %#v\n", fruits)
	fmt.Printf("Rujak: %#v\n", rujak)
	fmt.Printf("Copied: %d\n", nn)
	fmt.Printf("Fruits length: %d\n", len(fruits))
	fmt.Printf("Fruits capacity: %d\n", cap(fruits))
	fmt.Printf("Rujak length: %d\n", len(rujak))
	fmt.Printf("Rujak capacity: %d\n", cap(rujak))

	fmt.Printf("Slice of Rujak: %#v\n", sliceOfRujak)

	fmt.Printf("Slice of Rujak length: %d\n", len(sliceOfRujak))
	fmt.Printf("Slice of Rujak capacity: %d\n", cap(sliceOfRujak))

	fmt.Printf("Slice of Rujak 2: %#v\n", sliceOfRujak2)
	fmt.Printf("Slice of Rujak 2 length: %d\n", len(sliceOfRujak2))
	fmt.Printf("Slice of Rujak 2 capacity: %d\n", cap(sliceOfRujak2))
}
