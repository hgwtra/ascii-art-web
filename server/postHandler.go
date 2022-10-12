package server

import (
	"ascii/drawing"
	"errors"
	"net/http"
	"strings"
	"fmt"
)

func PostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorBadRequest(w, errors.New("only POST method allowed"))
		return
	}

	r.ParseForm()
	input := r.FormValue("text")
	banner := r.FormValue("fonts")
	fmt.Println(input, banner)
	if banner == "" || input == "" {
		errorBadRequest(w, errors.New("body is wrong, should contain text and banner fields"))
		return
	}

	var finalText string
	input = strings.ReplaceAll(input, "\r", "")
	sourceTextArr := strings.Split(input, "\n")
	for _, v := range sourceTextArr {
		if len(v) != 0 {
			res := drawing.Display(v, banner)
			finalText += res
		} else {
			finalText += "\n"
		}

	}
	if finalText == "" {
		errorInternalServer(w)
		return
	} else {
		Tpl.Execute(w, finalText)
	}

}
