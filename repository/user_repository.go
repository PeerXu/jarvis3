package repository

import (
	"sync"

	"github.com/PeerXu/jarvis3/user"
)

type userRepository struct {
	mtx          sync.RWMutex
	users        map[user.UserID]*user.User
	accessTokens map[string]user.UserID
	agentTokens  map[string]user.UserID
}

func NewUserRepository() *userRepository {
	admin := user.NewUser("admin", "admin", "pppeerxu@gmail.com")
	return &userRepository{
		mtx:          sync.RWMutex{},
		users:        map[user.UserID]*user.User{admin.ID: admin},
		accessTokens: map[string]user.UserID{},
		agentTokens:  map[string]user.UserID{},
	}
}

func (r *userRepository) CreateUser(u *user.User) (*user.User, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	r.users[u.ID] = u
	return u, nil
}

func (r *userRepository) DeleteUserByID(id user.UserID) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	delete(r.users, id)
	return nil
}

func (r *userRepository) GetUserByID(id user.UserID) (*user.User, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	if u, ok := r.users[id]; ok {
		return u, nil
	}
	return nil, user.ErrUserNotFound
}

func (r *userRepository) CreateAccessToken(id user.UserID, t *user.AccessToken) error {
	u, err := r.GetUserByID(id)
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
	r.accessTokens[t.Token] = id

	return nil
}

func (r *userRepository) DeleteAccessTokens(id user.UserID, ts []*user.AccessToken) error {
	u, err := r.GetUserByID(id)
	if err != nil {
		return err
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	var nts []*user.AccessToken

	for _, et := range u.AccessTokens {
		found := false
		for j, dt := range ts {
			if dt.Token == et.Token {
				found = true
				delete(r.accessTokens, dt.Token)
				ts = append(ts[:j], ts[j+1:]...)
				break
			}
		}
		if !found {
			nts = append(nts, et)
		}
	}
	u.AccessTokens = nts

	return nil
}

func (r *userRepository) LookupUserByUsername(username string) (*user.User, error) {
	for _, u := range r.users {
		if u.Username == username {
			return u, nil
		}
	}
	return nil, user.ErrUserNotFound
}

func (r *userRepository) LookupUserByAccessToken(t *user.AccessToken) (*user.User, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	userID, ok := r.accessTokens[t.Token]
	if !ok {
		return nil, user.ErrAccessTokenNotFound
	}
	u, err := r.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return u, nil

}

func (r *userRepository) CreateAgentToken(id user.UserID, t *user.AgentToken) error {
	u, err := r.GetUserByID(id)
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
	r.agentTokens[t.Token] = id

	return nil
}

func (r *userRepository) DeleteAgentTokens(id user.UserID, ts []*user.AgentToken) error {
	u, err := r.GetUserByID(id)
	if err != nil {
		return err
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	var nts []*user.AgentToken

	for _, et := range u.AgentTokens {
		found := false
		for j, dt := range ts {
			if dt.Token == et.Token {
				found = true
				delete(r.agentTokens, dt.Token)
				ts = append(ts[:j], ts[j+1:]...)
				break
			}
		}
		if !found {
			nts = append(nts, et)
		}
	}
	u.AgentTokens = nts

	return nil
}

func (r *userRepository) LookupUserByAgentToken(t *user.AgentToken) (*user.User, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	userID, ok := r.agentTokens[t.Token]
	if !ok {
		return nil, user.ErrAgentTokenNotFound
	}
	u, err := r.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *userRepository) LookupAgentTokenByName(id user.UserID, n string) (*user.AgentToken, error) {
	u, err := r.GetUserByID(id)
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
