package server

import (
	"ascii/drawing"
	"errors"
	"net/http"
	"strings"
)

var FinalText string

type AsciiData struct {
	DrawingText string
	Submit      bool
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	FinalText = ""
	if r.Method != http.MethodPost {
		errorBadRequest(w, errors.New("only POST method allowed"))
		return
	}

	r.ParseForm()
	input := r.FormValue("text")
	banner := r.FormValue("fonts")
	if banner == "" || input == "" {
		errorBadRequest(w, errors.New("body is wrong, should contain text and banner fields"))
		return
	}

	input = strings.ReplaceAll(input, "\r", "")
	sourceTextArr := strings.Split(input, "\n")
	for _, v := range sourceTextArr {
		if len(v) != 0 {
			res := drawing.Display(v, banner)
			FinalText += res
		} else {
			FinalText += "\n"
		}

	}
	if FinalText == "" {
		errorInternalServer(w)
		return
	} else {
		data := AsciiData{
			DrawingText: FinalText,
			Submit:      true,
		}
		Tpl.Execute(w, data)
	}
}
