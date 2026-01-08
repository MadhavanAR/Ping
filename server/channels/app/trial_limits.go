// Copyright (c) 2015-present Ping, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package app

import (
	"net/http"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/shared/request"
)

// TrialLimits defines the limits for free trial accounts
var TrialLimits = struct {
	MaxUsers     int64
	MaxMessages  int64
	MaxStorageGB int64
}{
	MaxUsers:     10,
	MaxMessages:  10000,
	MaxStorageGB: 10,
}

// IsTrialAccount checks if the current license is a trial license
func (a *App) IsTrialAccount() bool {
	license := a.License()
	return license != nil && license.IsTrialLicense()
}

// CheckTrialUserLimit checks if the trial user limit has been reached
func (a *App) CheckTrialUserLimit(rctx request.CTX) (bool, *model.AppError) {
	if !a.IsTrialAccount() {
		return false, nil
	}

	activeUserCount, appErr := a.Srv().Store().User().Count(model.UserCountOptions{})
	if appErr != nil {
		return false, model.NewAppError("CheckTrialUserLimit", "app.limits.check_trial_user_limit.store_error", nil, "", http.StatusInternalServerError).Wrap(appErr)
	}

	if activeUserCount >= TrialLimits.MaxUsers {
		return true, model.NewAppError("CheckTrialUserLimit", "api.user.create_user.trial_user_limit.exceeded", map[string]any{"limit": TrialLimits.MaxUsers}, "", http.StatusBadRequest)
	}

	return false, nil
}

// CheckTrialMessageLimit checks if the trial message limit has been reached
func (a *App) CheckTrialMessageLimit(rctx request.CTX) (bool, *model.AppError) {
	if !a.IsTrialAccount() {
		return false, nil
	}

	// Get total post count for the workspace
	postCount, appErr := a.GetPostsUsage()
	if appErr != nil {
		return false, model.NewAppError("CheckTrialMessageLimit", "app.limits.check_trial_message_limit.store_error", nil, "", http.StatusInternalServerError).Wrap(appErr)
	}

	if postCount >= TrialLimits.MaxMessages {
		return true, model.NewAppError("CheckTrialMessageLimit", "api.post.create_post.trial_message_limit.exceeded", map[string]any{"limit": TrialLimits.MaxMessages}, "", http.StatusBadRequest)
	}

	return false, nil
}

// CheckTrialStorageLimit checks if the trial storage limit has been reached
func (a *App) CheckTrialStorageLimit(rctx request.CTX, additionalBytes int64) (bool, *model.AppError) {
	if !a.IsTrialAccount() {
		return false, nil
	}

	// Get current storage usage in bytes
	currentStorageBytes, appErr := a.GetStorageUsage()
	if appErr != nil {
		return false, model.NewAppError("CheckTrialStorageLimit", "app.limits.check_trial_storage_limit.store_error", nil, "", http.StatusInternalServerError).Wrap(appErr)
	}

	// Convert GB limit to bytes (1 GB = 1024^3 bytes)
	maxStorageBytes := TrialLimits.MaxStorageGB * 1024 * 1024 * 1024

	if currentStorageBytes+additionalBytes > maxStorageBytes {
		return true, model.NewAppError("CheckTrialStorageLimit", "api.file.upload_file.trial_storage_limit.exceeded", map[string]any{"limit": TrialLimits.MaxStorageGB}, "", http.StatusBadRequest)
	}

	return false, nil
}

// GetTrialLimits returns the current trial limits information
// Note: This is a helper function. Actual limits are enforced through LicenseLimits
func (a *App) GetTrialLimitsInfo() map[string]int64 {
	if !a.IsTrialAccount() {
		return nil
	}

	return map[string]int64{
		"maxUsers":     TrialLimits.MaxUsers,
		"maxMessages":  TrialLimits.MaxMessages,
		"maxStorageGB": TrialLimits.MaxStorageGB,
	}
}

