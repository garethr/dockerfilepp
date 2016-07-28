package dockerfilepp

import (
	"testing"
)

func TestProcessRuns(t *testing.T) {
	replacements := make(map[string]string)
	Process(replacements, "Help")
}

func TestRenderPlain(t *testing.T) {
	output := render("plain", "")
	if string(output) != "plain" {
		t.Error("render not passing through content")
	}
}

func TestRenderUsesTemplate(t *testing.T) {
	output := render("COPY {{if .Value}}{{.Value}}{{else}}manifests{{end}} /manifests", "")
	if string(output) != "COPY manifests /manifests" {
		t.Error("render is not parsing Go templates correctly")
	}
}

func TestRenderUsesTemplateValues(t *testing.T) {
	output := render("COPY {{if .Value}}{{.Value}}{{else}}manifests{{end}} /manifests", "something")
	if string(output) != "COPY something /manifests" {
		t.Error("render is not substituting values into templates")
	}
}
