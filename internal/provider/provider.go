package provider

type Source string

const (
	SourcePluggy Source = "pluggy"
)

type Ref struct {
	ExternalID string
	Source     Source
}
