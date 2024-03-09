package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"fiber-orchestrator/structs"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"

	uuid "github.com/google/uuid"
)

func CreateUser(c *fiber.Ctx, userUuid uuid.UUID) ([]byte, int, error) {

	NewUser := new(structs.User)

	if err := c.BodyParser(NewUser); err != nil {
		return nil, http.StatusInternalServerError, errors.New("error parser userman")
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err1 := validate.Struct(NewUser)

	if err1 != nil {

		return nil, http.StatusInternalServerError, errors.New("error validate userman")

	}

	NewUser.ID = userUuid

	JsonUser, _ := json.Marshal(NewUser)

	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	viper.AutomaticEnv()

	host_auth, _ := viper.Get("HOST_AUTH").(string)

	auth_url := fmt.Sprintf("%s/auth/api/register", host_auth)

	req_user, err := http.NewRequest(http.MethodPost, auth_url, bytes.NewBuffer(JsonUser))

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("could not create request")
	}

	req_user.Header.Set("Content-Type", "application/json")

	res_user, err := http.DefaultClient.Do(req_user)

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("error making http request")
	}

	body_user, err := ioutil.ReadAll(res_user.Body)

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("error in read Body")
	}

	defer res_user.Body.Close()

	return body_user, res_user.StatusCode, nil
}
