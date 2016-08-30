package signing

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	jkth "github.com/PeerXu/jarvis3/kit/transport/http"
)

func MakeHandler(ctx context.Context, s Service) http.Handler {
	r := mux.NewRouter()

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	var signCreateUserHandler *kithttp.Server
	{
		signCreateUserEndpoint := makeSignCreateUserEndpoint(s)
		signCreateUserEndpoint = endpoint.Chain(
			ValidateAccessTokenMiddeware(s),
		)(signCreateUserEndpoint)
		signCreateUserOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		signCreateUserHandler = kithttp.NewServer(
			ctx,
			signCreateUserEndpoint,
			decodeSignCreateUserRequest,
			makeResponseEncoder(http.StatusOK),
			signCreateUserOptions...,
		)
	}

	var signDeleteUserHandler *kithttp.Server
	{
		signDeleteUserEndpoint := makeSignDeleteUserEndpoint(s)
		signDeleteUserEndpoint = endpoint.Chain(
			ValidateAccessTokenMiddeware(s),
		)(signDeleteUserEndpoint)
		signDeleteUserOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		signDeleteUserHandler = kithttp.NewServer(
			ctx,
			signDeleteUserEndpoint,
			decodeSignDeleteUserRequest,
			makeResponseEncoder(http.StatusNoContent),
			signDeleteUserOptions...,
		)
	}

	var signGetUserHandler *kithttp.Server
	{
		signGetUserEndpoint := makeSignGetUserEndpoint(s)
		signGetUserEndpoint = endpoint.Chain(
			ValidateAccessTokenMiddeware(s),
		)(signGetUserEndpoint)
		signGetUserOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		signGetUserHandler = kithttp.NewServer(
			ctx,
			signGetUserEndpoint,
			decodeSignGetUserRequest,
			makeResponseEncoder(http.StatusOK),
			signGetUserOptions...,
		)
	}

	var signLoginHandler *kithttp.Server
	{
		signLoginEndpoint := makeSignLoginEndpoint(s)
		signLoginOptions := opts
		signLoginHandler = kithttp.NewServer(
			ctx,
			signLoginEndpoint,
			decodeSignLoginRequest,
			makeResponseEncoder(http.StatusOK),
			signLoginOptions...,
		)
	}

	var signLogoutHandler *kithttp.Server
	{
		signLogoutEndpoint := makeSignLogoutEndpoint(s)
		signLogoutEndpoint = endpoint.Chain(
			ValidateAccessTokenMiddeware(s),
		)(signLogoutEndpoint)
		signLogoutOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		signLogoutHandler = kithttp.NewServer(
			ctx,
			signLogoutEndpoint,
			decodeSignLogoutRequest,
			makeResponseEncoder(http.StatusNoContent),
			signLogoutOptions...,
		)
	}

	var signCreateAgentTokenHandler *kithttp.Server
	{
		signCreateAgentTokenEndpoint := makeSignCreateAgentTokenEndpoint(s)
		signCreateAgentTokenEndpoint = endpoint.Chain(
			ValidateAccessTokenMiddeware(s),
		)(signCreateAgentTokenEndpoint)
		signCreateAgentTokenOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		signCreateAgentTokenHandler = kithttp.NewServer(
			ctx,
			signCreateAgentTokenEndpoint,
			decodeSignCreateAgentTokenRequest,
			makeResponseEncoder(http.StatusOK),
			signCreateAgentTokenOptions...,
		)
	}

	var signDeleteAgentTokenHandler *kithttp.Server
	{
		signDeleteAgentTokenEndpoint := makeSignDeleteAgentTokenEndpoint(s)
		signDeleteAgentTokenEndpoint = endpoint.Chain(
			ValidateAccessTokenMiddeware(s),
		)(signDeleteAgentTokenEndpoint)
		signDeleteAgentTokenOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		signDeleteAgentTokenHandler = kithttp.NewServer(
			ctx,
			signDeleteAgentTokenEndpoint,
			decodeSignDeleteAgentTokenRequest,
			makeResponseEncoder(http.StatusNoContent),
			signDeleteAgentTokenOptions...,
		)
	}

	var signValidateAccessTokenHandler *kithttp.Server
	{
		signValidateAccessTokenEndpoint := makeSignValidateAccessTokenEndpoint(s)
		signValidateAccessTokenOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		signValidateAccessTokenHandler = kithttp.NewServer(
			ctx,
			signValidateAccessTokenEndpoint,
			decodeSignValidateAccessTokenRequest,
			makeResponseEncoder(http.StatusOK),
			signValidateAccessTokenOptions...,
		)
	}

	r.Handle("/signing/v1/users", signCreateUserHandler).Methods("POST")
	r.Handle("/signing/v1/users/self", signValidateAccessTokenHandler).Methods("GET")
	r.Handle("/signing/v1/users/{username}", signDeleteUserHandler).Methods("DELETE")
	r.Handle("/signing/v1/users/{username}", signGetUserHandler).Methods("GET")
	r.Handle("/signing/v1/users/{username}/accessTokens", signLoginHandler).Methods("POST")
	r.Handle("/signing/v1/users/{username}/accessTokens", signLogoutHandler).Methods("DELETE")
	r.Handle("/signing/v1/users/{username}/agentTokens", signCreateAgentTokenHandler).Methods("POST")
	r.Handle("/signing/v1/users/{username}/agentTokens/{name}", signDeleteAgentTokenHandler).Methods("DELETE")

	return r
}

func encodeSignCreateUserRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/signing/v1/users").Methods("POST")
	r.Method, r.URL.Path = "POST", "/signing/v1/users"
	return encodeRequest(ctx, r, request)
}

func encodeSignDeleteUserRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/signing/v1/users/{username}").Methods("DELETE")
	r.Method, r.URL.Path = "DELETE", "/signing/v1/users/"+ctx.Value("Username").(string)
	return nil
}

func encodeSignGetUserRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/signing/v1/users/{username}").Methods("GET")
	r.Method, r.URL.Path = "GET", "/signing/v1/users/"+ctx.Value("Username").(string)
	return nil
}

func encodeSignLoginRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/signing/v1/users/{username}/accessTokens").Methods("POST")
	req := request.(signLoginRequest)
	r.Method, r.URL.Path = "POST", "/signing/v1/users/"+req.Username+"/accessTokens"
	return encodeRequest(ctx, r, request)
}

func encodeSignLogoutRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/signing/v1/users/{username}/accessTokens").Methods("DELETE")
	r.Method, r.URL.Path = "DELETE", "/signing/v1/users/"+ctx.Value("Username").(string)+"/accessTokens"
	return nil
}

func encodeSignCreateAgentTokenRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/signing/v1/users/{username}/agentTokens").Methods("POST")
	r.Method, r.URL.Path = "POST", "/signing/v1/users/"+ctx.Value("Username").(string)+"/agentTokens"
	return encodeRequest(ctx, r, request)
}

func encodeSignDeleteAgentTokenRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/signing/v1/users/{username}/agentTokens/{agentTokenName}").Methods("DELETE")
	req := request.(signDeleteAgentTokenRequest)
	r.Method, r.URL.Path = "DELETE", "/signing/v1/users/"+ctx.Value("Username").(string)+"/agentTokens/"+req.Name
	return nil
}

func encodeSignValidateAccessTokenRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/signing/v1/users/self").Methods("GET")
	r.Method, r.URL.Path = "GET", "/signing/v1/users/self"
	return nil
}

func decodeSignCreateUserResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	var response signCreateUserResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func decodeSignDeleteUserResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	return signDeleteUserResponse{}, nil
}

func decodeSignGetUserResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	var response signGetUserResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func decodeSignLoginResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	var response signLoginResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func decodeSignLogoutResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	return signLogoutResponse{}, nil
}

func decodeSignCreateAgentTokenResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	var response signCreateAgentTokenResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func decodeSignDeleteAgentTokenResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	return signDeleteAgentTokenResponse{}, nil
}

func decodeSignValidateAccessTokenResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	var response signValidateAccessTokenResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func decodeSignCreateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return signCreateUserRequest{
		Username: body.Username,
		Password: body.Password,
		Email:    body.Email,
	}, nil
}

func decodeSignDeleteUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	username := vars["username"]
	return signDeleteUserRequest{Username: username}, nil
}

func decodeSignGetUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	username := vars["username"]
	return signGetUserRequest{Username: username}, nil
}

func decodeSignLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	vars := mux.Vars(r)
	username := vars["username"]
	body.Username = username

	return signLoginRequest{
		Username: body.Username,
		Password: body.Password,
	}, nil
}

func decodeSignLogoutRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	username := vars["username"]
	return signLogoutRequest{Username: username}, nil
}

func decodeSignCreateAgentTokenRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return signCreateAgentTokenRequest{Name: body.Name}, nil
}

func decodeSignDeleteAgentTokenRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	name := vars["name"]
	return signDeleteAgentTokenRequest{Name: name}, nil
}

func decodeSignValidateAccessTokenRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return signValidateAccessTokenRequest{}, nil
}

func encodeRequest(ctx context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func makeResponseEncoder(code int) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		if err, ok := response.(error); ok {
			encodeError(ctx, err, w)
			return nil
		}
		return encodeResponse(ctx, w, code, response)
	}
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, code int, response interface{}) error {
	if code == 0 {
		code = http.StatusInternalServerError
	}
	w.WriteHeader(code)

	if code != http.StatusNoContent {
		return json.NewEncoder(w).Encode(response)
	}

	return nil
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	serr, ok := err.(*signError)
	if !ok {
		serr = newSignError(errorServerError, "unknown error", err)
	}

	var code int
	switch serr.Type {
	case errorInvalidRequest:
		code = http.StatusBadRequest
	case errorAccessDenied:
		code = http.StatusUnauthorized
	case errorNotFound:
		code = http.StatusNotFound
	case errorServerError:
	default:
		code = http.StatusInternalServerError
	}
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(serr)
}
