package main

import (
	//load the base ;-)
	"errors"
	"fmt"

	// load network dependencies
	"net/http"

	// load the controllers
	controller "app.roya_immobile.com/controllers"
	"github.com/gorilla/mux"

	// load the lib
	lib "app.roya_immobile.com/lib"

	// load the persistence
	"app.roya_immobile.com/persistence"
)

// initiate global vars
var (

	// @ string
	// set application name
	applicationName string

	// @ gorilla/mux
	// create the router
	muxRouting = mux.NewRouter()

	// @ lib/HandleFilePath
	// set handle file path lib globally
	_handleFilePath lib.HandleFilePath

	// @ persistence/EnvConfigVarsFilePath
	// load the environment variables
	envConfigFilePath string = persistence.EnvConfigVarsFilePath

	// sensible config vars
	sensibleConfigVars []string

	// @ lib/HandleSession
	// handle session globally
	_handleSession lib.HandleSession

	_envConfigVars lib.GetEnvConfigVars

	_handleRedirectAndPanic lib.HandleRedirectAndPanic
)

func main() {

	// init fmt just in case if you need
	// or simply disable
	var _ = fmt.Errorf

	//set file path class globally to handle getPath class
	setHandleFilePath()

	// set env config vars
	setEnvConfigVars()

	// declare config vars which are sensible
	sensibleConfigVars = []string{
		"applicationSrcName",
		"port",
		"publicStorage",
		"privateStorage",
		"templatePath",
		"emailTemplatePath",
		"smtpDetails",
		"validSessionTimeInMins",
		"validStaticSessionTimeInDays",
	}

	// set sensible vars in config handler
	declareSensibleConfigVars()

	// set application name
	setApplicationName()
	//	fmt.Println(applicationName)

	muxRouting = mux.NewRouter()

	// This will serve files under http://localhost:8000/static/<filename>
	// muxRouting.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir(publicDir))))
	s := http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))
	muxRouting.PathPrefix("/public/").Handler(s)

	http.Handle("/", muxRouting)

	//the base is dashboard
	//muxRouting.HandleFunc("/", routerDefault)

	// handle the routing
	handleRouting()

	var port = "9092"

	printDBDetails()

	// _envConfigVars.PrintAllConfigVars()

	printPortDetails(port)

	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%s", port), RecoverWrap(muxRouting)))

	//http.ListenAndServe(_envConfigVars.GetConfVar("port"), RecoverWrap(muxRouting))

} // end main

func printPortDetails(port string) {

	fmt.Println("-------------------------")
	fmt.Println("Running on port", port)
	fmt.Printf("Visit: %s\n", _envConfigVars.GetConfVar("baseURL"))
	fmt.Println("-------------------------")
	fmt.Println()

}

func printDBDetails() {
	var _dbAdapter = new(lib.DbAdapter)
	_dbAdapter.Connect("", "", &_handleRedirectAndPanic)
	_dbAdapter.PrintDBDetails()
}

func setHandleFilePath() {

	_handleFilePath = lib.HandleFilePath{}
	_handleFilePath.SetEnvConfigVars(applicationName, envConfigFilePath)

}

func setHandleSession(w http.ResponseWriter, r *http.Request) {

	_handleSession = lib.HandleSession{}
	_handleSession.Init(w, r)

}

func setEnvConfigVars() {
	_envConfigVars = lib.GetEnvConfigVars{}
	_envConfigVars.Initiate(envConfigFilePath)
}

func declareSensibleConfigVars() {
	_envConfigVars.SetSensibleConfigSecurityVars(sensibleConfigVars)
}

func setApplicationName() {
	applicationName = "farah_web_app"
}

func setHandleRedirectAndPanic(w http.ResponseWriter, r *http.Request) {
	_handleRedirectAndPanic = lib.HandleRedirectAndPanic{}
	_handleRedirectAndPanic.SetHttpParams(w, r)
}

func handleRouting() {

	// /deviceinfo/put
	muxRouting.HandleFunc("/home/manage", routerHomeManage)
	muxRouting.HandleFunc("/home/manage/", routerHomeManage)

	/*------------------------- address route start -------------------------------------*/
	// /address/addedit

	/*------------------------- address route end   -------------------------------------*/
	//the base is dashboard
	//muxRouting.HandleFunc("/", routerDefault)
	//muxRouting.HandleFunc("", routerDefault)

	// check the controller and action
	// handle the Module based routing
	// ca = Controller Action
	// to check only numerals use: {id:[0-9]+}
	//muxRouting.HandleFunc("/ca/{controller}/{action}/", qualifyControllerAndAction)
	// the duplicate url without the trailing /
	//muxRouting.HandleFunc("/ca/{controller}/{action}", qualifyControllerAndAction)

	//url to handle panic errors
	//muxRouting.HandleFunc("/error/error-500/{paniccode}", handle500Page)

	//custom 404 page
	//muxRouting.NotFoundHandler = http.HandlerFunc(handle404Page)

	// the login page
	//muxRouting.HandleFunc("/authenticate/login", routerAuthenticateLogin)

}

/*
// activate if required
func qualifyControllerAndAction(w http.ResponseWriter, r *http.Request) {

	var _routeMaker = new(lib.RouteMaker)
	_routeMaker.MakeRoutes(mux.Vars(r), _handleFilePath)

}
*/

// To handle PANIC in run time
func RecoverWrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			r := recover()
			if r != nil {
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}
				sendMeMail(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func sendMeMail(err error) {
	// send mail
	fmt.Println("Hey! you see it paniced, you want to have a mail? then edit me in mail.go")
	fmt.Println(err)
}

func routerHomeManage(w http.ResponseWriter, r *http.Request) {

	// set the panic handle
	setHandleRedirectAndPanic(w, r)

	// set the session handler
	setHandleSession(w, r)

	var _controllerhomeView = new(controller.ControllerHome)

	_controllerhomeView.ActionManage(mux.Vars(r), r, w, &_handleFilePath, &_handleSession, &_handleRedirectAndPanic, &_envConfigVars)

}
