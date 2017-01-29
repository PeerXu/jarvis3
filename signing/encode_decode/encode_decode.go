package encode_decode

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	jerrors "github.com/PeerXu/jarvis3/errors"
	. "github.com/PeerXu/jarvis3/signing/error"
	"github.com/PeerXu/jarvis3/user"
)

func EncodeSignCreateUserRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/signing/v1/users").Methods("POST")
	r.Method, r.URL.Path = "POST", "/signing/v1/users"
	return EncodeRequest(ctx, r, request)
}

func EncodeSignDeleteUserByIDRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/signing/v1/users/{user_id}").Methods("DELETE")
	r.Method, r.URL.Path = "DELETE", "/signing/v1/users/"+ctx.Value("UserID").(string)
	return nil
}

func EncodeSignGetUserByIDRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/signing/v1/users/{user_id}").Methods("GET")
	r.Method, r.URL.Path = "GET", "/signing/v1/users/"+ctx.Value("UserID").(string)
	return nil
}

func EncodeSignLoginRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/signing/v1/users/self/access_tokens").Methods("POST")
	r.Method, r.URL.Path = "POST", "/signing/v1/users/self/access_tokens"
	return EncodeRequest(ctx, r, request)
}

func EncodeSignLogoutRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/signing/v1/users/{user_id}/access_tokens").Methods("DELETE")
	r.Method, r.URL.Path = "DELETE", "/signing/v1/users/"+ctx.Value("UserID").(string)+"/access_tokens"
	return nil
}

func EncodeSignCreateAgentTokenRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/signing/v1/users/{user_id}/agent_tokens").Methods("POST")
	r.Method, r.URL.Path = "POST", "/signing/v1/users/"+ctx.Value("UserID").(string)+"/agent_tokens"
	return EncodeRequest(ctx, r, request)
}

func EncodeSignDeleteAgentTokenRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/signing/v1/users/{user_id}/agent_tokens/{agentTokenName}").Methods("DELETE")
	req := request.(SignDeleteAgentTokenRequest)
	r.Method, r.URL.Path = "DELETE", "/signing/v1/users/"+ctx.Value("UserID").(string)+"/agent_tokens/"+req.Name
	return nil
}

func EncodeSignValidateTokenRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/signing/v1/users/self").Methods("GET")
	r.Method, r.URL.Path = "GET", "/signing/v1/users/self"
	return nil
}

func DecodeSignCreateUserResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	var response SignCreateUserResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func DecodeSignDeleteUserByIDResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	return SignDeleteUserResponse{}, nil
}

func DecodeSignGetUserByIDResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	var response SignGetUserResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func DecodeSignLoginResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	var response SignLoginResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func DecodeSignLogoutResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	return SignLogoutResponse{}, nil
}

func DecodeSignCreateAgentTokenResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	var response SignCreateAgentTokenResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func DecodeSignDeleteAgentTokenResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	return SignDeleteAgentTokenResponse{}, nil
}

func DecodeSignValidateTokenResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	var response SignValidateTokenResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func DecodeSignCreateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return SignCreateUserRequest{
		Username: body.Username,
		Password: body.Password,
		Email:    body.Email,
	}, nil
}

func DecodeSignDeleteUserByIDRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	return SignDeleteUserRequest{ID: userID}, nil
}

func DecodeSignGetUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	return SignGetUserRequest{ID: userID}, nil
}

func DecodeSignLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return SignLoginRequest{
		Username: body.Username,
		Password: body.Password,
	}, nil
}

func DecodeSignLogoutRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	return SignLogoutRequest{ID: userID}, nil
}

func DecodeSignCreateAgentTokenRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return SignCreateAgentTokenRequest{Name: body.Name}, nil
}

func DecodeSignDeleteAgentTokenRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	name := vars["name"]
	return SignDeleteAgentTokenRequest{Name: name}, nil
}

func DecodeSignValidateTokenRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return SignValidateTokenRequest{}, nil
}

func EncodeRequest(ctx context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func MakeResponseEncoder(code int) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		if err, ok := response.(error); ok {
			EncodeError(ctx, err, w)
			return nil
		}
		return EncodeResponse(ctx, w, code, response)
	}
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, code int, response interface{}) error {
	if code == 0 {
		code = http.StatusInternalServerError
	}
	w.WriteHeader(code)

	if code != http.StatusNoContent {
		return json.NewEncoder(w).Encode(response)
	}

	return nil
}

func EncodeError(ctx context.Context, err error, w http.ResponseWriter) {
	jerr, ok := err.(jerrors.JarvisError)
	if !ok {
		jerr = NewSignError(jerrors.ErrorServerError, "unknown error", err)
	}

	var code int
	switch jerr.Type {
	case jerrors.ErrorInvalidRequest:
		code = http.StatusBadRequest
	case jerrors.ErrorAccessDenied:
		code = http.StatusUnauthorized
	case jerrors.ErrorNotFound:
		code = http.StatusNotFound
	case jerrors.ErrorServerError:
	default:
		code = http.StatusInternalServerError
	}
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(jerr)
}

type AccessTokenBody struct {
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expire_at"`
}

type AccessTokensBody []AccessTokenBody

func EncodeAccessToken2AccessTokenBody(t *user.AccessToken) AccessTokenBody {
	return AccessTokenBody{
		Token:    t.Token,
		ExpireAt: t.ExpireAt,
	}
}

func EncodeAccessTokens2AccessTokensBody(ts []*user.AccessToken) AccessTokensBody {
	var ats AccessTokensBody
	for _, t := range ts {
		ats = append(ats, EncodeAccessToken2AccessTokenBody(t))
	}
	return ats
}

func DecodeAccessTokenBody2AccessToken(b AccessTokenBody) *user.AccessToken {
	return &user.AccessToken{
		Token:    b.Token,
		ExpireAt: b.ExpireAt,
	}
}

func DecodeAccessTokensBody2AccessTokens(bs AccessTokensBody) []*user.AccessToken {
	var ats []*user.AccessToken
	for _, b := range bs {
		ats = append(ats, DecodeAccessTokenBody2AccessToken(b))
	}
	return ats
}

type AgentTokenBody struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type AgentTokensBody []AgentTokenBody

func EncodeAgentToken2AgentTokenBody(t *user.AgentToken) AgentTokenBody {
	return AgentTokenBody{
		Name:  t.Name,
		Token: t.Token,
	}
}

func EncodeAgentTokens2AgentTokensBody(ts []*user.AgentToken) AgentTokensBody {
	var ats AgentTokensBody
	for _, t := range ts {
		ats = append(ats, EncodeAgentToken2AgentTokenBody(t))
	}
	return ats
}

func DecodeAgentTokenBody2AgentToken(b AgentTokenBody) *user.AgentToken {
	return &user.AgentToken{
		Name:  b.Name,
		Token: b.Token,
	}
}

func DecodeAgentTokensBody2AgentTokens(bs AgentTokensBody) []*user.AgentToken {
	var ats []*user.AgentToken
	for _, b := range bs {
		ats = append(ats, DecodeAgentTokenBody2AgentToken(b))
	}
	return ats
}

type UserBody struct {
	ID           string            `json:"id"`
	Username     string            `json:"username"`
	Email        string            `json:"email"`
	AccessTokens []AccessTokenBody `json:"access_tokens"`
	AgentTokens  []AgentTokenBody  `json:"agent_tokens"`
}

func EncodeUser2UserBody(u *user.User) UserBody {
	return UserBody{
		ID:           u.ID.String(),
		Username:     u.Username,
		Email:        u.Email.String(),
		AccessTokens: EncodeAccessTokens2AccessTokensBody(u.AccessTokens),
		AgentTokens:  EncodeAgentTokens2AgentTokensBody(u.AgentTokens),
	}
}

func DecodeUserBody2User(b UserBody) *user.User {
	return &user.User{
		ID:           user.UserID(b.ID),
		Username:     b.Username,
		Email:        user.Email(b.Email),
		AccessTokens: DecodeAccessTokensBody2AccessTokens(b.AccessTokens),
		AgentTokens:  DecodeAgentTokensBody2AgentTokens(b.AgentTokens),
	}
}

type SignCreateUserRequest struct {
	Username string
	Password string
	Email    string
}

type SignCreateUserResponse UserBody

type SignDeleteUserRequest struct {
	ID string `json:"-"`
}

type SignDeleteUserResponse struct{}

type SignGetUserRequest struct {
	ID string `json:"-"`
}

type SignGetUserResponse UserBody

type SignLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignLoginResponse UserBody

type SignLogoutRequest struct {
	ID string `json:"-"`
}

type SignLogoutResponse struct{}

func (res SignLogoutResponse) StatusCode() int {
	return http.StatusNoContent
}

type SignCreateAgentTokenRequest struct {
	Name string `json:"name"`
}

type SignCreateAgentTokenResponse AgentTokenBody

type SignDeleteAgentTokenRequest struct {
	Name string
}

type SignDeleteAgentTokenResponse struct{}

func (res SignDeleteAgentTokenResponse) StatusCode() int {
	return http.StatusNoContent
}

type SignValidateTokenRequest struct{}

type SignValidateTokenResponse UserBody

func (res SignValidateTokenResponse) StatusCode() int {
	return http.StatusOK
}
