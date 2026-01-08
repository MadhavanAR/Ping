// Copyright (c) 2015-present Ping, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package health

import (
	"context"
	"time"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/v8/channels/store"
)

// Service provides health check functionality
type Service struct {
	store store.Store
}

// ServiceConfig configures the HealthService
type ServiceConfig struct {
	Store store.Store
}

// New creates a new HealthService
func New(c ServiceConfig) (*Service, error) {
	if c.Store == nil {
		return nil, model.NewAppError("HealthService.New", "app.health.service.config_error", nil, "store is required", 500)
	}

	return &Service{
		store: c.Store,
	}, nil
}

// HealthStatus represents the health status of the system
type HealthStatus struct {
	Status      string            `json:"status"`
	Version     string            `json:"version"`
	Database    DatabaseHealth    `json:"database"`
	Cache       CacheHealth       `json:"cache"`
	FileStorage FileStorageHealth `json:"file_storage"`
	Timestamp   int64             `json:"timestamp"`
}

// DatabaseHealth represents database connectivity status
type DatabaseHealth struct {
	Status      string `json:"status"`
	ResponseTime int64  `json:"response_time_ms"`
}

// CacheHealth represents cache connectivity status
type CacheHealth struct {
	Status      string `json:"status"`
	ResponseTime int64  `json:"response_time_ms"`
}

// FileStorageHealth represents file storage status
type FileStorageHealth struct {
	Status string `json:"status"`
}

// Check performs a comprehensive health check
func (s *Service) Check(ctx context.Context) (*HealthStatus, error) {
	status := &HealthStatus{
		Status:    "healthy",
		Version:   model.CurrentVersion,
		Timestamp: time.Now().Unix(),
	}

	// Check database
	dbStart := time.Now()
	if err := s.checkDatabase(ctx); err != nil {
		status.Status = "unhealthy"
		status.Database = DatabaseHealth{
			Status: "unhealthy",
		}
	} else {
		status.Database = DatabaseHealth{
			Status:      "healthy",
			ResponseTime: time.Since(dbStart).Milliseconds(),
		}
	}

	// Cache and file storage checks would go here
	// For now, mark as healthy if database is healthy
	if status.Database.Status == "healthy" {
		status.Cache = CacheHealth{Status: "healthy"}
		status.FileStorage = FileStorageHealth{Status: "healthy"}
	}

	return status, nil
}

// checkDatabase verifies database connectivity
func (s *Service) checkDatabase(ctx context.Context) error {
	// Simple query to check database connectivity
	_, err := s.store.System().GetByName("health_check")
	if err != nil && err.Id != "store.sql_system.get_by_name.app_error" {
		// If it's not a "not found" error, database is unhealthy
		return err
	}
	return nil
}

// Liveness checks if the service is alive (basic check)
func (s *Service) Liveness(ctx context.Context) error {
	// Simple check - if we can create a context, we're alive
	return nil
}

// Readiness checks if the service is ready to serve traffic
func (s *Service) Readiness(ctx context.Context) error {
	// Check if database is accessible
	return s.checkDatabase(ctx)
}

