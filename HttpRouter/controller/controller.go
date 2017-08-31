package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetTaskProcess(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("requestid"))
}

func PostTaskProcess(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "hello, something wrong happened for sending message to %s!\n", ps.ByName("requestid"))
	} else {
		fmt.Fprint(w, "Your enter is %s\n", []byte(body))
	}
}
