/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package properties

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type CompanyData interface {
	MakeDir(appName string) (string, error)

	GetDir(appName string) string

	LoadJsonFile(appName, fileName string, v any) (bool, error)

	SaveJsonFile(appName, fileName string, v any) error
}

type implCompanyData struct {
	companyName string
}

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
