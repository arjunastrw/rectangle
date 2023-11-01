package main

import (
	"fmt"

	"github.com/arjunastrw/rectangle/shape"
)

func main() {
	firstFloor := shape.Rectangle{Width: 7.5, Length: 6.5}
	secondFloor := shape.Rectangle{Width: 4.5, Length: 5.5}

	totalArea := firstFloor.Area() + secondFloor.Area()

	fmt.Println("Total Area :", totalArea)

}
