package controllers

import (
	"fmt"

	"net/http"

	"app.roya_immobile.com/helpers"
	"app.roya_immobile.com/models"

	"app.roya_immobile.com/persistence"

	"app.roya_immobile.com/lib"
)

type ControllerHome struct {
	_handleRedirectAndPanic *lib.HandleRedirectAndPanic
	_handleFilePath         *lib.HandleFilePath
	_handleSession          *lib.HandleSession
	_templateEngine         lib.TemplateEngine
	_translator             lib.Translator_i18n
	_helperHeaderDetails    helpers.HelperHeaderDetails
	// Set Model if requiredfile:///media/ouknik/7432A60D32A5D480/thechpool/techpool-meinport-app/persistence/struct_user_details.go

	_modelAnnonce models.ModelAnnonce
	_modelUser    models.ModelUserapp

	_flashMessages                lib.HandleFlashMessages
	_userDetails                  persistence.UserDetails
	doNotRedirectUserOnFailToAuth bool
}

// trigger a dummy fmt to avoid
// 'imported but not used' error
var _ = fmt.Errorf("")

func (c *ControllerHome) init(action string,
	_getVars map[string]string,
	r *http.Request,
	w http.ResponseWriter,
	_handleFilePath *lib.HandleFilePath,
	_handleSession *lib.HandleSession,
	_handleRedirectAndPanic *lib.HandleRedirectAndPanic,
	_envConfigVars *lib.GetEnvConfigVars) bool {

	// First initaite the Panic Handler
	c._handleRedirectAndPanic = _handleRedirectAndPanic

	// Set the FilePath Class
	c._handleFilePath = _handleFilePath

	// Set Session Handler
	c._handleSession = _handleSession

	// Initiate the Model, if required
	c._modelAnnonce = models.ModelAnnonce{}

	// Requires only Panic Handler DbTable
	// and Prefix are to be setted in Model
	c._modelAnnonce.Init(c._handleRedirectAndPanic)

	// set the translator
	c._translator = lib.Translator_i18n{}
	c._translator.InitTranslator(c._handleFilePath, c._handleSession, c._handleRedirectAndPanic)

	// Initiate the Template Engine
	c._templateEngine = lib.TemplateEngine{}
	c._templateEngine.CreateTemplate(_getVars, c._handleFilePath, w, &c._translator)

	// Init FlashMessages
	c._flashMessages = lib.HandleFlashMessages{}
	c._flashMessages.Init(c._handleSession)

	// Set all flash Messages to view
	c._templateEngine.ViewData["FlashMessages"] = c._flashMessages.GetAllFlashMessages()

	return c._handleRedirectAndPanic.HasPanicError
}

func (c *ControllerHome) ActionManage(_getVars map[string]string,
	r *http.Request,
	w http.ResponseWriter,
	_handleFilePath *lib.HandleFilePath,
	_handleSession *lib.HandleSession,
	_handleRedirectAndPanic *lib.HandleRedirectAndPanic,
	_envConfigVars *lib.GetEnvConfigVars) {
	// redirect user if authentification fails?
	c.doNotRedirectUserOnFailToAuth = true
	// Initiate Controller & look for Panic Error
	if c.init("manage", _getVars, r, w, _handleFilePath, _handleSession, _handleRedirectAndPanic, _envConfigVars) {
		c._handleRedirectAndPanic.RedirectToPanicUrl()
		return
	}

	annonceFromDb := c._modelAnnonce.GetAllAnnonceForView()

	//fmt.Println(annonceFromDb)
	//fmt.Println(annonceFromDb)

	c._templateEngine.ViewData["Annonces"] = annonceFromDb
	// fmt.Println(c.restaurantFromDb)
	// make iteration keys and
	// assign iteration keys to view
	keysForIteration := lib.StripTableFieldPrefixFromKeysForView(c._modelAnnonce.DBAdapter.GetDbTableFieldPrefix(), c._modelAnnonce.GetDbFieldsAsKeysForIteration(), "Annonces_")
	c._templateEngine.ViewData["KeysForIteration"] = keysForIteration

	// Final Panic Error Check before irreverseable
	// header is created
	if c._handleRedirectAndPanic.HasPanicError {
		c._handleRedirectAndPanic.RedirectToPanicUrl()
		return
	}
	c._templateEngine.Execute("home", "manage")

}
