package variables

import (
	"fmt"
)

func VariablesEnteras() {
	var miVarInt int
	miVarInt64 := int64(12)
	miVarInt16 := int16(32)
	
	fmt.Println("variable int definida por default", miVarInt)
	fmt.Println("variable int definida en 64 bits", miVarInt64)
	fmt.Println("variable int definida en 16 bits", miVarInt16)
}