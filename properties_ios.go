//go:build !ci && darwin && ios

/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package properties

import (
	"os"
	"path/filepath"
)

// AppDataDir resolves the data directory under the iOS
// ~/Library/Application Support base.
func AppDataDir(companyName, appName string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, "Library", "Application Support", companyName, appName), nil
}
