package handlers

import (
	infra "api/infra/pg"
	"api/repositories"
	"api/types"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetPeople(ctx *gin.Context) {

	var person types.GetPerson
	err := ctx.ShouldBindUri(&person)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
	}

	log.Printf(person.ID)
	ctx.Value(person.ID)
}

func CreatePerson(ctx *gin.Context) {
	var person types.PostPerson

	err := parseInput(ctx, &person)

	if err != nil {
		log.Printf("err %s", err.Error())
		ctx.Status(400)
		return
	}

	error := repositories.CreatePersonRepository(infra.DbClient, person)

	if error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, error)
		return
	}

	ctx.Status(201)
}

func parseInput(ctx *gin.Context, person *types.PostPerson) error {

	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(&person)

	if err != nil {
		return err
	}

	validate := validator.New()
	errValidation := validate.Struct(person)

	if errValidation != nil {
		return errValidation
	}

	return nil

}
