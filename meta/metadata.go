package meta

type Metadata struct {
	Package  string
	Name     string
	Version  string
	Revision string
}

func NewMetadata(version string, revision string) *Metadata {
	return &Metadata{
		Package:  "sarif-converter",
		Name:     "SARIF Converter",
		Version:  version,
		Revision: revision,
	}
}
