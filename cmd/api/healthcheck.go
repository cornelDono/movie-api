package main

import (
	"net/http"
)

type Status struct {
	Status      string `json:"name"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	status := Status{
		Status:      "available",
		Environment: app.config.env,
		Version:     version,
	}

	app.writeJSON(w, http.StatusOK, status, nil)
}
