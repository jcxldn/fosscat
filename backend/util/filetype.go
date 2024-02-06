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
	if _, err := file.File.Read(buf); err == nil {
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
