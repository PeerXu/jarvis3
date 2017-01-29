package signing

import (
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	jkth "github.com/PeerXu/jarvis3/kit/transport/http"
	. "github.com/PeerXu/jarvis3/signing/encode_decode"
	. "github.com/PeerXu/jarvis3/signing/endpoint"
	. "github.com/PeerXu/jarvis3/signing/service"
)

func MakeHandler(ctx context.Context, s Service) http.Handler {
	r := mux.NewRouter()

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(EncodeError),
	}

	var signCreateUserHandler *kithttp.Server
	{
		signCreateUserEndpoint := MakeSignCreateUserEndpoint(s)
		signCreateUserEndpoint = endpoint.Chain(
			ValidateTokenMiddeware(s),
		)(signCreateUserEndpoint)
		signCreateUserOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		signCreateUserHandler = kithttp.NewServer(
			ctx,
			signCreateUserEndpoint,
			DecodeSignCreateUserRequest,
			MakeResponseEncoder(http.StatusOK),
			signCreateUserOptions...,
		)
	}

	var signDeleteUserHandler *kithttp.Server
	{
		signDeleteUserEndpoint := MakeSignDeleteUserByIDEndpoint(s)
		signDeleteUserEndpoint = endpoint.Chain(
			ValidateTokenMiddeware(s),
		)(signDeleteUserEndpoint)
		signDeleteUserOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		signDeleteUserHandler = kithttp.NewServer(
			ctx,
			signDeleteUserEndpoint,
			DecodeSignDeleteUserByIDRequest,
			MakeResponseEncoder(http.StatusNoContent),
			signDeleteUserOptions...,
		)
	}

	var signGetUserHandler *kithttp.Server
	{
		signGetUserEndpoint := MakeSignGetUserByIDEndpoint(s)
		signGetUserEndpoint = endpoint.Chain(
			ValidateTokenMiddeware(s),
		)(signGetUserEndpoint)
		signGetUserOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		signGetUserHandler = kithttp.NewServer(
			ctx,
			signGetUserEndpoint,
			DecodeSignGetUserRequest,
			MakeResponseEncoder(http.StatusOK),
			signGetUserOptions...,
		)
	}

	var signLoginHandler *kithttp.Server
	{
		signLoginEndpoint := MakeSignLoginEndpoint(s)
		signLoginOptions := opts
		signLoginHandler = kithttp.NewServer(
			ctx,
			signLoginEndpoint,
			DecodeSignLoginRequest,
			MakeResponseEncoder(http.StatusOK),
			signLoginOptions...,
		)
	}

	var signLogoutHandler *kithttp.Server
	{
		signLogoutEndpoint := MakeSignLogoutEndpoint(s)
		signLogoutEndpoint = endpoint.Chain(
			ValidateTokenMiddeware(s),
		)(signLogoutEndpoint)
		signLogoutOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		signLogoutHandler = kithttp.NewServer(
			ctx,
			signLogoutEndpoint,
			DecodeSignLogoutRequest,
			MakeResponseEncoder(http.StatusNoContent),
			signLogoutOptions...,
		)
	}

	var signCreateAgentTokenHandler *kithttp.Server
	{
		signCreateAgentTokenEndpoint := MakeSignCreateAgentTokenEndpoint(s)
		signCreateAgentTokenEndpoint = endpoint.Chain(
			ValidateTokenMiddeware(s),
		)(signCreateAgentTokenEndpoint)
		signCreateAgentTokenOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		signCreateAgentTokenHandler = kithttp.NewServer(
			ctx,
			signCreateAgentTokenEndpoint,
			DecodeSignCreateAgentTokenRequest,
			MakeResponseEncoder(http.StatusOK),
			signCreateAgentTokenOptions...,
		)
	}

	var signDeleteAgentTokenHandler *kithttp.Server
	{
		signDeleteAgentTokenEndpoint := MakeSignDeleteAgentTokenEndpoint(s)
		signDeleteAgentTokenEndpoint = endpoint.Chain(
			ValidateTokenMiddeware(s),
		)(signDeleteAgentTokenEndpoint)
		signDeleteAgentTokenOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		signDeleteAgentTokenHandler = kithttp.NewServer(
			ctx,
			signDeleteAgentTokenEndpoint,
			DecodeSignDeleteAgentTokenRequest,
			MakeResponseEncoder(http.StatusNoContent),
			signDeleteAgentTokenOptions...,
		)
	}

	var signValidateTokenHandler *kithttp.Server
	{
		signValidateTokenEndpoint := MakeSignValidateTokenEndpoint(s)
		signValidateTokenOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		signValidateTokenHandler = kithttp.NewServer(
			ctx,
			signValidateTokenEndpoint,
			DecodeSignValidateTokenRequest,
			MakeResponseEncoder(http.StatusOK),
			signValidateTokenOptions...,
		)
	}

	r.Handle("/signing/v1/users", signCreateUserHandler).Methods("POST")
	r.Handle("/signing/v1/users/self", signValidateTokenHandler).Methods("GET")
	r.Handle("/signing/v1/users/self/access_tokens", signLoginHandler).Methods("POST")
	r.Handle("/signing/v1/users/{user_id}", signDeleteUserHandler).Methods("DELETE")
	r.Handle("/signing/v1/users/{user_id}", signGetUserHandler).Methods("GET")
	r.Handle("/signing/v1/users/{user_id}/access_tokens", signLogoutHandler).Methods("DELETE")
	r.Handle("/signing/v1/users/{user_id}/agent_tokens", signCreateAgentTokenHandler).Methods("POST")
	r.Handle("/signing/v1/users/{user_id}/agent_tokens/{name}", signDeleteAgentTokenHandler).Methods("DELETE")

	return r
}
