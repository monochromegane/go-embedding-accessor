package accessor

import (
	"fmt"
	"path/filepath"
)

func Generate(pkg, outputDir, name, files string) error {
	if files == "" {
		return fmt.Errorf("Empty go:embed directive. files is required.")
	}

	return writeToFile(
		filepath.Join(outputDir, fmt.Sprintf("%s_embedding.go", name)),
		embeddingTemplates.toString(),
		data{
			Package: pkg,
			Name:    name,
			Files:   files,
		},
	)
}
