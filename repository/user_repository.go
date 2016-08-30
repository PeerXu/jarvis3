package repository

import (
	"sync"

	"github.com/PeerXu/jarvis3/user"
)

type userRepository struct {
	mtx          sync.RWMutex
	users        map[string]*user.User
	accessTokens map[string]*user.User
	agentTokens  map[string]*user.User
}

func NewUserRepository() *userRepository {
	return &userRepository{
		mtx:          sync.RWMutex{},
		users:        map[string]*user.User{"admin": user.NewUser("admin", "admin", "jarvis3@gmail.com")},
		accessTokens: map[string]*user.User{},
		agentTokens:  map[string]*user.User{},
	}
}

func (r *userRepository) CreateUser(u *user.User) (*user.User, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	r.users[u.Username] = u
	return u, nil
}

func (r *userRepository) DeleteUser(username string) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	delete(r.users, username)
	return nil
}

func (r *userRepository) GetUser(username string) (*user.User, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	if u, ok := r.users[username]; ok {
		return u, nil
	}

	return nil, user.ErrUserNotFound
}

func (r *userRepository) CreateAccessToken(u *user.User, t *user.AccessToken) error {
	u, err := r.GetUser(u.Username)
	if err != nil {
		return err
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	for _, at := range u.AccessTokens {
		if at.Token == t.Token {
			return user.ErrUnknown
		}
	}

	u.AccessTokens = append(u.AccessTokens, t)
	r.accessTokens[t.Token] = u

	return nil
}

func (r *userRepository) DeleteAccessTokens(u *user.User, ts []*user.AccessToken) error {
	u, err := r.GetUser(u.Username)
	if err != nil {
		return err
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	var nts []*user.AccessToken

	for _, et := range u.AccessTokens {
		flg := false
		for j, dt := range ts {
			if dt.Token == et.Token {
				flg = true
				delete(r.accessTokens, dt.Token)
				ts = append(ts[:j], ts[j+1:]...)
				break
			}
		}
		if !flg {
			nts = append(nts, et)
		}
	}
	u.AccessTokens = nts

	return nil
}

func (r *userRepository) LookupUserByAccessToken(t *user.AccessToken) (*user.User, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	u, ok := r.accessTokens[t.Token]
	if !ok {
		return nil, user.ErrAccessTokenNotFound
	}
	return u, nil

}

func (r *userRepository) CreateAgentToken(u *user.User, t *user.AgentToken) error {
	u, err := r.GetUser(u.Username)
	if err != nil {
		return err
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	for _, at := range u.AgentTokens {
		if at.Token == t.Token {
			return user.ErrUnknown
		}
	}

	u.AgentTokens = append(u.AgentTokens, t)
	r.agentTokens[t.Token] = u

	return nil
}

func (r *userRepository) DeleteAgentTokens(u *user.User, ts []*user.AgentToken) error {
	u, err := r.GetUser(u.Username)
	if err != nil {
		return err
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	var nts []*user.AgentToken

	for _, et := range u.AgentTokens {
		flg := false
		for j, dt := range ts {
			if dt.Token == et.Token {
				flg = true
				delete(r.agentTokens, dt.Token)
				ts = append(ts[:j], ts[j+1:]...)
				break
			}
		}
		if !flg {
			nts = append(nts, et)
		}
	}
	u.AgentTokens = nts

	return nil
}

func (r *userRepository) LookupUserByAgentToken(t *user.AgentToken) (*user.User, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	u, ok := r.agentTokens[t.Token]
	if !ok {
		return nil, user.ErrAgentTokenNotFound
	}

	return u, nil
}

func (r *userRepository) LookupAgentTokenByName(u *user.User, n string) (*user.AgentToken, error) {
	u, err := r.GetUser(u.Username)
	if err != nil {
		return nil, err
	}

	r.mtx.RLock()
	defer r.mtx.RUnlock()

	for _, t := range u.AgentTokens {
		if t.Name == n {
			return t, nil
		}
	}

	return nil, user.ErrAgentTokenNotFound
}
