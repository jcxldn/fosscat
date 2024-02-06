package database

import (
	"io"

	"github.com/google/uuid"
	"github.com/jcxldn/fosscat/backend/graph/model"
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

func UploadFiles(db *gorm.DB, files []*model.UploadFile, uploader structs.User) ([]*structs.File, error) {
	fileStructs := []*structs.File{}
	for _, fileModel := range files {
		fileType, err := util.GetFileTypeForFile(fileModel.File)
		if err == nil && fileType == structs.FileTypesImage {
			// Create a new file struct with a unique UUID
			file := createFileStruct(db)

			// Set other fields
			file.Name = fileModel.File.Filename
			file.Uploader = uploader
			file.Type = structs.FileTypesImage
			data, readErr := io.ReadAll(fileModel.File.File)
			if readErr != nil {
				return nil, err
			}
			file.Data = data

			// Add to database
			db.Create(&file)

			// Add to fileStructs
			fileStructs = append(fileStructs, file)
		} else {
			return nil, err
		}
	}

	return fileStructs, nil
}
