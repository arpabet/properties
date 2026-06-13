//go:build !ci && (android || mobile)

/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package properties

import "path/filepath"

// AppDataDir resolves the data directory under the Android /data/data base.
func AppDataDir(companyName, appName string) (string, error) {
	return filepath.Join("/data", "data", companyName, appName), nil
}
