// +build ci !linux,!darwin,!windows,!freebsd,!openbsd,!netbsd

/*
 * Copyright (c) 2023 Zander Schwid & Co. LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package properties

func AppDataDir(companyName, appName string) string {
	return filepath.Join("/tmp", companyName, appName)
}
