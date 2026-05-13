package api

import (
	"fmt"
	"io"

	"github.com/danielgtaylor/huma/v2"
)

func WriteOpenAPI(w io.Writer, api huma.API) error {
	spec, err := api.OpenAPI().DowngradeYAML()
	if err != nil {
		return fmt.Errorf("downgrade openapi spec: %w", err)
	}

	if _, err := w.Write(spec); err != nil {
		return fmt.Errorf("write openapi spec: %w", err)
	}

	return nil
}
