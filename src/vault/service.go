package vault

import (
	"golang.org/x/crypto/bcrypt"
	"context"
	"encoding/json"
	"net/http"
	"github.com/go-kit/kit/endpoint"
	"errors"
)
// Service provides password hashing capabilities
type Service interface {
	Hash(ctx context.Context, password string) (string, error)
	Validate (ctx context.Context, password, hash string) (bool, error)
}

type vaultService struct{}

func NewService() Service{
	return vaultService{}
}

func (vaultService) Hash(ctx context.Context, password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (vaultService) Validate(ctx context.Context, password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, nil
	}
	return true, nil
}

type hashRequest struct {
	Password string `json:"password"`
}

type hashResponse struct {
	Hash string `json:"hash"`
	Err string `json:"err,omitempty"`
}

// implement http.DecoderRequestFunc
func decodeHashRequest (ctx context.Context, r *http.Request) (interface{}, error) {
	var req hashRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

type validateRequest struct {
	Password string `json:"password"`
	Hash string `json:"hash"`
}

type validateResponse struct {
	Valid bool `json:"valid"`
	Err string `json:"err,omitempty"`
}

func decodeValidateRequest (ctx context.Context, r *http.Request) (interface{}, error) {
	var req validateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse (ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// Endoints represent a single RPC method. The definition is inside the endpoint package
// type Endpoint func (ctx context.Context, request interface{}) (response interface{}, error)

// Generate an endpoint from any Service implementation  that does hashing 
func MakeHashEndPoint (srv Service) endpoint.Endpoint {
	return  func(ctx context.Context, request interface{}) ( interface{},  error) {
		req := request.(hashRequest)
		v, err := srv.Hash(ctx, req.Password)
		if err != nil {
			return hashResponse{v, err.Error()}, nil
		}
		return hashResponse{v, ""}, nil
	}
}

// Generate an endpoint from any Service implementation  that does validating the hash
func MakeValidateEndPoint (srv Service) endpoint.Endpoint {
	return func (ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(validateRequest)
		v, err := srv.Validate(ctx, req.Password, req.Hash)
		if err != nil {
			return validateResponse {false, err.Error()}, nil
		}
		return validateResponse {v, ""}, nil
	}
}

// Wrapping endpoints into a Service implementation
type Endpoints struct {
	HashEndpoint endpoint.Endpoint
	ValidateEndpoint endpoint.Endpoint
}

// implement Service interface
func (e Endpoints) Hash(ctx context.Context, password string) (string, error) {
	req := hashRequest{Password: password}
	resp, err := e.HashEndpoint(ctx, req) // call server to handle RPC, later on we will supply all method to the server
	if err != nil {
		return "", err
	}
	hashResp := resp.(hashResponse)
	if hashResp.Err != "" {
		return "", errors.New(hashResp.Err)
	}
	return hashResp.Hash, nil
}
// implement Service interface
func (e Endpoints) Validate (ctx context.Context, password, hash string) (bool, error) {
	req := validateRequest{Password: password, Hash: hash}
	resp, err := e.ValidateEndpoint(ctx, req)
	if err != nil {
		return false, err
	}

	validateResp := resp.(validateResponse)
	if validateResp.Err != "" {
		return false, errors.New(validateResp.Err)
	}
	return validateResp.Valid, nil
}
