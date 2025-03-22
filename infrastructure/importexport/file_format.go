package importexport

// FileFormat представляет формат файла для импорта/экспорта
type FileFormat string

const (
	// CSV формат CSV
	CSV FileFormat = "csv"
	// JSON формат JSON
	JSON FileFormat = "json"
	// YAML формат YAML
	YAML FileFormat = "yaml"
)
