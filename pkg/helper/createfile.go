package helper

import (
	"bufio"
	"fmt"
	"os"
)

/*
create a file with the ascii art wanted from the user with the selected file extension type
*/

func Createfile(fileExtension, result string) error {
	file, err := os.Create("file-request/file." + fileExtension)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	fmt.Fprint(writer, result)
	return writer.Flush()
}
