package new

import (
	"path"
	"text/template"
	"os"
	"log"
	"runtime"
)

// BeOnly entry point.
func BeOnly(name, beType string) {

	server 				:= BackendServer()
	database 			:= Database()
	serverTesting := ServerTesting()
	

	if name != "" && server != "" && database != "" && serverTesting != "" {
		create(name, server, database, serverTesting, beType)
	} else {
		return 
	}
}

func create(name, server, database, testing, beType string) {
	// basically checking if this is being used to create a backend for another project type (SPA, Multi, or MVC)
	if beType != "html" {
		Mkdir(name)
		CreateFile("../files/common/.gitignore", name + "/.gitignore")
		CreateFile("../files/common/README.md", name + "/README.md")
	}
	Mkdir(name + "/server")
	Mkdir(name + "/server/controllers")

	createServer(name, server, beType)
	createDatabase(name, database)
	// make testing folder go get the testing lib 


	// also ask about creating docker file 

}

func createServer(name, server, beType string) {
	switch server {
	case "Chi":
		chi(name, beType)
	case "HTTP Router":
		http(name, beType)
	case "Gin (recommended)":
		gin(name, beType)
	case "Echo":
		echo(name, beType)
	case "Gorilla mux":
		gorilla(name, beType)
	}
}

func chi(name, beType string) {
	CreateFile("../files/chi/.chi.go", name + "/server/server.go")
	CreateFile("../files/chi/.routes.go", name + "/routes.go")

}

func http(name, beType string) {
	CreateFile("../files/http/.http_router.go", name + "/server/server.go")
	CreateFile("../files/http/.routes.go", name + "/server/routes.go")
}

func gin(name, beType string) {
	CreateFile("../files/gin/.gin.go", name + "/server/server.go")
	// CreateFile("../files/gin/.routes.go", name + "/server/routes.go")
	importControllerPath("../files/gin/routes.tmpl", name + "/server/routes.go", name + "/server/controllers")
	
	if beType == "html" {

	} else {
		CreateFile("../files/gin/.jsonController.go", name + "/server/controllers/hello.go")
	}
}

func echo(name, beType string) {
	CreateFile("../files/echo/.echo.go", name + "/server/server.go")
	importControllerPath("../files/echo/routes.tmpl", name + "/server/routes.go", name + "/server/controllers")
	if beType == "html" {

	} else {
		CreateFile("../files/echo/.jsonController.go", name + "/server/controllers/hello.go")
	}
}

func gorilla(name, beType string) {
	CreateFile("../files/gorilla/.gorilla.go", name + "/server/server.go")
	importControllerPath("../files/gorilla/routes.tmpl", name + "/server/routes.go", name + "/server/controllers")
	if beType == "html" {
		// html controller
	} else {
		// json controller 
		CreateFile("../files/gorilla/.jsonController.go", name + "/server/controllers/hello.go")
	}
}

func createDatabase(name, database string) {
	if database == "Postgres" {
		Mkdir(name + "/server/models")
	} else if database == "MongoDB" {
		Mkdir(name + "/server/models")
	} else {
		return 
	}	
}

type Route struct {
	Path string 
}

func importControllerPath(templatePath, createFilePath, folderPath string) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
			panic("No caller information")
	}

	tmpl, _ := template.ParseFiles(path.Join(filename, templatePath))
	path := GetPath(folderPath)
	route := Route{
		Path: path,
	}
	file, err := os.Create(createFilePath)
	if err != nil {
		log.Println("Create file: ", err)
	}

	error := tmpl.Execute(file, route)
	if error != nil {
		log.Fatalf("template execution: %s", error)
	}
}