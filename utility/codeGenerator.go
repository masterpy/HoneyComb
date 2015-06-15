package utility

import (
	//"fmt"
	"math/rand"
	"strconv"
	"time"
)

//Generate a random integer
//use time.Now().UTC().UnixNano() as Seed
func RandomInt() int {

	magicNumber := 10000
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(magicNumber)
}

//I think this is good enouth
//if this not work, you can add ip
func GenerateCode(origin string) string {
	str := origin + strconv.Itoa(RandomInt())
	//fmt.Println("GENERATE CODE" + strconv.Itoa(RandomInt()))
	result := string(Md5sum(str))
	return result
}
