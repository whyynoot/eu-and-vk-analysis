package app

import (
	"encoding/json"
	"eu-and-vk-analysis/backend/client_models"
	"eu-and-vk-analysis/backend/server"
	_ "eu-and-vk-analysis/docs"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Server for router, and server handlers

type AnalyticsServer struct {
	analytics *backend.Analytics
}

func NewAnalyticsServer() (*AnalyticsServer, error) {
	analytics, err := backend.NewAnalytics()
	if err != nil {
		return nil, err
	}
	return &AnalyticsServer{analytics: analytics}, nil
}

func (ts *AnalyticsServer) closeDB() {
	ts.analytics.CloseDB()
}

func renderJSON(w http.ResponseWriter, v interface{}, code int) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(js)
}

// Get interests ... Get interests by performance
// @Summary Get interests
// @Description Get interests by performance
// @Tags Interests
// @Param filter path string true "Filter" Enums(bad, good, excellent, three)
// @Success 200 {object} client_models.Response
// @Failure 400,500 {object} client_models.BadResponse
// @Router /interests/{filter} [get]
func (ts *AnalyticsServer) interestsHandler(w http.ResponseWriter, req *http.Request) {
	InputPerformance := mux.Vars(req)["filter"]
	status, err := ts.analytics.CheckCorrectPerformance(InputPerformance)
	if err != nil {
		renderJSON(w, client_models.BadResponse{
			Status: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	response := ts.analytics.AnalyseInterests(status)
	if response.Status != "OK" {
		renderJSON(w, client_models.BadResponse{Status: response.Status}, http.StatusInternalServerError)
	} else {
		renderJSON(w, response, http.StatusOK)
	}
}

// Get students ... Get students by filter
// @Summary Get students by filter
// @Description Get students by filter
// @Description Currently only supporting vk group id
// @Tags Students
// @Param filter path string true "Filter"
// @Success 200 {object} client_models.Response
// @Failure 400,500 {object} client_models.BadResponse
// @Router /students/{filter} [get]
func (ts *AnalyticsServer) studentsHandler(w http.ResponseWriter, req *http.Request) {
	InputGroupId := mux.Vars(req)["filter"]
	GroupId, err := strconv.Atoi(InputGroupId)
	if err != nil {
		log.Println(err)
		renderJSON(w, client_models.BadResponse{
			Status: "Filter Not Supported",
		}, http.StatusBadRequest)
		return
	}

	response := ts.analytics.AnalyseStudents(GroupId)
	if response.Status != "OK" {
		renderJSON(w, client_models.BadResponse{Status: response.Status}, http.StatusInternalServerError)
	} else {
		renderJSON(w, response, http.StatusOK)
	}
}

// App for running, initing server and router

type App struct {
	router         *mux.Router
	server         *http.Server
	analyticsSever *AnalyticsServer
}

type ServerConfig struct {
	Port string `envconfig:"PORT" default:"8000"`
}

func NewApp() *App {
	// Initializing logger, and setting it up
	//file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	//if err != nil {
	//	log.Fatalf("Error setting logs output file %v", err)
	//}
	//log.SetOutput(file)
	log.SetOutput(os.Stdout)

	serverConfig := ServerConfig{}
	err := envconfig.Process("", &serverConfig)
	if err != nil {
		log.Fatalf("Error getting config data %v", err)
	}

	app := new(App)

	// Initializing router
	app.router = mux.NewRouter()
	app.analyticsSever, err = NewAnalyticsServer()
	if err != nil {
		log.Fatalf("Fatal error on initsiliazing analytics server %v", err)
	}

	//Router handling some functions
	app.router.HandleFunc("/interests/{filter}", app.analyticsSever.interestsHandler).Methods("GET")
	app.router.HandleFunc("/students/{filter}", app.analyticsSever.studentsHandler).Methods("GET")

	app.router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	app.router.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./frontend/")))
	app.router.PathPrefix("/images/").Handler(http.FileServer(http.Dir("./frontend/")))
	app.router.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/html/")))

	app.NewServer(serverConfig.Port)

	return app
}

func (app *App) Run() {
	defer app.analyticsSever.closeDB()
	log.Println("Staring server")
	log.Fatal(app.server.ListenAndServe())
}

func (app *App) NewServer(port string) {
	app.server = &http.Server{
		Handler:      app.router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
