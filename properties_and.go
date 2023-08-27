// +build !ci
// +build android mobile

/*
 * Copyright (c) 2023 Zander Schwid & Co. LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package properties

import "path/filepath"

func AppDataDir(companyName, appName string) string {
	return filepath.Join("/data", "data", companyName, appName)
}

