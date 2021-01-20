package bottom

import (
	"fmt"
)

func main() {
	encodeTest := encode("test")
	fmt.Println(encodeTest)

	decodeTest, err := decode("💖💖✨🥺,👉👈💖💖,👉👈💖💖✨🥺👉👈💖💖✨🥺,👉👈")
	if err != nil {
		return
	}
	fmt.Println(decodeTest)

	decodeTest2, err := decode("erroneous data")
	if err != nil {
		fmt.Println("failed to decode.")
		return
	}
	fmt.Println(decodeTest2)
}
