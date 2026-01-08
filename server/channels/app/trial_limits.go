// Copyright (c) 2015-present Ping, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package app

import (
	"context"
	"net/http"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/shared/request"
	"github.com/mattermost/mattermost/server/v8/channels/app/trial"
)

// TrialLimits defines the limits for free trial accounts
// This is kept for backward compatibility
var TrialLimits = trial.TrialLimits

// IsTrialAccount checks if the current license is a trial license
func (a *App) IsTrialAccount() bool {
	license := a.License()
	return license != nil && license.IsTrialLicense()
}

// CheckTrialUserLimit checks if the trial user limit has been reached
func (a *App) CheckTrialUserLimit(rctx request.CTX) (bool, *model.AppError) {
	if a.Srv().trialService == nil {
		return false, nil
	}
	return a.Srv().trialService.CheckUserLimit(rctx.Context())
}

// CheckTrialMessageLimit checks if the trial message limit has been reached
func (a *App) CheckTrialMessageLimit(rctx request.CTX) (bool, *model.AppError) {
	if a.Srv().trialService == nil {
		return false, nil
	}
	return a.Srv().trialService.CheckMessageLimit(rctx.Context())
}

// CheckTrialStorageLimit checks if the trial storage limit has been reached
func (a *App) CheckTrialStorageLimit(rctx request.CTX, additionalBytes int64) (bool, *model.AppError) {
	if a.Srv().trialService == nil {
		return false, nil
	}
	return a.Srv().trialService.CheckStorageLimit(rctx.Context(), additionalBytes)
}

// GetTrialLimitsInfo returns the current trial limits information
func (a *App) GetTrialLimitsInfo() map[string]int64 {
	if a.Srv().trialService == nil {
		return nil
	}
	return a.Srv().trialService.GetLimitsInfo()
}

