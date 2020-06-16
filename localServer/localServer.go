package localServer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"webServerForWebProject/lock"
	"webServerForWebProject/mimestypes"
	"webServerForWebProject/soldierFileJson"
)

const (
	cleaningProjectProgramLoc = "\\build\\dist"
)

//hendle the general pages the client request like css and js file.
func hendleGeneral(w http.ResponseWriter, path string /* contentType string */) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("error, %v", err)
		http.NotFound(w, nil)
		return
	}
	contentType, err := mimestypes.GetFileContentType(path)
	if err != nil {
		log.Printf("error, %v", err)
		http.NotFound(w, nil)
		return
	}
	w.Header().Add("Content-Type", contentType)
	w.Write(f)
}

//the main hendler of the http server
func hendler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var path string = ".\\build"
		if r.URL.Path == "/" {
			path += "\\index.html"
		} else {
			replacer := strings.NewReplacer("/", "\\")
			//objectName b:= helperFunctions.Strcutter(r.URL.Path, '/', true)
			objectName := replacer.Replace(r.URL.Path)
			path += objectName
		}
		hendleGeneral(w, path)
		break
	case http.MethodPost:
		hendlePostRequest(w, r)
	}
}

//hendlePostRequest hendle the post method in this server.
func hendlePostRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	jsonFile, err := soldierFileJson.CreateTheJSON(r.Form)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "write valid pogram")
		return
	}
	file, err := json.MarshalIndent(jsonFile, "", " ")
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "write valid pogram")
		return
	}
	err = ioutil.WriteFile("."+cleaningProjectProgramLoc+"\\soldier_file.json", file, 0644)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "write valid pogram")
		return
	}
	a := startThePogram()
	if a != lock.OK {
		fmt.Fprint(w, "write valid pogram")
		return
	}
	fmt.Fprint(w, "The excel is ready!")
	excelFolder := dir + cleaningProjectProgramLoc + "\\excelFolder"
	cmd := exec.Command("explorer", excelFolder)
	cmd.Dir = excelFolder
	err = cmd.Run()
}

func startThePogram() string {
	lock.Lock.Lock()
	defer lock.Lock.Unlock()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	lock.StartLock(dir + cleaningProjectProgramLoc + "\\.LOCK")
	cmd := exec.Command(dir + cleaningProjectProgramLoc + "\\CleaningProject.exe")
	cmd.Dir = dir + cleaningProjectProgramLoc
	err = cmd.Start()
	if err != nil {
		log.Println(err)
	}
	state := <-lock.C
	return state
}

//open the browser with the current url
func openbrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

//StartLocalServer start the local server of the pogram.
func StartLocalServer() {
	openbrowser("http://localhost:8080/")
	http.HandleFunc("/", hendler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
