// Copyright (c) 2015-present Ping, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package trial

import (
	"context"
	"net/http"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/v8/channels/utils"
)

// UsageAdapter adapts App usage methods to UsageService interface
type UsageAdapter struct {
	postStore    PostStore
	fileInfoStore FileInfoStore
}

// PostStore interface for post operations
type PostStore interface {
	AnalyticsPostCount(options *model.PostCountOptions) (int64, error)
}

// FileInfoStore interface for file operations
type FileInfoStore interface {
	GetStorageUsage(allowFromCache, includeDeleted bool) (int64, error)
}

// NewUsageAdapter creates a new usage adapter
func NewUsageAdapter(postStore PostStore, fileInfoStore FileInfoStore) *UsageAdapter {
	return &UsageAdapter{
		postStore:     postStore,
		fileInfoStore: fileInfoStore,
	}
}

// GetPostsUsage returns the total posts count
func (u *UsageAdapter) GetPostsUsage(ctx context.Context) (int64, error) {
	count, err := u.postStore.AnalyticsPostCount(&model.PostCountOptions{
		ExcludeDeleted: true,
		UsersPostsOnly: true,
		AllowFromCache: true,
	})
	if err != nil {
		return 0, model.NewAppError("GetPostsUsage", "app.post.analytics_posts_count.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
	}
	return utils.RoundOffToZeroesResolution(float64(count), 3), nil
}

// GetStorageUsage returns the sum of files' sizes
func (u *UsageAdapter) GetStorageUsage(ctx context.Context) (int64, error) {
	usage, err := u.fileInfoStore.GetStorageUsage(true, false)
	if err != nil {
		return 0, model.NewAppError("GetStorageUsage", "app.usage.get_storage_usage.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
	}
	return usage, nil
}

