package main

import "crypto/rand"
import "fmt"
import "flag"
import "time"

func NewPassword(pLen int) string {

	vRis := make([]byte, pLen)

	rand.Read(vRis)
	for vCnt := 0; vCnt < pLen; vCnt++ {
		vRis[vCnt] = vRis[vCnt]%93 + 33
		//vRis[vCnt] = byte(33 + rand.Intn(93))
	}

	return string(vRis)
}

func main() {

	vLengthParamater := flag.Int("length", 16, "Password length")
	vCountParamater := flag.Int("count", 20, "Numer of password to generate")
	flag.Parse()

	for vCnt := 0; vCnt < *vCountParamater; vCnt++ {
		fmt.Println(NewPassword(*vLengthParamater))
		time.Sleep(time.Millisecond)
	}
}
