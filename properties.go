/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

// Package properties locates a per-user, per-company application data directory
// across desktop and mobile platforms and provides convenience helpers for
// loading and saving JSON files within it. The resolved directory is
// <base>/<companyName>/<appName>, where <base> depends on the operating system
// (see AppDataDir).
package properties

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// CompanyData resolves and manages the data directory for a single company,
// scoped to individual applications by appName.
type CompanyData interface {

	// MakeDir creates the app data directory (if needed) and returns its path.
	MakeDir(appName string) (string, error)

	// GetDir returns the app data directory path without creating it.
	GetDir(appName string) string

	// LoadJsonFile reads and unmarshals fileName into v. The bool reports
	// whether the file existed; a missing file is not an error.
	LoadJsonFile(appName, fileName string, v any) (bool, error)

	// SaveJsonFile marshals v to JSON and writes it to fileName, creating the
	// app data directory if necessary.
	SaveJsonFile(appName, fileName string, v any) error
}

type implCompanyData struct {
	companyName string
}

// Locate returns a CompanyData rooted at the given company name. It performs no
// I/O; the directory is resolved lazily by the returned value's methods.
func Locate(companyName string) CompanyData {
	return &implCompanyData{companyName: companyName}
}

func (t *implCompanyData) MakeDir(appName string) (string, error) {
	dir, err := AppDataDir(t.companyName, appName)
	if err != nil {
		return "", err
	}
	err = os.MkdirAll(dir, 0700)
	return dir, err
}

func (t *implCompanyData) GetDir(appName string) string {
	dir, _ := AppDataDir(t.companyName, appName)
	return dir
}

func (t *implCompanyData) LoadJsonFile(appName, fileName string, v any) (bool, error) {
	dir, err := AppDataDir(t.companyName, appName)
	if err != nil {
		return false, err
	}
	fullPath := filepath.Join(dir, fileName)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return false, nil
	}
	blob, err := os.ReadFile(fullPath)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal(blob, v)
	return err == nil, err
}

func (t *implCompanyData) SaveJsonFile(appName, fileName string, v any) error {
	blob, err := json.Marshal(v)
	if err != nil {
		return err
	}
	dir, err := t.MakeDir(appName)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, fileName), blob, 0660)
}
