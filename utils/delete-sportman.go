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

func DeleteSportmen(c *fiber.Ctx, userUuid uuid.UUID) ([]byte, int, error) {

	NewSportmen := new(structs.SportmenDelete)

	if err := c.BodyParser(NewSportmen); err != nil {
		return nil, http.StatusInternalServerError, errors.New("error parser sportman")
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err1 := validate.Struct(NewSportmen)

	if err1 != nil {

		return nil, http.StatusInternalServerError, errors.New("error validate sportman")

	}

	NewSportmen.UserId = userUuid

	JsonSportmen, _ := json.Marshal(NewSportmen)

	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	viper.AutomaticEnv()

	host_sport, _ := viper.Get("HOST_SPORT").(string)

	sport_url := fmt.Sprintf("%s/sport", host_sport)

	req_sport, err := http.NewRequest(http.MethodDelete, sport_url, bytes.NewBuffer(JsonSportmen))

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("could not create request")
	}

	req_sport.Header.Set("Content-Type", "application/json")

	res_sport, err := http.DefaultClient.Do(req_sport)

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("error making http request")
	}

	body_sport, err := ioutil.ReadAll(res_sport.Body)

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("error in read Body")
	}

	defer res_sport.Body.Close()

	return body_sport, res_sport.StatusCode, nil
}
