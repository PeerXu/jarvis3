package user

type Repository interface {
	CreateUser(*User) (*User, error)
	DeleteUser(username string) error
	GetUser(username string) (*User, error)
	CreateAccessToken(*User, *AccessToken) error
	DeleteAccessTokens(*User, []*AccessToken) error
	LookupUserByAccessToken(*AccessToken) (*User, error)
	CreateAgentToken(*User, *AgentToken) error
	DeleteAgentTokens(*User, []*AgentToken) error
	LookupUserByAgentToken(*AgentToken) (*User, error)
	LookupAgentTokenByName(*User, string) (*AgentToken, error)
}
