package lib

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
)

func MultipartToSinglepart(multi multipart.File) []byte {

	buf := bytes.NewBuffer(nil)

	_, err := io.Copy(buf, multi)
	if err != nil {
		fmt.Println(err)
	}

	return buf.Bytes()
}
