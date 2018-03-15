package new

import (
	// "fmt"
	"os"
	"io/ioutil"
	// "path/filepath"
	// "io"
	"path"
	"log"
	"runtime"
	"strings"
	"github.com/spf13/afero"	
)


// CreateFile takes file path of a local file and copies it to the output path 
func CreateFile(fileSrc, output string) {

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
			panic("No caller information")
	}


	file, err := ioutil.ReadFile(path.Join(filename, fileSrc))
	if err != nil {
		log.Fatal(err)
	}


	appFS := afero.NewOsFs()
	error := afero.WriteFile(appFS, output, file, 0644)
	if error != nil {
		log.Fatal(error)
	}
}

//Mkdir helper function to make a directory  
func Mkdir(name string) {
	path := name
	err := os.MkdirAll(path, 0700)

	if err != nil {
		log.Println("Error creating directory")
		log.Println(err)
		os.Exit(1)
	}
}

// GetPath returns users github.com/user/ path
func GetPath() string {
  dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	path := ""
	if strings.Contains(dir, "github.com") {
		path = strings.Split(dir, "src/")[1]	
	} else {
		path = os.Getenv("GOPATH")
	}
	return path
}
	
