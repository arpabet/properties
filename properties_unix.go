//go:build (linux || openbsd || freebsd || netbsd) && !android && !ci

/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package properties

import (
	"os"
	"path/filepath"
)

// AppDataDir resolves the data directory under the ~/.config base on Linux and
// the BSDs.
func AppDataDir(companyName, appName string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".config", companyName, appName), nil
}
