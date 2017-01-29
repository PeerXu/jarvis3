package endpoint

import (
	"net/url"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"

	jerrors "github.com/PeerXu/jarvis3/errors"
	jhttptransport "github.com/PeerXu/jarvis3/kit/transport/http"
	. "github.com/PeerXu/jarvis3/signing/encode_decode"
	. "github.com/PeerXu/jarvis3/signing/service"
	"github.com/PeerXu/jarvis3/user"
)

type Endpoints struct {
	SignCreateUserEndpoint       endpoint.Endpoint
	SignDeleteUserByIDEndpoint   endpoint.Endpoint
	SignGetUserByIDEndpoint      endpoint.Endpoint
	SignLoginEndpoint            endpoint.Endpoint
	SignLogoutEndpoint           endpoint.Endpoint
	SignCreateAgentTokenEndpoint endpoint.Endpoint
	SignDeleteAgentTokenEndpoint endpoint.Endpoint
	SignValidateTokenEndpoint    endpoint.Endpoint
}

func MakeClientEndpoints(
	instance string,
	encodeRequestFactory jhttptransport.EncodeRequestFuncMiddleware,
	decodeResponseFactory jhttptransport.DecodeResponseFuncMiddleware,
	opts []httptransport.ClientOption) (Endpoints, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	tgt, err := url.Parse(instance)
	if err != nil {
		return Endpoints{}, err
	}

	tgt.Path = ""

	var signCreateUserClient endpoint.Endpoint
	{
		signCreateUserEncodeRequest := encodeRequestFactory(EncodeSignCreateUserRequest)
		signCreateUserDecodeResponse := decodeResponseFactory(DecodeSignCreateUserResponse)
		signCreateUserOptions := opts
		signCreateUserClient = httptransport.NewClient("POST", tgt, signCreateUserEncodeRequest, signCreateUserDecodeResponse, signCreateUserOptions...).Endpoint()
	}

	var signDeleteUserByIDClient endpoint.Endpoint
	{
		signDeleteUserByIDEncodeRequest := encodeRequestFactory(EncodeSignDeleteUserByIDRequest)
		signDeleteUserByIDDecodeResponse := decodeResponseFactory(DecodeSignDeleteUserByIDResponse)
		signDeleteUserByIDOptions := opts
		signDeleteUserByIDClient = httptransport.NewClient("DELETE", tgt, signDeleteUserByIDEncodeRequest, signDeleteUserByIDDecodeResponse, signDeleteUserByIDOptions...).Endpoint()
	}

	var signGetUserByIDClient endpoint.Endpoint
	{
		signGetUserByIDEncodeRequest := encodeRequestFactory(EncodeSignGetUserByIDRequest)
		signGetUserByIDDecodeResponse := decodeResponseFactory(DecodeSignGetUserByIDResponse)
		signGetUserByIDOptions := opts
		signGetUserByIDClient = httptransport.NewClient("GET", tgt, signGetUserByIDEncodeRequest, signGetUserByIDDecodeResponse, signGetUserByIDOptions...).Endpoint()
	}

	var signLoginClient endpoint.Endpoint
	{
		signLoginEncodeRequest := EncodeSignLoginRequest
		signLoginDecodeResponse := decodeResponseFactory(DecodeSignLoginResponse)
		signLoginOptions := opts
		signLoginClient = httptransport.NewClient("POST", tgt, signLoginEncodeRequest, signLoginDecodeResponse, signLoginOptions...).Endpoint()
	}

	var signLogoutClient endpoint.Endpoint
	{
		signLogoutEncodeRequest := encodeRequestFactory(EncodeSignLogoutRequest)
		signLogoutDecodeResponse := decodeResponseFactory(DecodeSignLogoutResponse)
		signLogoutOptions := opts
		signLogoutClient = httptransport.NewClient("DELETE", tgt, signLogoutEncodeRequest, signLogoutDecodeResponse, signLogoutOptions...).Endpoint()
	}

	var signCreateAgentTokenClient endpoint.Endpoint
	{
		signCreateAgentTokenEncodeRequest := encodeRequestFactory(EncodeSignCreateAgentTokenRequest)
		signCreateAgentTokenDecodeResponse := decodeResponseFactory(DecodeSignCreateAgentTokenResponse)
		signCreateAgentTokenOptions := opts
		signCreateAgentTokenClient = httptransport.NewClient("POST", tgt, signCreateAgentTokenEncodeRequest, signCreateAgentTokenDecodeResponse, signCreateAgentTokenOptions...).Endpoint()
	}

	var signDeleteAgentTokenClient endpoint.Endpoint
	{
		signDeleteAgentTokenEncodeRequest := encodeRequestFactory(EncodeSignDeleteAgentTokenRequest)
		signDeleteAgentTokenDecodeResponse := decodeResponseFactory(DecodeSignDeleteAgentTokenResponse)
		signDeleteAgentTokenOptions := opts
		signDeleteAgentTokenClient = httptransport.NewClient("DELETE", tgt, signDeleteAgentTokenEncodeRequest, signDeleteAgentTokenDecodeResponse, signDeleteAgentTokenOptions...).Endpoint()
	}

	var signValidateTokenClient endpoint.Endpoint
	{
		signValidateTokenEncodeRequest := encodeRequestFactory(EncodeSignValidateTokenRequest)
		signValidateTokenDecodeResponse := decodeResponseFactory(DecodeSignValidateTokenResponse)
		signValidateTokenOptions := opts
		signValidateTokenClient = httptransport.NewClient("GET", tgt, signValidateTokenEncodeRequest, signValidateTokenDecodeResponse, signValidateTokenOptions...).Endpoint()
	}

	return Endpoints{
		SignCreateUserEndpoint:       signCreateUserClient,
		SignDeleteUserByIDEndpoint:   signDeleteUserByIDClient,
		SignGetUserByIDEndpoint:      signGetUserByIDClient,
		SignLoginEndpoint:            signLoginClient,
		SignLogoutEndpoint:           signLogoutClient,
		SignCreateAgentTokenEndpoint: signCreateAgentTokenClient,
		SignDeleteAgentTokenEndpoint: signDeleteAgentTokenClient,
		SignValidateTokenEndpoint:    signValidateTokenClient,
	}, nil
}

func (e Endpoints) CreateUser(ctx context.Context, username string, password string, email string) (*user.User, error) {
	request := SignCreateUserRequest{
		Username: username,
		Password: password,
		Email:    email,
	}
	response, err := e.SignCreateUserEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	res := response.(SignCreateUserResponse)
	return &user.User{
		Username:     res.Username,
		Email:        user.Email(res.Email),
		AccessTokens: []*user.AccessToken{},
		AgentTokens:  []*user.AgentToken{},
	}, nil
}

func (e Endpoints) DeleteUserByID(ctx context.Context, id user.UserID) error {
	request := SignDeleteUserRequest{ID: id.String()}
	_, err := e.SignDeleteUserByIDEndpoint(ctx, request)
	return err
}

func (e Endpoints) GetUserByID(ctx context.Context, id user.UserID) (*user.User, error) {
	request := SignGetUserRequest{ID: id.String()}
	response, err := e.SignGetUserByIDEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}

	res := response.(SignGetUserResponse)
	return DecodeUserBody2User(UserBody(res)), nil
}

func (e Endpoints) Login(ctx context.Context, username string, password string) (*user.User, error) {
	request := SignLoginRequest{Username: username, Password: password}
	response, err := e.SignLoginEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}

	if err, ok := response.(jerrors.JarvisError); ok {
		return nil, err
	}

	res := response.(SignLoginResponse)
	return DecodeUserBody2User(UserBody(res)), nil
}

func (e Endpoints) Logout(ctx context.Context, id user.UserID) error {
	request := SignLogoutRequest{ID: id.String()}
	response, err := e.SignLogoutEndpoint(ctx, request)
	if err != nil {
		return err
	}

	if err, ok := response.(jerrors.JarvisError); ok {
		return err
	}

	return err
}

func (e Endpoints) CreateAgentToken(ctx context.Context, name string) (*user.AgentToken, error) {
	request := SignCreateAgentTokenRequest{Name: name}
	response, err := e.SignCreateAgentTokenEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	if err, ok := response.(jerrors.JarvisError); ok {
		return nil, err
	}

	res := response.(SignCreateAgentTokenResponse)
	agentToken := &user.AgentToken{Name: res.Name, Token: res.Token}
	return agentToken, nil
}

func (e Endpoints) DeleteAgentToken(ctx context.Context, name string) error {
	request := SignDeleteAgentTokenRequest{Name: name}
	response, err := e.SignDeleteAgentTokenEndpoint(ctx, request)
	if err != nil {
		return err
	}
	if err, ok := response.(jerrors.JarvisError); ok {
		return err
	}

	return nil
}

func (e Endpoints) ValidateToken(ctx context.Context) (*user.User, error) {
	request := SignValidateTokenRequest{}
	response, err := e.SignValidateTokenEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	if err, ok := response.(jerrors.JarvisError); ok {
		return nil, err
	}

	res := response.(SignValidateTokenResponse)

	var accessTokens []*user.AccessToken
	for _, t := range res.AccessTokens {
		accessTokens = append(accessTokens, &user.AccessToken{Token: t.Token, ExpireAt: t.ExpireAt})
	}

	var agentTokens []*user.AgentToken
	for _, t := range res.AgentTokens {
		agentTokens = append(agentTokens, &user.AgentToken{Token: t.Token, Name: t.Name})
	}

	u := &user.User{
		Username:     res.Username,
		Email:        user.Email(res.Email),
		AccessTokens: accessTokens,
		AgentTokens:  agentTokens,
	}
	return u, nil
}

func MakeSignCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignCreateUserRequest)
		u, err := s.CreateUser(ctx, req.Username, req.Password, req.Email)
		if err != nil {
			return err, nil
		}
		return EncodeUser2UserBody(u), nil
	}
}

func MakeSignDeleteUserByIDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignDeleteUserRequest)
		err := s.DeleteUserByID(ctx, user.UserID(req.ID))
		if err != nil {
			return err, nil
		}
		return SignDeleteUserResponse{}, nil
	}
}

func MakeSignGetUserByIDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignGetUserRequest)
		u, err := s.GetUserByID(ctx, user.UserID(req.ID))
		if err != nil {
			return err, nil
		}

		return EncodeUser2UserBody(u), nil
	}
}

func MakeSignLoginEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignLoginRequest)
		u, err := s.Login(ctx, req.Username, req.Password)
		if err != nil {
			return err, nil
		}
		return EncodeUser2UserBody(u), nil
	}
}

func MakeSignLogoutEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignLogoutRequest)
		err := s.Logout(ctx, user.UserID(req.ID))
		if err != nil {
			return err, nil
		}
		return SignLogoutResponse{}, nil
	}
}

func MakeSignCreateAgentTokenEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignCreateAgentTokenRequest)
		t, err := s.CreateAgentToken(ctx, req.Name)
		if err != nil {
			return err, nil
		}
		return EncodeAgentToken2AgentTokenBody(t), nil
	}
}

func MakeSignDeleteAgentTokenEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignDeleteAgentTokenRequest)
		err := s.DeleteAgentToken(ctx, req.Name)
		if err != nil {
			return err, nil
		}
		return SignDeleteAgentTokenResponse{}, nil
	}
}

func MakeSignValidateTokenEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		u, err := s.ValidateToken(ctx)
		if err != nil {
			return err, nil
		}
		return EncodeUser2UserBody(u), nil
	}
}
