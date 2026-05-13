package main

import (
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("Internal Error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJsonError(w, http.StatusInternalServerError, "The server encountered a  problem")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw("Bad Request", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJsonError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw("Not Found Error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJsonError(w, http.StatusNotFound, "Not found")
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("Conflict Response", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJsonError(w, http.StatusConflict, "Conflict error")
}

func (app *application) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("Unauthorized Error Response", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJsonError(w, http.StatusUnauthorized, "Unauthorized error")
}

func (app *application) unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("Unauthorized Basic Error Response", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	w.Header().Set("www-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	writeJsonError(w, http.StatusUnauthorized, "Unauthorized basic error")
}

func (app *application) forbiddenErrorResponse(w http.ResponseWriter, r *http.Request) {
	app.logger.Errorw("Forbidden Error Response", "method", r.Method, "path", r.URL.Path)

	writeJsonError(w, http.StatusForbidden, "Forbidden error")
}
