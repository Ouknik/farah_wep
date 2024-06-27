package helpers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"app.roya_immobile.com/lib"
	"app.roya_immobile.com/models"
	"app.roya_immobile.com/persistence"
)

type HelperCloudToken struct {
	_modelFileservetoken   models.ModelFileservetoken
	_persistenceCloudToken persistence.CloudToken
	cloudUrl               string
	cloudAction            persistence.CloudAction
	fileSize               int64
}

func (h *HelperCloudToken) Init(dbSqlConnection *sql.DB, _handleRedirectAndPanic *lib.HandleRedirectAndPanic, cloudUrl string) {
	h.makeRequiredDbModels(dbSqlConnection, _handleRedirectAndPanic)
	h._persistenceCloudToken = persistence.CloudToken{}

	h.cloudUrl = cloudUrl
}

func (h *HelperCloudToken) GetDeleteCloudToken(fileName, portalKey string) (persistence.CloudToken, error) {
	var token, err = h._modelFileservetoken.FindByFileTokenForDelete(portalKey, fileName)
	if err != nil {
		return persistence.CloudToken{}, err
	}

	h._persistenceCloudToken.FileName = fileName
	h._persistenceCloudToken.Token = token
	h._persistenceCloudToken.CloudUrl = h.makeCloudUrl(token, h.cloudAction.Delete())
	return h._persistenceCloudToken, nil
}

func (h *HelperCloudToken) GetUploadCloudToken(fileName, filePath, destination, userId, portalKey string) (persistence.CloudToken, error) {

	var fileError = h.makeFileDetails(fileName, filePath)
	if fileError != nil {
		return persistence.CloudToken{}, fileError
	}

	// make the 16 dig token
	var token = lib.GenerateRandomAlphaNumericString(16)
	// append portal key and user id
	h._persistenceCloudToken.PortalKey = portalKey
	h._persistenceCloudToken.CreatedBy = userId
	h._persistenceCloudToken.Token = token
	h._persistenceCloudToken.Destination = destination
	h._persistenceCloudToken.CloudUrl = h.makeCloudUrl(token, h.cloudAction.Upload())
	h.dumpFileDetailsInDb()

	return h._persistenceCloudToken, nil

}

func (h *HelperCloudToken) dumpFileDetailsInDb() {
	h._modelFileservetoken.InsertFileDetails(h._persistenceCloudToken)
}

func (h *HelperCloudToken) makeRequiredDbModels(dbSqlConnection *sql.DB, _handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	h._modelFileservetoken = models.ModelFileservetoken{}
	h._modelFileservetoken.InitWithoutConnection(dbSqlConnection, _handleRedirectAndPanic)
}

func (h *HelperCloudToken) getFileSize(filePath string) (int64, error) {
	fileStats, err := os.Stat(filePath)
	if err != nil {
		log.Println("File stats could not be read:", err)
		return -1, err
	}
	return fileStats.Size(), nil
}

func (h *HelperCloudToken) makeFileDetails(fileName, filePath string) error {

	fileSize, err := h.getFileSize(fmt.Sprintf("%s/%s", filePath, fileName))
	if err != nil {
		log.Println(err)
		return err
	}
	// set file size globally to avoid empty content-length problem
	h.fileSize = fileSize

	fileType, err := h.getFileContentType(fmt.Sprintf("%s/%s", filePath, fileName))
	if err != nil {
		log.Println(err)
		return err
	}

	h._persistenceCloudToken.FileName = fileName
	h._persistenceCloudToken.FileSize = lib.Int64ToStr(fileSize)
	h._persistenceCloudToken.FileMimeType = fileType

	return nil

}

func (h *HelperCloudToken) getFileContentType(filePath string) (string, error) {
	ouput, err := os.Open(filePath)
	if err != nil {
		log.Println("file could not be opened", err)
		return "", err
	}
	defer ouput.Close()

	// to sniff the content type only the first
	// 512 bytes are used.
	buf := make([]byte, 512)
	_, err = ouput.Read(buf)
	if err != nil {
		return "", err
	}
	return http.DetectContentType(buf), nil
}

func (h *HelperCloudToken) makeCloudUrl(cloudToken string, cloudAction persistence.CloudAction) string {
	return fmt.Sprintf("%s/%s/%s", h.cloudUrl, cloudAction, cloudToken)
}
