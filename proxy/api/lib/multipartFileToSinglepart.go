package lib

import (
	"bytes"
	"io"
	"mime/multipart"
)

func MultipartToSinglepart(multi multipart.File) []byte {

	buf := bytes.NewBuffer(nil)

	_, err := io.Copy(buf, multi)
	if err != nil {
		ErrorHandler(err, "system")
	}

	return buf.Bytes()
}
