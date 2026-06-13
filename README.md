# properties

A small, dependency-free Go library to locate a per-user, per-company application
data directory across desktop and mobile platforms, with convenience helpers for
loading and saving JSON files.

## Install

```bash
go get go.arpabet.com/properties
```

## Usage

```go
package main

import (
	"fmt"

	"go.arpabet.com/properties"
)

type Settings struct {
	Theme   string `json:"theme"`
	Volume  int    `json:"volume"`
}

func main() {
	data := properties.Locate("AwesomeCompany")

	// Save a JSON file under the app's data directory.
	settings := &Settings{Theme: "dark", Volume: 7}
	if err := data.SaveJsonFile("SuperApp", "settings.json", settings); err != nil {
		panic(err)
	}

	// Load it back. The bool reports whether the file existed.
	var loaded Settings
	found, err := data.LoadJsonFile("SuperApp", "settings.json", &loaded)
	if err != nil {
		panic(err)
	}
	if found {
		fmt.Printf("%+v\n", loaded)
	}

	// Resolve the directory path without writing anything.
	fmt.Println(data.GetDir("SuperApp"))
}
```

## API

```go
func Locate(companyName string) CompanyData
```

```go
type CompanyData interface {
	// MakeDir creates the app data directory (if needed) and returns its path.
	MakeDir(appName string) (string, error)

	// GetDir returns the app data directory path without creating it.
	GetDir(appName string) string

	// LoadJsonFile reads and unmarshals fileName into v.
	// The bool reports whether the file existed; a missing file is not an error.
	LoadJsonFile(appName, fileName string, v interface{}) (bool, error)

	// SaveJsonFile marshals v to JSON and writes it to fileName,
	// creating the app data directory if necessary.
	SaveJsonFile(appName, fileName string, v interface{}) error
}
```

## Resolved locations

The directory is `<base>/<companyName>/<appName>`, where `<base>` is:

| Platform            | Base directory                       |
| ------------------- | ------------------------------------ |
| macOS / iOS         | `~/Library/Application Support`      |
| Windows             | `~/AppData/Roaming`                  |
| Linux / *BSD        | `~/.config`                          |
| Android (`mobile`)  | `/data/data`                         |
| Other / fallback    | `/tmp`                               |

## License

BUSL-1.1. See [LICENSE](LICENSE).
