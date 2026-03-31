// Package rclone 提供基于 rclone 的统一存储后端实现。
// 按需引入 rclone backend，避免 backend/all 导致二进制膨胀。
package rclone

import (
	_ "github.com/rclone/rclone/backend/drive"
	_ "github.com/rclone/rclone/backend/ftp"
	_ "github.com/rclone/rclone/backend/local"
	_ "github.com/rclone/rclone/backend/s3"
	_ "github.com/rclone/rclone/backend/webdav"
)
