package main

import (
	"net/http"
)

type SystemInfo struct {
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	sysInfo := SystemInfo{
		Environment: app.config.env,
		Version:     version,
	}

	status := envelope{
		"status":      "available",
		"system_info": sysInfo,
	}

	err := app.writeJSON(w, http.StatusOK, status, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
