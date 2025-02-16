package utils

import (
	"math/rand"
	"strconv"
)

func RandomNumber() string {
	return strconv.Itoa(rand.Intn(900000)+100000);
}