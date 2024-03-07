package router

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func WebRouter(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL.Path)

	resourcePath := getResourcePath(r.URL.Path)

	index, err := os.ReadFile(resourcePath)

	//return html
	if err != nil {
		handleNotFoundPage(w)
	}

	w.Write(index)

}

func getResourcePath(path string) string {
	//check if path has direct file route like "index.html" at last child

	splitted := strings.Split(path, "/")

	lastChild := splitted[len(splitted)-1]

	//has full path to file  -> .html
	if strings.Contains(lastChild, ".") {
		return strings.Join([]string{SERVE_FOLDER_PATH, path}, "")
	}

	//return default index.html
	return strings.Join([]string{SERVE_FOLDER_PATH, path, "/index.html"}, "")
}

func handleNotFoundPage(w http.ResponseWriter) {
	errPagePAth := strings.Join([]string{SERVE_FOLDER_PATH, PAGE_404}, "/")
	errPage, err := os.ReadFile(errPagePAth)

	if err != nil {
		fmt.Println("Fatal, couldn't get 404 page")
		return
	}

	fmt.Println("Couldn't find resource, serving 404 page")

	w.Write(errPage)
}
