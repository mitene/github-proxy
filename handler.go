package proxy

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
)

func RepoHandler(w http.ResponseWriter, r *http.Request) {
	archiveDir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatalln(err)
	}
	defer os.RemoveAll(archiveDir)

	ref := "master"
	if r.FormValue("ref") != "" {
		ref = r.FormValue("ref")
	}

	fileType := "tgz"
	if r.FormValue("type") != "" {
		fileType = r.FormValue("type")
	}

	path := "/"
	if r.FormValue("path") != "" {
		path = r.FormValue("path")
	}

	vars := mux.Vars(r)
	log.Printf("Start getting GitHub content: %s/%s\n", vars["owner"], vars["repo"])
	log.Printf("options: ref: %s, path: %s, type: %s\n", ref, path, fileType)

	client := NewGithubClient()
	file, err := client.MakeArchive(vars["owner"], vars["repo"], ref, path, archiveDir, fileType)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusForbidden)
		return
	}
	defer os.Remove(file)
	log.Printf("Successfully retrieved: %s\n", file)

	fileInfo, err := os.Stat(file)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusForbidden)
		return
	}
	fileName := filepath.Base(file[:len(file)])
	fileSize := strconv.Itoa(int(fileInfo.Size()))

	log.Printf("get: file: %s\n", file)
	log.Printf("get: filesize: %s\n", fileSize)

	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	if mimetype := mime.TypeByExtension(filepath.Ext(file)); mimetype != "" {
		w.Header().Add("Content-Type", mimetype)
	} else {
		w.Header().Add("Content-Type", "application/octet-stream")
	}
	w.Header().Add("Content-Length", fileSize)
	http.ServeFile(w, r, file)
}
