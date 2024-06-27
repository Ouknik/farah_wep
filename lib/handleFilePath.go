package lib

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"app.roya_immobile.com/persistence"

	//	"strings"
	"sync"
)

type HandleFilePath struct {
	_envConfigVars  persistence.EnvConfigVars
	applicationName string
}

func (h *HandleFilePath) GetApplicationPath() string {

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	return h.CleanPath(fmt.Sprintf("%s/../src/%s", dir, h.applicationName))

}

func (h *HandleFilePath) GetSrcPath() string {
	rootDir, _ := os.Getwd()
	fmt.Println(rootDir)
	return rootDir
}

func (h *HandleFilePath) GetPublicPath() string {
	return "/public"
}

func (h *HandleFilePath) GetPrivatePath() string {
	return "./storage_private"
}

/*
func (h *HandleFilePath) GetPublicStoragePath() string {
	return h._envConfigVars.EnvConfigVarList["publicStoragePath"]
}*/

func (h *HandleFilePath) GetAnnoncePicPath_Large() string {
	return "./public/annoncepic_large"
}

func (h *HandleFilePath) GetAnnoncePicPath_Thumbnails() string {
	return "./public/annoncepic_thumbnails"
}

func (h *HandleFilePath) GetProfilePicPath_Large() string {
	return "/public/profilepic_large"
}

func (h *HandleFilePath) GetProfilePicPath_Thumbnails() string {
	return "/public/profilepic_thumbnails"
}

/*
func (h *HandleFilePath) GetProfilePicCDNPath_Large() string {
	return h._envConfigVars.EnvConfigVarList["profilePicCDNPath_large"]
}*/

/*
	func (h *HandleFilePath) GetProfilePicCDNPath_Thumbnails() string {
		return h._envConfigVars.EnvConfigVarList["profilePicCDNPath_thumbnails"]
	}
*/
func (h *HandleFilePath) GetAnnoncePicCDNPath_Large() string {
	return "/public/annoncepic_large"
}

func (h *HandleFilePath) GetAnnoncePicCDNPath_Thumbnails() string {
	return "/public/annoncepic_thumbnails"
}

/*
func (h *HandleFilePath) GetRestaurantImgPath_Large() string {
	return h._envConfigVars.EnvConfigVarList["restaurantImgStorage_large"]
}

func (h *HandleFilePath) GetRestaurantImgPath_Thumbnails() string {
	return h._envConfigVars.EnvConfigVarList["restaurantImgStorage_thumbnails"]
}

func (h *HandleFilePath) GetRestaurantImgCDNPath_Large() string {
	return h._envConfigVars.EnvConfigVarList["restaurantImgCDNPath_large"]
}

func (h *HandleFilePath) GetRestaurantImgCDNPath_Thumbnails() string {
	return h._envConfigVars.EnvConfigVarList["restaurantImgCDNPath_thumbnails"]
}
*/

// a common stroage for all large
// images from company logos to
// event images
/*
func (h *HandleFilePath) GetCommonPublicPath_Large() string {
	return h._envConfigVars.EnvConfigVarList["commonPublicStorage_large"]
}
*/
// a common stroage for all image
// thumbnails from company logos
// to event images
/*
func (h *HandleFilePath) GetCommonPublicPath_Thumbnails() string {
	return h._envConfigVars.EnvConfigVarList["commonPublicStorage_thumbnails"]
}
*/
func (h *HandleFilePath) GetThemePath() string {
	return "/public/themes/farah"
}

func (h *HandleFilePath) GetTemplatePath() string {
	return "./templates/farah"
}

func (h *HandleFilePath) GetLanguagePath() string {
	//return fmt.Sprintf("%s/i18n_languages/", h.GetApplicationPath())
	return fmt.Sprintf("%s/i18n_languages/", h.GetSrcPath())
}

func (h *HandleFilePath) GetFormPath() string {
	//return fmt.Sprintf("%s/forms/", h.GetApplicationPath())
	return fmt.Sprintf("%s/forms/", h.GetSrcPath())
}

func (h *HandleFilePath) CleanPath(pathToClean string) string {

	var t, _ = filepath.Abs(pathToClean)
	return t

}

func (h *HandleFilePath) SetEnvConfigVars(applicationName string, envConfigFilePath string) {

	var _envConf = new(GetEnvConfigVars)

	h._envConfigVars = _envConf.Initiate(envConfigFilePath)

	h.applicationName = applicationName

}

// Check if file exists
func (h *HandleFilePath) FileExists(fileWithAbsolutPath string) bool {
	if _, err := os.Stat(fileWithAbsolutPath); err == nil {
		// File exists
		return true
	}
	// something might be wrong
	return false
}

func (h *HandleFilePath) GetFileSize(filePath string) (int64, error) {

	fi, e := os.Stat(filePath)
	if e != nil {
		return 0, e
	}
	return fi.Size(), nil

}

func (h *HandleFilePath) RemoveExtensionFromFileName(filename string) string {
	return TrimSuffix(filename, filepath.Ext(filename))
}

func (h *HandleFilePath) CreateEmptyFile(filename, path string) error {
	// If the file doesn't exist, create it, or append to the file
	// fmt.Printf("%s/%s\n", path, filename)
	f, err := os.OpenFile(fmt.Sprintf("%s/%s", path, filename), os.O_CREATE, 0644)
	// as the file object is not required anymore
	// close file object immediatly and not defer
	f.Close()
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}
	return err
}

func (h *HandleFilePath) WriteNotAppendContentToFile(fileToWrite, contentToWrite string, waitGroup sync.WaitGroup) {
	if !h.FileExists(fileToWrite) {
		return
	}

	waitGroup.Add(1)

	// overwrite file if it exists
	file, err := os.OpenFile(fileToWrite, os.O_RDWR, 0644)

	// close fo on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	datawriter := bufio.NewWriter(file)

	datawriter.WriteString(contentToWrite)

	if err = datawriter.Flush(); err != nil {
		panic(err)
	}

	waitGroup.Done()
}

// move or rename file, basically
// executes the mv (linux move)
// command
func (h *HandleFilePath) MoveFile(source, destination string) {
	// move file
	var err = os.Rename(source, destination)
	if err != nil {
		fmt.Printf("%v", err)
	}
}
