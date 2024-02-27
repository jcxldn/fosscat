package database

import (
	"io"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/jcxldn/fosscat/backend/structs"
	"github.com/jcxldn/fosscat/backend/util"
	"gorm.io/gorm"
)

func createFileStruct(db *gorm.DB) *structs.File {
	file := structs.File{}

	isFreeUuid := false
	for !isFreeUuid {
		// Generate a UUID for the user id.
		file.ID = uuid.New()
		// Check that the UUID has not been used already
		// If true, it will break out of this for loop and continue.
		isFreeUuid = util.IsUuidFree[structs.Entity](db, file.ID)
	}

	return &file
}

func UploadFiles(db *gorm.DB, files []*multipart.FileHeader, uploader structs.User) ([]*structs.File, error) {
	fileStructs := []*structs.File{}
	for _, file := range files {
		fileReader, err1 := file.Open()
		if err1 == nil {
			fileType, err := util.GetFileTypeForFile(fileReader)
			if err == nil && fileType == structs.FileTypesImage {
				// Create a new file struct with a unique UUID
				fileStruct := createFileStruct(db)

				// Set other fields
				fileStruct.Name = file.Filename
				fileStruct.Uploader = uploader
				fileStruct.Type = structs.FileTypesImage
				data, readErr := io.ReadAll(fileReader)
				if readErr != nil {
					return nil, err
				}
				fileStruct.Data = data

				// Add to database
				db.Create(&fileStruct)

				// Add to fileStructs
				fileStructs = append(fileStructs, fileStruct)
			} else {
				return nil, err
			}
		} else {
			return nil, err1
		}
	}

	return fileStructs, nil
}
