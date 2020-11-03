package encrypt

import (
	"fmt"
	"log"
	"testing"
)

func Test1(t *testing.T) {
	//crypted, _ := DesEncrypt("")
	//fmt.Println(crypted)
	decrypted, _ := DesDeCrypt("1dd93979e9d0715062fdfd4e3dc041f1")
	fmt.Println(decrypted)
}

func Test2(t *testing.T) {
	decryted, err := DesEncrypt("daipengyuan")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(decryted)
}
