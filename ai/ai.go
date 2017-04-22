package ai

type Request struct {
	Id        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Lang      string `json:"lang"`
	Result    Result `json:"result"`
}

type Result struct {
	Action           string            `json:"action"`
	ActionIncomplete bool              `json:"actionIncomplete"`
	Contexts         []Context         `json:"contexts"`
	Parameters       map[string]string `json:"parameters"`
	ResolvedQuery    string            `json:"resolvedQuery"`
	Source           string            `json:"source"`
}

type Context struct {
	Lifespan   int64             `json:"lifespan"`
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters"`
}

type Response struct {
	ContextOut  []Context         `json:"contextOut"`
	Data        map[string]string `json:"data"`
	DisplayText string            `json:"displayText"`
	Speech      string            `json:"speech"`
	Source      string            `json:"source"`
}

type Event struct {
	Name string            `json:"name"`
	Data map[string]string `json:"data"`
}
