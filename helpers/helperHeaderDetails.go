package helpers

import (
	"app.roya_immobile.com/lib"
)

type HelperHeaderDetails struct {
	_handleFilePath *lib.HandleFilePath
	_handleSession  *lib.HandleSession
	headerDetails   map[string]string
}

func (h *HelperHeaderDetails) Init(_handleFilePath *lib.HandleFilePath, _handleSession *lib.HandleSession) {
	h._handleFilePath = _handleFilePath
	h._handleSession = _handleSession
	h.headerDetails = map[string]string{}

	h.SetPublicPath(h._handleFilePath.GetPublicPath())
	h.SetThemePath(h._handleFilePath.GetThemePath())
	//h.SetProfilePicCDNPath_Large(h._handleFilePath.GetProfilePicCDNPath_Large())
	//h.SetProfilePicCDNPath_Thumbnails(h._handleFilePath.GetProfilePicCDNPath_Thumbnails())
	h.SetLang(h._handleSession.GetLanguageFromCookie())
}

func (h *HelperHeaderDetails) SetLang(lang string) {
	h.headerDetails["Lang"] = lang
}

func (h *HelperHeaderDetails) SetKeywords(keywords []string) {
	h.headerDetails["Keywords"] = lib.Implode(keywords, ", ")
}

func (h *HelperHeaderDetails) SetDescription(description string) {
	h.headerDetails["Description"] = description
}

func (h *HelperHeaderDetails) SetPageTitle(pageTitle string) {
	h.headerDetails["PageTitle"] = pageTitle
}

func (h *HelperHeaderDetails) GetHeaderDetails() map[string]string {
	return h.headerDetails
}

// depricated
// use {{.EnvConfVars}} instead
// it is generic and safe as this
// method considers sensible params
func (h *HelperHeaderDetails) SetProfilePicCDNPath_Thumbnails(path string) {
	h.headerDetails["ProfilePicCDNPath_Thumbnails"] = path
}

// depricated
// use {{.EnvConfVars}} instead
// it is generic and safe as this
// method considers sensible params
func (h *HelperHeaderDetails) SetPublicPath(path string) {
	h.headerDetails["PublicPath"] = path
}

// depricated
// use {{.EnvConfVars}} instead
// it is generic and safe as this
// method considers sensible params
func (h *HelperHeaderDetails) SetProfilePicCDNPath_Large(path string) {
	h.headerDetails["ProfilePicCDNPath_Large"] = path
}

// depricated
// use {{.EnvConfVars}} instead
// it is generic and safe as this
// method considers sensible params
func (h *HelperHeaderDetails) SetThemePath(path string) {
	h.headerDetails["ThemePath"] = path
}
