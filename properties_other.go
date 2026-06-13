//go:build ci || (!linux && !darwin && !windows && !freebsd && !openbsd && !netbsd)

/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package properties

import "path/filepath"

func AppDataDir(companyName, appName string) (string, error) {
	return filepath.Join("/tmp", companyName, appName), nil
}
