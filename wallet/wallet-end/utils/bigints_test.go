package utils

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"math/big"
	"testing"
)

func TestUint8ArrayToBigInt(t *testing.T) {
	x := []byte{229, 195, 87, 247, 232, 92, 48, 202, 171, 203, 143,
		154, 125, 216, 139, 89, 17, 78, 40, 189, 118, 204,
		144, 14, 230, 191, 57, 17, 36, 181, 224, 110, 67,
		133, 4, 26, 156, 175, 182, 67, 115, 247, 215, 221,
		40, 187, 225, 111, 98, 13, 70, 67, 36, 42, 231,
		181, 155, 234, 94, 46, 214, 72, 100, 131}

	bigint := Uint8ArrayToBigInt(x)

	assert.Equal(t,
		"12033667936854278646467929070026001295165595627649794109200610712065050964767432450763253419115894771589000955150062892551461416593509556352107661459219587",
		bigint.String())
}

func TestBigIntToArray(t *testing.T) {
	x, _ := new(big.Int).SetString("12033667936854278646467929070026001295165595627649794109200610712065050964767432450763253419115894771589000955150062892551461416593509556352107661459219587", 10)

	array := BigIntToArray(86, 3, x)

	fmt.Println(array)
}