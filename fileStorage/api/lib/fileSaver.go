package lib

import (
	"compress/gzip"
	"os"
	"path/filepath"
	"strings"
)

func FileSaver(file []byte, path string) error {
	
	path = strings.Replace(path, filepath.Ext(path), ".gz", -1)

	dst, err := os.Create(path)
	if err != nil {
		ErrorHandler(err, "system")
		return err
	}
	defer dst.Close()

	w := gzip.NewWriter(dst)
	w.Write(file)

	w.Close()

	return nil
}
