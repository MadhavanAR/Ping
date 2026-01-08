// Copyright (c) 2015-present Ping, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package api4

import (
	"encoding/json"
	"net/http"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/v8/channels/app/health"
)

func (api *API) InitHealth() {
	api.BaseRoutes.Root.Handle("/health", api.APIHandler(getHealth)).Methods("GET")
	api.BaseRoutes.Root.Handle("/health/live", api.APIHandler(getLiveness)).Methods("GET")
	api.BaseRoutes.Root.Handle("/health/ready", api.APIHandler(getReadiness)).Methods("GET")
}

func getHealth(c *Context, w http.ResponseWriter, r *http.Request) {
	healthService, err := health.New(health.ServiceConfig{
		Store: c.App.Srv().Store(),
	})
	if err != nil {
		c.Err = model.NewAppError("getHealth", "api.health.service_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	status, err := healthService.Check(c.AppContext)
	if err != nil {
		c.Err = model.NewAppError("getHealth", "api.health.check_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if status.Status == "unhealthy" {
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	json, err := json.Marshal(status)
	if err != nil {
		c.Err = model.NewAppError("getHealth", "api.marshal_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	w.Write(json)
}

func getLiveness(c *Context, w http.ResponseWriter, r *http.Request) {
	healthService, err := health.New(health.ServiceConfig{
		Store: c.App.Srv().Store(),
	})
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	if err := healthService.Liveness(c.AppContext); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func getReadiness(c *Context, w http.ResponseWriter, r *http.Request) {
	healthService, err := health.New(health.ServiceConfig{
		Store: c.App.Srv().Store(),
	})
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	if err := healthService.Readiness(c.AppContext); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

