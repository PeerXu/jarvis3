package signing

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"

	jhttptransport "github.com/PeerXu/jarvis3/kit/transport/http"
	"github.com/PeerXu/jarvis3/user"
	"github.com/PeerXu/jarvis3/utils"
)

type Endpoints struct {
	SignCreateUserEndpoint          endpoint.Endpoint
	SignDeleteUserEndpoint          endpoint.Endpoint
	SignGetUserEndpoint             endpoint.Endpoint
	SignLoginEndpoint               endpoint.Endpoint
	SignLogoutEndpoint              endpoint.Endpoint
	SignCreateAgentTokenEndpoint    endpoint.Endpoint
	SignDeleteAgentTokenEndpoint    endpoint.Endpoint
	SignValidateAccessTokenEndpoint endpoint.Endpoint
}

func MakeClientEndpoints(instance string, environment utils.Environment) (Endpoints, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	tgt, err := url.Parse(instance)
	if err != nil {
		return Endpoints{}, err
	}

	tgt.Path = ""

	opts := []httptransport.ClientOption{}

	var signCreateUserClient endpoint.Endpoint
	{
		signCreateUserEncodeRequest := jhttptransport.EncodeRequestFuncChain(
			LoadEnvironmentIntoContextMiddleware(environment),
			AssignAccessTokenToContextMiddleware,
		)(encodeSignCreateUserRequest)
		signCreateUserDecodeResponse := jhttptransport.DecodeResponseFuncChain(
			DecodeErrorResponseMiddleware,
		)(decodeSignCreateUserResponse)
		signCreateUserOptions := opts
		signCreateUserClient = httptransport.NewClient("POST", tgt, signCreateUserEncodeRequest, signCreateUserDecodeResponse, signCreateUserOptions...).Endpoint()
	}

	var signDeleteUserClient endpoint.Endpoint
	{
		signDeleteUserEncodeRequest := jhttptransport.EncodeRequestFuncChain(
			LoadEnvironmentIntoContextMiddleware(environment),
			AssignAccessTokenToContextMiddleware,
		)(encodeSignDeleteUserRequest)
		signDeleteUserDecodeResponse := jhttptransport.DecodeResponseFuncChain(
			DecodeErrorResponseMiddleware,
		)(decodeSignDeleteUserResponse)
		signDeleteUserOptions := opts
		signDeleteUserClient = httptransport.NewClient("DELETE", tgt, signDeleteUserEncodeRequest, signDeleteUserDecodeResponse, signDeleteUserOptions...).Endpoint()
	}

	var signGetUserClient endpoint.Endpoint
	{
		signGetUserEncodeRequest := jhttptransport.EncodeRequestFuncChain(
			LoadEnvironmentIntoContextMiddleware(environment),
			AssignAccessTokenToContextMiddleware,
		)(encodeSignGetUserRequest)
		signGetUserDecodeResponse := jhttptransport.DecodeResponseFuncChain(
			DecodeErrorResponseMiddleware,
		)(decodeSignGetUserResponse)
		signGetUserOptions := opts
		signGetUserClient = httptransport.NewClient("GET", tgt, signGetUserEncodeRequest, signGetUserDecodeResponse, signGetUserOptions...).Endpoint()
	}

	var signLoginClient endpoint.Endpoint
	{
		signLoginEncodeRequest := encodeSignLoginRequest
		signLoginDecodeResponse := jhttptransport.DecodeResponseFuncChain(
			DecodeErrorResponseMiddleware,
		)(decodeSignLoginResponse)
		signLoginOptions := opts
		signLoginClient = httptransport.NewClient("POST", tgt, signLoginEncodeRequest, signLoginDecodeResponse, signLoginOptions...).Endpoint()
	}

	var signLogoutClient endpoint.Endpoint
	{
		signLogoutEncodeRequest := jhttptransport.EncodeRequestFuncChain(
			LoadEnvironmentIntoContextMiddleware(environment),
			AssignAccessTokenToContextMiddleware,
		)(encodeSignLogoutRequest)
		signLogoutDecodeResponse := jhttptransport.DecodeResponseFuncChain(
			DecodeErrorResponseMiddleware,
		)(decodeSignLogoutResponse)
		signLogoutOptions := opts
		signLogoutClient = httptransport.NewClient("DELETE", tgt, signLogoutEncodeRequest, signLogoutDecodeResponse, signLogoutOptions...).Endpoint()
	}

	var signCreateAgentTokenClient endpoint.Endpoint
	{
		signCreateAgentTokenEncodeRequest := jhttptransport.EncodeRequestFuncChain(
			LoadEnvironmentIntoContextMiddleware(environment),
			AssignAccessTokenToContextMiddleware,
		)(encodeSignCreateAgentTokenRequest)
		signCreateAgentTokenDecodeResponse := jhttptransport.DecodeResponseFuncChain(
			DecodeErrorResponseMiddleware,
		)(decodeSignCreateAgentTokenResponse)
		signCreateAgentTokenOptions := opts
		signCreateAgentTokenClient = httptransport.NewClient("POST", tgt, signCreateAgentTokenEncodeRequest, signCreateAgentTokenDecodeResponse, signCreateAgentTokenOptions...).Endpoint()
	}

	var signDeleteAgentTokenClient endpoint.Endpoint
	{
		signDeleteAgentTokenEncodeRequest := jhttptransport.EncodeRequestFuncChain(
			LoadEnvironmentIntoContextMiddleware(environment),
			AssignAccessTokenToContextMiddleware,
		)(encodeSignDeleteAgentTokenRequest)
		signDeleteAgentTokenDecodeResponse := jhttptransport.DecodeResponseFuncChain(
			DecodeErrorResponseMiddleware,
		)(decodeSignDeleteAgentTokenResponse)
		signDeleteAgentTokenOptions := opts
		signDeleteAgentTokenClient = httptransport.NewClient("DELETE", tgt, signDeleteAgentTokenEncodeRequest, signDeleteAgentTokenDecodeResponse, signDeleteAgentTokenOptions...).Endpoint()
	}

	var signValidateAccessTokenClient endpoint.Endpoint
	{
		signValidateAccessTokenEncodeRequest := jhttptransport.EncodeRequestFuncChain(
			LoadEnvironmentIntoContextMiddleware(environment),
			AssignAccessTokenToContextMiddleware,
		)(encodeSignValidateAccessTokenRequest)
		signValidateAccessTokenDecodeResponse := jhttptransport.DecodeResponseFuncChain(
			DecodeErrorResponseMiddleware,
		)(decodeSignValidateAccessTokenResponse)
		signValidateAccessTokenOptions := opts
		signValidateAccessTokenClient = httptransport.NewClient("GET", tgt, signValidateAccessTokenEncodeRequest, signValidateAccessTokenDecodeResponse, signValidateAccessTokenOptions...).Endpoint()
	}

	return Endpoints{
		SignCreateUserEndpoint:          signCreateUserClient,
		SignDeleteUserEndpoint:          signDeleteUserClient,
		SignGetUserEndpoint:             signGetUserClient,
		SignLoginEndpoint:               signLoginClient,
		SignLogoutEndpoint:              signLogoutClient,
		SignCreateAgentTokenEndpoint:    signCreateAgentTokenClient,
		SignDeleteAgentTokenEndpoint:    signDeleteAgentTokenClient,
		SignValidateAccessTokenEndpoint: signValidateAccessTokenClient,
	}, nil
}

func (e Endpoints) CreateUser(ctx context.Context, username string, password string, email string) (*user.User, error) {
	request := signCreateUserRequest{
		Username: username,
		Password: password,
		Email:    email,
	}
	response, err := e.SignCreateUserEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	res := response.(signCreateUserResponse)
	return &user.User{
		Username:     res.Username,
		Email:        user.Email(res.Email),
		AccessTokens: []*user.AccessToken{},
		AgentTokens:  []*user.AgentToken{},
	}, nil
}

func (e Endpoints) DeleteUser(ctx context.Context, username string) error {
	request := signDeleteUserRequest{Username: username}
	_, err := e.SignDeleteUserEndpoint(ctx, request)
	return err
}

func (e Endpoints) GetUser(ctx context.Context, username string) (*user.User, error) {
	request := signGetUserRequest{Username: username}
	response, err := e.SignGetUserEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}

	res := response.(signGetUserResponse)
	var accessTokens []*user.AccessToken
	for _, t := range res.AccessTokens {
		accessTokens = append(accessTokens, &user.AccessToken{Token: t.Token, ExpireAt: t.ExpireAt})
	}
	var agentTokens []*user.AgentToken
	for _, t := range res.AgentTokens {
		agentTokens = append(agentTokens, &user.AgentToken{Name: t.Name, Token: t.Token})
	}
	return &user.User{
		Username:     res.Username,
		Email:        user.Email(res.Email),
		AccessTokens: accessTokens,
		AgentTokens:  agentTokens,
	}, nil
}

func (e Endpoints) Login(ctx context.Context, username string, password string) (*user.AccessToken, error) {
	request := signLoginRequest{Username: username, Password: password}
	response, err := e.SignLoginEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}

	if err, ok := convertSignError(response); ok {
		return nil, err
	}

	res := response.(signLoginResponse)
	accessToken := &user.AccessToken{Token: res.AccessToken, ExpireAt: res.ExpireAt}
	return accessToken, nil
}

func (e Endpoints) Logout(ctx context.Context, username string) error {
	request := signLogoutRequest{Username: username}
	response, err := e.SignLogoutEndpoint(ctx, request)
	if err != nil {
		return err
	}

	if err, ok := convertSignError(response); ok {
		return err
	}

	return err
}

func (e Endpoints) CreateAgentToken(ctx context.Context, name string) (*user.AgentToken, error) {
	request := signCreateAgentTokenRequest{Name: name}
	response, err := e.SignCreateAgentTokenEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	if err, ok := convertSignError(response); ok {
		return nil, err
	}

	res := response.(signCreateAgentTokenResponse)
	agentToken := &user.AgentToken{Name: res.Name, Token: res.Token}
	return agentToken, nil
}

func (e Endpoints) DeleteAgentToken(ctx context.Context, name string) error {
	request := signDeleteAgentTokenRequest{Name: name}
	response, err := e.SignDeleteAgentTokenEndpoint(ctx, request)
	if err != nil {
		return err
	}
	if err, ok := convertSignError(response); ok {
		return err
	}

	return nil
}

func (e Endpoints) ValidateAccessToken(ctx context.Context) (*user.User, error) {
	request := signValidateAccessTokenRequest{}
	response, err := e.SignValidateAccessTokenEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	if err, ok := convertSignError(response); ok {
		return nil, err
	}

	res := response.(signValidateAccessTokenResponse)

	var accessTokens []*user.AccessToken
	for _, t := range res.AccessTokens {
		accessTokens = append(accessTokens, &user.AccessToken{Token: t.Token, ExpireAt: t.ExpireAt})
	}

	u := &user.User{
		Username:     res.Username,
		Email:        user.Email(res.Email),
		AccessTokens: accessTokens,
	}
	return u, nil
}

type accessTokenBody struct {
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expireAt"`
}

func encodeUser2AccessTokens(u *user.User) []*accessTokenBody {
	accessTokens := []*accessTokenBody{}
	for _, t := range u.AccessTokens {
		accessTokens = append(accessTokens, &accessTokenBody{
			Token:    t.Token,
			ExpireAt: t.ExpireAt,
		})
	}
	return accessTokens
}

type agentTokenBody struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

func encodeUser2AgentTokens(u *user.User) []*agentTokenBody {
	agentTokens := []*agentTokenBody{}
	for _, t := range u.AgentTokens {
		agentTokens = append(agentTokens, &agentTokenBody{
			Name:  t.Name,
			Token: t.Token,
		})
	}
	return agentTokens
}

type signCreateUserRequest struct {
	Username string
	Password string
	Email    string
}

type signCreateUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func makeSignCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(signCreateUserRequest)
		u, err := s.CreateUser(ctx, req.Username, req.Password, req.Email)
		if err != nil {
			return err, nil
		}
		return signCreateUserResponse{
			Username: u.Username,
			Email:    u.Email.String(),
		}, nil
	}
}

type signDeleteUserRequest struct {
	Username string `json:"-"`
}

type signDeleteUserResponse struct{}

func makeSignDeleteUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(signDeleteUserRequest)
		err := s.DeleteUser(ctx, req.Username)
		if err != nil {
			return err, nil
		}
		return signDeleteUserResponse{}, nil
	}
}

type signGetUserRequest struct {
	Username string
}

type signGetUserResponse struct {
	Username     string             `json:"username"`
	Email        string             `json:"email"`
	AccessTokens []*accessTokenBody `json:"accessTokens"`
	AgentTokens  []*agentTokenBody  `json:"agentTokens"`
}

func encodeUser2signGetUserResponse(u *user.User) signGetUserResponse {
	res := signGetUserResponse{
		Username:     u.Username,
		Email:        u.Email.String(),
		AccessTokens: encodeUser2AccessTokens(u),
		AgentTokens:  encodeUser2AgentTokens(u),
	}
	return res
}

func makeSignGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(signGetUserRequest)
		u, err := s.GetUser(ctx, req.Username)
		if err != nil {
			return err, nil
		}

		return encodeUser2signGetUserResponse(u), nil
	}
}

type signLoginRequest struct {
	Username string `json:"-"`
	Password string `json:"password"`
}

type signLoginResponse struct {
	AccessToken string    `json:"accessToken"`
	ExpireAt    time.Time `json:"expireAt"`
}

func makeSignLoginEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(signLoginRequest)
		t, err := s.Login(ctx, req.Username, req.Password)
		if err != nil {
			return err, nil
		}
		return signLoginResponse{
			AccessToken: t.Token,
			ExpireAt:    t.ExpireAt,
		}, nil
	}
}

type signLogoutRequest struct {
	Username string `json:"-"`
}

type signLogoutResponse struct{}

func (res signLogoutResponse) StatusCode() int {
	return http.StatusNoContent
}

func makeSignLogoutEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(signLogoutRequest)
		err := s.Logout(ctx, req.Username)
		if err != nil {
			return err, nil
		}
		return signLogoutResponse{}, nil
	}
}

type signCreateAgentTokenRequest struct {
	Name string `json:"name"`
}

type signCreateAgentTokenResponse struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

func makeSignCreateAgentTokenEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(signCreateAgentTokenRequest)
		t, err := s.CreateAgentToken(ctx, req.Name)
		if err != nil {
			return err, nil
		}
		return signCreateAgentTokenResponse{
			Name:  t.Name,
			Token: t.Token,
		}, nil
	}
}

type signDeleteAgentTokenRequest struct {
	Name string
}

type signDeleteAgentTokenResponse struct{}

func (res signDeleteAgentTokenResponse) StatusCode() int {
	return http.StatusNoContent
}

func makeSignDeleteAgentTokenEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(signDeleteAgentTokenRequest)
		err := s.DeleteAgentToken(ctx, req.Name)
		if err != nil {
			return err, nil
		}
		return signDeleteAgentTokenResponse{}, nil
	}
}

type signValidateAccessTokenRequest struct{}

type signValidateAccessTokenResponse struct {
	Username     string             `json:"username"`
	Email        string             `json:"email"`
	AccessTokens []*accessTokenBody `json:"accessTokens"`
}

func (res signValidateAccessTokenResponse) StatusCode() int {
	return http.StatusOK
}

func makeSignValidateAccessTokenEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		u, err := s.ValidateAccessToken(ctx)
		if err != nil {
			return err, nil
		}

		res := signValidateAccessTokenResponse{
			Username:     u.Username,
			Email:        string(u.Email),
			AccessTokens: encodeUser2AccessTokens(u),
		}
		return res, nil
	}
}
