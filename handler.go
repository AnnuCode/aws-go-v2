package main

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, data any) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

type envelope map[string]any

type Health struct {
	Status    string
	Container string
	Project   string
}

type Code struct {
	Code_name struct {
		S string `json:"S"`
	} `json:"- code_name"`
	Secret_code struct {
		S string `json:"S"`
	} `json:"- secret_code"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {

	health := Health{
		Status:    "healthy",
		Container: "link to Docker container",
		Project:   "https://github.com/AnnuCode/devops-practice",
	}

	js, err := json.MarshalIndent(health, "", "\t")
	if err != nil {
		http.Error(w, "There was an error", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	//commented part below is an example to test database operation(Get) whose result can then be passed as a JSON http response

	// output, err := tableBasics.GetMovie("First movie", 2001)
	// if err != nil {
	// 	log.Println("There was an error in finding movie")
	// }
	// js, err := json.MarshalIndent(output, "", "\t")
	// if err != nil {
	// 	http.Error(w, "There was an error", http.StatusInternalServerError)
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(js)

}

func secretHandler(w http.ResponseWriter, r *http.Request) {

	secret_code := Code{
		Code_name: struct {
			S string "json:\"S\""
		}{S: "doctor"},
		Secret_code: struct {
			S string "json:\"S\""
		}{S: "ssshhh!!"},
	}

	js, err := json.MarshalIndent(envelope{"- secret_code": secret_code}, "", "\t")
	if err != nil {
		http.Error(w, "There was an error", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
