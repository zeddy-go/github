package payload

type Error struct {
	Message          string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
}
