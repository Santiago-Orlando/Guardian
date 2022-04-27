package lib

import "os"


func DeleteFile(path string) error {

	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}