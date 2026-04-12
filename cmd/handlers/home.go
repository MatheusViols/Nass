package handlers


import (
	"net/http"
	"io/ioutil"
)

func Home(res http.ResponseWriter, req *http.Request) {
	htmlFile, err := ioutil.ReadFile("../../templates/index.html")
	if err != nil {
		http.Error(res, "Could not supply home page", http.StatusInternalServerError)
		return		
	}

	res.WriteHeader(http.StatusOK)
	res.Write(htmlFile)
}