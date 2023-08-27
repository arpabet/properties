// +build !ci
// +build darwin,ios

/*
 * Copyright (c) 2023 Zander Schwid & Co. LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package properties

import (
	"os"
	"path/filepath"
)

func AppDataDir(companyName, appName string) string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, "Library", "Preferences", companyName, appName)
}
