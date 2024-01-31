package structs

import (
	"fmt"
)

type FileTypes string

const (
	FileTypesErrorDetermining FileTypes = "ERROR_DETERMINING"
	FileTypesGeneric          FileTypes = "GENERIC"
	FileTypesImage            FileTypes = "IMAGE"
)

func (e *FileTypes) GetFileTypeForString(str string) (FileTypes, error) {
	switch str {
	case string(FileTypesErrorDetermining):
		return FileTypesErrorDetermining, nil
	case string(FileTypesGeneric):
		return FileTypesGeneric, nil
	case string(FileTypesImage):
		return FileTypesImage, nil
	}
	return FileTypesErrorDetermining, fmt.Errorf("unable to determine file type from type string '%s'", str)
}
