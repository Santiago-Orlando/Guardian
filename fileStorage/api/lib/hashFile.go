package lib

import (
	"crypto/sha256"
	"fmt"
	//"io"
	//"io/ioutil"
)

func HashFile(file []byte) string {

	sha256 := sha256.Sum256(file)

	return fmt.Sprintf("%x", sha256)
}
