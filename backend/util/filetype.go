package util

import (
	"net/http"
	"slices"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jcxldn/fosscat/backend/structs"
)

var imageFileTypes = []string{"image/jpeg", "image/png"}

func GetFileTypeForFile(file graphql.Upload) (structs.FileTypes, error) {
	buf := make([]byte, 512) // only need first 512 bytes to determine file type per docs

	// Read the first 512 bytes of the file
	_, err := file.File.Read(buf)

	// Seek back to the beginning of the file
	// so that future reads (such as saving in the DB) contain the whole file
	// Inspired to a similar problem at: https://stackoverflow.com/a/61729136
	file.File.Seek(0, 0)

	if err == nil {
		// Read file successfully
		contentType := http.DetectContentType(buf)
		if slices.Contains(imageFileTypes, contentType) {
			return structs.FileTypesImage, nil
		} else {
			return structs.FileTypesGeneric, nil
		}
	} else {
		return structs.FileTypesErrorDetermining, err
	}
}
