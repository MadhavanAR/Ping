// Copyright (c) 2015-present Ping, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package trial

import (
	"context"
	"net/http"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/v8/channels/store"
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

// Service provides trial limit management functionality following the service layer pattern
type Service struct {
	userStore     store.UserStore
	licenseFn     func() *model.License
	usageService  UsageService
}

// ServiceConfig configures the TrialService
type ServiceConfig struct {
	UserStore    store.UserStore
	LicenseFn    func() *model.License
	UsageService UsageService
}

// UsageService provides usage statistics
type UsageService interface {
	GetPostsUsage(ctx context.Context) (int64, error)
	GetStorageUsage(ctx context.Context) (int64, error)
}

// New creates a new TrialService with dependency injection
func New(c ServiceConfig) (*Service, error) {
	if c.UserStore == nil || c.LicenseFn == nil {
		return nil, model.NewAppError("TrialService.New", "app.trial.service.config_error", nil, "required dependencies missing", http.StatusInternalServerError)
	}

	return &Service{
		userStore:    c.UserStore,
		licenseFn:    c.LicenseFn,
		usageService: c.UsageService,
	}, nil
}

// IsTrialAccount checks if the current license is a trial license
func (s *Service) IsTrialAccount() bool {
	license := s.licenseFn()
	return license != nil && license.IsTrialLicense()
}

// CheckUserLimit checks if the trial user limit has been reached
func (s *Service) CheckUserLimit(ctx context.Context) (bool, *model.AppError) {
	if !s.IsTrialAccount() {
		return false, nil
	}

	activeUserCount, appErr := s.userStore.Count(model.UserCountOptions{})
	if appErr != nil {
		return false, model.NewAppError("CheckUserLimit", "app.trial.check_user_limit.store_error", nil, "", http.StatusInternalServerError).Wrap(appErr)
	}

	if activeUserCount >= TrialLimits.MaxUsers {
		return true, model.NewAppError("CheckUserLimit", "api.user.create_user.trial_user_limit.exceeded", map[string]any{"limit": TrialLimits.MaxUsers}, "", http.StatusBadRequest)
	}

	return false, nil
}

// CheckMessageLimit checks if the trial message limit has been reached
func (s *Service) CheckMessageLimit(ctx context.Context) (bool, *model.AppError) {
	if !s.IsTrialAccount() {
		return false, nil
	}

	if s.usageService == nil {
		return false, nil
	}

	postCount, err := s.usageService.GetPostsUsage(ctx)
	if err != nil {
		return false, model.NewAppError("CheckMessageLimit", "app.trial.check_message_limit.store_error", nil, "", http.StatusInternalServerError).Wrap(err)
	}

	if postCount >= TrialLimits.MaxMessages {
		return true, model.NewAppError("CheckMessageLimit", "api.post.create_post.trial_message_limit.exceeded", map[string]any{"limit": TrialLimits.MaxMessages}, "", http.StatusBadRequest)
	}

	return false, nil
}

// CheckStorageLimit checks if the trial storage limit has been reached
func (s *Service) CheckStorageLimit(ctx context.Context, additionalBytes int64) (bool, *model.AppError) {
	if !s.IsTrialAccount() {
		return false, nil
	}

	if s.usageService == nil {
		return false, nil
	}

	currentStorageBytes, err := s.usageService.GetStorageUsage(ctx)
	if err != nil {
		return false, model.NewAppError("CheckStorageLimit", "app.trial.check_storage_limit.store_error", nil, "", http.StatusInternalServerError).Wrap(err)
	}

	// Convert GB limit to bytes (1 GB = 1024^3 bytes)
	maxStorageBytes := TrialLimits.MaxStorageGB * 1024 * 1024 * 1024

	if currentStorageBytes+additionalBytes > maxStorageBytes {
		return true, model.NewAppError("CheckStorageLimit", "api.file.upload_file.trial_storage_limit.exceeded", map[string]any{"limit": TrialLimits.MaxStorageGB}, "", http.StatusBadRequest)
	}

	return false, nil
}

// GetLimitsInfo returns the current trial limits information
func (s *Service) GetLimitsInfo() map[string]int64 {
	if !s.IsTrialAccount() {
		return nil
	}

	return map[string]int64{
		"maxUsers":     TrialLimits.MaxUsers,
		"maxMessages":  TrialLimits.MaxMessages,
		"maxStorageGB": TrialLimits.MaxStorageGB,
	}
}

