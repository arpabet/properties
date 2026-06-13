/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package properties_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"go.arpabet.com/properties"
)

const appName = "TestApp"

// uniqueCompany returns a company name that is extremely unlikely to collide
// with a real application's data directory, so tests never touch real data.
func uniqueCompany(t *testing.T) string {
	t.Helper()
	return fmt.Sprintf("properties-test-%d-%d", os.Getpid(), time.Now().UnixNano())
}

// cleanup removes the company directory created by a test. The resolved
// directory is <base>/<company>/<app>, so its parent is the company root.
func cleanup(t *testing.T, data properties.CompanyData) {
	t.Helper()
	companyDir := filepath.Dir(data.GetDir(appName))
	if err := os.RemoveAll(companyDir); err != nil {
		t.Logf("cleanup failed for %q: %v", companyDir, err)
	}
}

func TestLocate(t *testing.T) {
	if data := properties.Locate("AwesomeCompany"); data == nil {
		t.Fatal("Locate returned nil")
	}
}

func TestGetDir(t *testing.T) {
	company := uniqueCompany(t)
	data := properties.Locate(company)
	defer cleanup(t, data)

	dir := data.GetDir(appName)
	if dir == "" {
		t.Fatal("GetDir returned an empty path")
	}
	if !filepath.IsAbs(dir) {
		t.Errorf("GetDir = %q, want an absolute path", dir)
	}
	if got := filepath.Base(dir); got != appName {
		t.Errorf("GetDir base = %q, want %q", got, appName)
	}
	if got := filepath.Base(filepath.Dir(dir)); got != company {
		t.Errorf("GetDir company segment = %q, want %q", got, company)
	}

	// GetDir must not create anything on disk.
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		t.Errorf("GetDir created the directory %q (stat err = %v); it should be inert", dir, err)
	}
}

func TestMakeDir(t *testing.T) {
	data := properties.Locate(uniqueCompany(t))
	defer cleanup(t, data)

	dir, err := data.MakeDir(appName)
	if err != nil {
		t.Fatalf("MakeDir: %v", err)
	}
	if dir != data.GetDir(appName) {
		t.Errorf("MakeDir = %q, GetDir = %q; want equal", dir, data.GetDir(appName))
	}

	info, err := os.Stat(dir)
	if err != nil {
		t.Fatalf("stat after MakeDir: %v", err)
	}
	if !info.IsDir() {
		t.Errorf("%q is not a directory", dir)
	}

	// MakeDir must be idempotent.
	if _, err := data.MakeDir(appName); err != nil {
		t.Errorf("second MakeDir: %v", err)
	}
}

type settings struct {
	Theme  string `json:"theme"`
	Volume int    `json:"volume"`
}

func TestSaveAndLoadJsonFile(t *testing.T) {
	data := properties.Locate(uniqueCompany(t))
	defer cleanup(t, data)

	const fileName = "settings.json"
	want := &settings{Theme: "dark", Volume: 7}

	// SaveJsonFile should create the directory tree as needed.
	if err := data.SaveJsonFile(appName, fileName, want); err != nil {
		t.Fatalf("SaveJsonFile: %v", err)
	}

	var got settings
	found, err := data.LoadJsonFile(appName, fileName, &got)
	if err != nil {
		t.Fatalf("LoadJsonFile: %v", err)
	}
	if !found {
		t.Fatal("LoadJsonFile reported the saved file as not found")
	}
	if got != *want {
		t.Errorf("round trip = %+v, want %+v", got, *want)
	}
}

func TestLoadJsonFileMissing(t *testing.T) {
	data := properties.Locate(uniqueCompany(t))
	defer cleanup(t, data)

	var got settings
	found, err := data.LoadJsonFile(appName, "does-not-exist.json", &got)
	if err != nil {
		t.Fatalf("LoadJsonFile on missing file returned error: %v", err)
	}
	if found {
		t.Error("LoadJsonFile reported a missing file as found")
	}
}

func TestLoadJsonFileInvalid(t *testing.T) {
	data := properties.Locate(uniqueCompany(t))
	defer cleanup(t, data)

	dir, err := data.MakeDir(appName)
	if err != nil {
		t.Fatalf("MakeDir: %v", err)
	}
	const fileName = "broken.json"
	if err := os.WriteFile(filepath.Join(dir, fileName), []byte("{ not json"), 0660); err != nil {
		t.Fatalf("seeding invalid file: %v", err)
	}

	var got settings
	found, err := data.LoadJsonFile(appName, fileName, &got)
	if err == nil {
		t.Error("LoadJsonFile on invalid JSON returned nil error")
	}
	if found {
		t.Error("LoadJsonFile reported success on invalid JSON")
	}
}
