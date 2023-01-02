// Package version provides version information.
package version

import (
	"bytes"
	"fmt"
	"runtime"
	"text/template"

	"github.com/sentinelos/tasker/pkg/constants"
)

var (
	// Tag is set at build time.
	Tag string
	// SHA is set at build time.
	SHA string
)

const versionTemplate = constants.AppName + `:
 Tag:        {{ .Tag }}
 SHA:        {{ .SHA }}
 Go version: {{ .GoVersion }}
 OS/Arch:    {{ .Os }}/{{ .Arch }}`

// Version contains verbose version information.
type Version struct {
	Tag       string
	SHA       string
	GoVersion string
	Os        string
	Arch      string
}

// PrintLongVersion prints verbose version information.
func PrintLongVersion() {
	v := Version{
		Tag:       Tag,
		SHA:       SHA,
		GoVersion: runtime.Version(),
		Os:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}

	var wr bytes.Buffer

	tmpl, err := template.New("version").Parse(versionTemplate)
	if err != nil {
		fmt.Println(err)
	}

	err = tmpl.Execute(&wr, v)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(wr.String())
}

// PrintShortVersion prints the tag and sha.
func PrintShortVersion() {
	fmt.Printf("%s %s-%s\n", constants.AppName, Tag, SHA)
}
