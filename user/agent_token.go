package user

import (
	"fmt"
	"strings"
)

type AgentToken struct {
	Name  string
	Token string
}

func (t *AgentToken) String() string {
	return fmt.Sprintf("{Name: %v, Token: %v}", t.Name, t.Token)
}

func NewAgentToken(name string) *AgentToken {
	return &AgentToken{
		Name:  name,
		Token: genTokenString(),
	}
}

func ParseAgentTokenFromString(s string) (*AgentToken, error) {
	xs := strings.Split(s, ":")
	if len(xs) != 2 {
		return nil, ErrAgentTokenNotFound
	}

	typ, token := xs[0], xs[1]

	if strings.ToLower(strings.Trim(typ, " ")) != "agt" {
		return nil, ErrAgentTokenNotFound
	}

	return &AgentToken{Token: strings.Trim(token, " ")}, nil
}
