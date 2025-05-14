package google

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/configs"
	"golang.org/x/oauth2"
	"io"
)

const userUrl = "https://www.googleapis.com/oauth2/v1/userinfo"

func (handler *HandlerGoogleAuth) login(c *fiber.Ctx) error {
	path := configs.ConfigGoogle()
	url := path.AuthCodeURL("state")

	return c.Redirect(url)
}

func (handler *HandlerGoogleAuth) callback(c *fiber.Ctx) error {
	code := c.FormValue("code")
	session, err := SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	token, err := configs.ConfigGoogle().Exchange(c.Context(), code)
	if err != nil {
		return err
	}

	user, err := handler.getUser(token)
	if err != nil {
		return err
	}

	jsonBytes, err := json.Marshal(&user)
	if err != nil {
		panic(err)
	}

	session.Set("user", string(jsonBytes))
	err = session.Save()
	if err != nil {
		return err
	}

	return c.Redirect("/auth/google/profile")
}

func (handler *HandlerGoogleAuth) getUser(token *oauth2.Token) (ResponseGoogle, error) {
	client := configs.ConfigGoogle().Client(context.Background(), token)

	response, err := client.Get(userUrl)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(response.Body)

	if err != nil {
		return ResponseGoogle{}, fmt.Errorf("%v", err)
	}

	var user ResponseGoogle

	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		return ResponseGoogle{}, fmt.Errorf("%v", err)
	}
	return user, nil
}
