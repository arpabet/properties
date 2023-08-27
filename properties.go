/*
 * Copyright (c) 2023 Zander Schwid & Co. LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package properties

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type CompanyData interface {

	MakeDir(appName string) (string, error)

	GetDir(appName string) string

	LoadJsonFile(appName, fileName string, v interface{}) (bool, error)

	SaveJsonFile(appName, fileName string, v interface{}) error

}

type implCompanyData struct {
	companyName string
}

func Locate(companyName string) CompanyData {
	return &implCompanyData{companyName: companyName}
}

func (t *implCompanyData) MakeDir(appName string) (string, error) {
	dir := AppDataDir(t.companyName, appName)
	err := os.MkdirAll(dir, 0700)
	return dir, err
}

func (t *implCompanyData) GetDir(appName string) string {
	return AppDataDir(t.companyName, appName)
}

func (t *implCompanyData) LoadJsonFile(appName, fileName string, v interface{}) (bool, error) {
	dir := AppDataDir(t.companyName, appName)
	fullPath := filepath.Join(dir, fileName)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return false, nil
	}
	blob, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal(blob, v)
	return err == nil, err
}

func (t *implCompanyData) SaveJsonFile(appName, fileName string, v interface{}) error {
	blob, err := json.Marshal(v)
	if err != nil {
		return err
	}
	dir, err := t.MakeDir(appName)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join(dir, fileName), blob, 0660)
}
