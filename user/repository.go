package user

type Repository interface {
	CreateUser(*User) (*User, error)
	DeleteUserByID(id UserID) error
	GetUserByID(id UserID) (*User, error)
	CreateAccessToken(UserID, *AccessToken) error
	DeleteAccessTokens(UserID, []*AccessToken) error
	LookupUserByUsername(username string) (*User, error)
	LookupUserByAccessToken(*AccessToken) (*User, error)
	CreateAgentToken(UserID, *AgentToken) error
	DeleteAgentTokens(UserID, []*AgentToken) error
	LookupUserByAgentToken(*AgentToken) (*User, error)
	LookupAgentTokenByName(UserID, string) (*AgentToken, error)
}
