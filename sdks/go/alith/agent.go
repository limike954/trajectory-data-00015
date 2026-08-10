package alith

// Agent stores in-memory configuration for an Alith agent.
type Agent struct {
	ID       string
	Name     string
	Model    string
	APIKey   string
	BaseURL  string
	Preamble string
}

// NewAgent creates an Agent with the provided name and model.
func NewAgent(name, model string) *Agent {
	return &Agent{
		Name:  name,
		Model: model,
	}
}

// WithCredentials sets API credentials on the same Agent and returns it.
func (a *Agent) WithCredentials(apiKey, baseURL string) *Agent {
	a.APIKey = apiKey
	a.BaseURL = baseURL
	return a
}

// WithPreamble sets the Agent preamble on the same Agent and returns it.
func (a *Agent) WithPreamble(preamble string) *Agent {
	a.Preamble = preamble
	return a
}
