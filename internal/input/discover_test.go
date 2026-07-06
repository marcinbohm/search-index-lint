package input

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestDiscoverIgnoresDefaultDirectories(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "keep.json", `{}`)
	for _, dir := range []string{".git", "node_modules", "vendor", "dist", "build"} {
		writeFile(t, root, filepath.Join(dir, "ignored.json"), `{}`)
	}

	sources, err := Discover(root)
	if err != nil {
		t.Fatalf("Discover returned error: %v", err)
	}

	if len(sources) != 1 {
		t.Fatalf("Discover returned %d sources, want 1: %#v", len(sources), sources)
	}
	if filepath.Base(sources[0].Path) != "keep.json" {
		t.Fatalf("Discover returned %q, want keep.json", sources[0].Path)
	}
}

func TestDiscoverIgnoresLocalDirectory(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "schema.json", `{}`)
	writeFile(t, root, filepath.Join(".local", "notes.json"), `{bad`)

	sources, err := Discover(root)
	if err != nil {
		t.Fatalf("Discover returned error: %v", err)
	}

	if len(sources) != 1 {
		t.Fatalf("Discover returned %d sources, want 1", len(sources))
	}
	if filepath.Base(sources[0].Path) != "schema.json" {
		t.Fatalf("Discover returned %q, want schema.json", sources[0].Path)
	}
}

func TestDiscoverReturnsDeterministicOrder(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "z.json", `{}`)
	writeFile(t, root, filepath.Join("nested", "a.json"), `{}`)
	writeFile(t, root, "b.json", `{}`)

	sources, err := Discover(root)
	if err != nil {
		t.Fatalf("Discover returned error: %v", err)
	}

	var got []string
	for _, source := range sources {
		got = append(got, filepath.ToSlash(source.RelativePath))
	}

	want := []string{
		filepath.ToSlash(relativePath(filepath.Join(root, "b.json"))),
		filepath.ToSlash(relativePath(filepath.Join(root, "nested", "a.json"))),
		filepath.ToSlash(relativePath(filepath.Join(root, "z.json"))),
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Discover order = %v, want %v", got, want)
	}
}

func TestDiscoverSkipsUnsupportedFiles(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "mapping.json", `{}`)
	writeFile(t, root, "samples.jsonl", `{}`)
	writeFile(t, root, "samples.ndjson", `{}`)
	writeFile(t, root, "README.md", `{bad`)
	writeFile(t, root, "main.go", `{bad`)
	writeFile(t, root, "data.bin", `{bad`)
	writeFile(t, root, "config.yaml", `{bad`)

	sources, err := Discover(root)
	if err != nil {
		t.Fatalf("Discover returned error: %v", err)
	}

	var got []string
	for _, source := range sources {
		got = append(got, filepath.Base(source.Path))
	}

	want := []string{"mapping.json", "samples.jsonl", "samples.ndjson"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Discover returned files %v, want %v", got, want)
	}
}

func TestDiscoverMissingPathReturnsClearError(t *testing.T) {
	_, err := Discover(filepath.Join(t.TempDir(), "missing"))
	if err == nil {
		t.Fatal("Discover returned nil error for missing path")
	}
}

func writeFile(t *testing.T, root, name, content string) {
	t.Helper()
	path := filepath.Join(root, name)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}
}
