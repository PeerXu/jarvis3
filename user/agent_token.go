package user

type AgentToken struct {
	Name  string
	Token string
}

func NewAgentToken(name string) *AgentToken {
	return &AgentToken{
		Name:  name,
		Token: genTokenString(),
	}
}
