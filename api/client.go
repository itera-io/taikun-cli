package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"taikun-cli/utils"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/itera-io/taikungoclient/client"
	"github.com/itera-io/taikungoclient/client/auth"
	"github.com/itera-io/taikungoclient/client/keycloak"
	"github.com/itera-io/taikungoclient/models"
)

const TaikunEmailEnvVar = "TAIKUN_EMAIL"
const TaikunPasswordEnvVar = "TAIKUN_PASSWORD"
const TaikunKeycloakEmailEnvVar = "TAIKUN_KEYCLOAK_EMAIL"
const TaikunKeycloakPasswordEnvVar = "TAIKUN_KEYCLOAK_PASSWORD"
const TaikunApiHostEnvVar = "TAIKUN_API_HOST"

type Client struct {
	Client *client.Taikungoclient

	email               string
	password            string
	useKeycloakEndpoint bool

	token        string
	refreshToken string
}

func NewClient() (*Client, error) {
	email, keycloakEnabled := os.LookupEnv(TaikunKeycloakEmailEnvVar)
	password := os.Getenv(TaikunKeycloakPasswordEnvVar)

	if !keycloakEnabled {
		email = os.Getenv(TaikunEmailEnvVar)
		password = os.Getenv(TaikunPasswordEnvVar)
	}

	if email == "" || password == "" {
		return nil, fmt.Errorf(
			`Please set your Taikun credentials.
To authenticate with your Taikun account, set the following environment variables:
%s
%s

To authenticate with Keycloak, set the following environment variables:
%s
%s

To override the default API host, set %s.`,
			TaikunKeycloakEmailEnvVar,
			TaikunKeycloakPasswordEnvVar,
			TaikunEmailEnvVar,
			TaikunPasswordEnvVar,
			TaikunApiHostEnvVar,
		)
	}

	transportConfig := client.DefaultTransportConfig()
	if apiHost, apiHostIsSet := os.LookupEnv(TaikunApiHostEnvVar); apiHostIsSet {
		transportConfig = transportConfig.WithHost(apiHost)
	}

	return &Client{
		Client:              client.NewHTTPClientWithConfig(nil, transportConfig),
		email:               email,
		password:            password,
		useKeycloakEndpoint: keycloakEnabled,
	}, nil
}

type jwtData struct {
	Nameid     string `json:"nameid"`
	Email      string `json:"email"`
	UniqueName string `json:"unique_name"`
	Role       string `json:"role"`
	Nbf        int    `json:"nbf"`
	Exp        int    `json:"exp"`
	Iat        int    `json:"iat"`
}

func (apiClient *Client) AuthenticateRequest(c runtime.ClientRequest, _ strfmt.Registry) error {

	if len(apiClient.token) == 0 {

		if !apiClient.useKeycloakEndpoint {
			loginResult, err := apiClient.Client.Auth.AuthLogin(
				auth.NewAuthLoginParams().WithV(utils.ApiVersion).WithBody(
					&models.LoginCommand{Email: apiClient.email, Password: apiClient.password},
				), nil,
			)
			if err != nil {
				return err
			}
			apiClient.token = loginResult.Payload.Token
			apiClient.refreshToken = loginResult.Payload.RefreshToken
		} else {
			loginResult, err := apiClient.Client.Keycloak.KeycloakLogin(
				keycloak.NewKeycloakLoginParams().WithV(utils.ApiVersion).WithBody(
					&models.LoginWithKeycloakCommand{Email: apiClient.email, Password: apiClient.password},
				), nil,
			)
			if err != nil {
				return err
			}
			apiClient.token = loginResult.Payload.Token
			apiClient.refreshToken = loginResult.Payload.RefreshToken
		}

	}

	if apiClient.hasTokenExpired() {

		refreshResult, err := apiClient.Client.Auth.AuthRefreshToken(
			auth.NewAuthRefreshTokenParams().WithV(utils.ApiVersion).WithBody(
				&models.RefreshTokenCommand{
					RefreshToken: apiClient.refreshToken,
					Token:        apiClient.token,
				}), nil,
		)
		if err != nil {
			return err
		}

		apiClient.token = refreshResult.Payload.Token
		apiClient.refreshToken = refreshResult.Payload.RefreshToken
	}

	err := c.SetHeaderParam("Authorization", fmt.Sprintf("Bearer %s", apiClient.token))
	if err != nil {
		return err
	}

	return nil
}

func (apiClient *Client) hasTokenExpired() bool {
	jwtSplit := strings.Split(apiClient.token, ".")
	if len(jwtSplit) != 3 {
		return true
	}

	data, err := base64.RawURLEncoding.DecodeString(jwtSplit[1])
	if err != nil {
		return true
	}

	jwtData := jwtData{}
	err = json.Unmarshal(data, &jwtData)
	if err != nil {
		return true
	}

	tm := time.Unix(int64(jwtData.Exp), 0)

	return tm.Before(time.Now())
}
