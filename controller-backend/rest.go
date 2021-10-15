package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"

	"github.com/avssvd/ncr-test-golang/gen/rest/restapi"
	"github.com/avssvd/ncr-test-golang/gen/rest/restapi/operations"
	db "github.com/avssvd/ncr-test-golang/gen/sqlc"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

var ErrorDBBadResponse = errors.New("database error")
var ErrorControllerNotFound = errors.New("controller not found")
var ErrorControllerAlreadyExists = errors.New("controller already exists")

func restServe(serverPort int, db *db.Queries) {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewBackendAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = serverPort

	api.GetControllersHandler = operations.GetControllersHandlerFunc(GetControllersHandlerFunc(db))
	api.GetControllerIndicationsHandler = operations.GetControllerIndicationsHandlerFunc(GetControllerIndicationsHandlerFunc(db))
	api.PostControllerHandler = operations.PostControllerHandlerFunc(PostControllerHandlerFunc(db))
	api.DeleteControllerHandler = operations.DeleteControllerHandlerFunc(DeleteControllerHandlerFunc(db))

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}

func GetControllersHandlerFunc(db *db.Queries) func(operations.GetControllersParams) middleware.Responder {
	return func(params operations.GetControllersParams) middleware.Responder {
		controllers, err := db.ListControllers(context.Background())
		if err != nil {
			log.Println(err)
			return operations.NewGetControllersInternalServerError().WithPayload(&operations.GetControllersInternalServerErrorBody{Error: ErrorDBBadResponse.Error()})
		}

		body := operations.GetControllersOKBody{}
		for _, controller := range controllers {
			body.Controllers = append(body.Controllers, &operations.GetControllersOKBodyControllersItems0{
				CreatedAt: strfmt.DateTime(controller.CreatedAt),
				Serial:    controller.Serial,
			})
		}

		return operations.NewGetControllersOK().WithPayload(&body)
	}
}

func GetControllerIndicationsHandlerFunc(db *db.Queries) func(operations.GetControllerIndicationsParams) middleware.Responder {
	return func(params operations.GetControllerIndicationsParams) middleware.Responder {
		_, err := db.GetController(context.Background(), params.Controller.Serial)
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return operations.NewGetControllerIndicationsBadRequest().WithPayload(&operations.GetControllerIndicationsBadRequestBody{Error: ErrorControllerNotFound.Error()})
		case err != nil:
			log.Println(err)
			return operations.NewGetControllerIndicationsInternalServerError().WithPayload(&operations.GetControllerIndicationsInternalServerErrorBody{Error: ErrorDBBadResponse.Error()})
		}

		indications, err := db.ListIndicationsByController(context.Background(), params.Controller.Serial)
		if err != nil {
			log.Println(err)
			return operations.NewGetControllerIndicationsInternalServerError().WithPayload(&operations.GetControllerIndicationsInternalServerErrorBody{Error: ErrorDBBadResponse.Error()})
		}

		body := operations.GetControllerIndicationsOKBody{}
		for _, indication := range indications {
			temp, err := strconv.ParseFloat(indication.Indication, 32)
			if err != nil {
				log.Println(err)
			}
			body.Indications = append(body.Indications, &operations.GetControllerIndicationsOKBodyIndicationsItems0{
				Indication: float32(temp),
				SentAt:     strfmt.DateTime(indication.SentAt),
			})
		}

		return operations.NewGetControllerIndicationsOK().WithPayload(&body)
	}

}

func PostControllerHandlerFunc(db *db.Queries) func(operations.PostControllerParams) middleware.Responder {
	return func(params operations.PostControllerParams) middleware.Responder {
		_, err := db.GetController(context.Background(), params.Controller.Serial)
		switch {
		case err == nil:
			return operations.NewPostControllerBadRequest().WithPayload(&operations.PostControllerBadRequestBody{Error: ErrorControllerAlreadyExists.Error()})
		case !errors.Is(err, sql.ErrNoRows):
			log.Println(err)
			return operations.NewPostControllerInternalServerError().WithPayload(&operations.PostControllerInternalServerErrorBody{Error: ErrorDBBadResponse.Error()})
		}

		err = db.CreateController(context.Background(), params.Controller.Serial)
		if err != nil {
			log.Println(err)
			return operations.NewPostControllerInternalServerError().WithPayload(&operations.PostControllerInternalServerErrorBody{Error: ErrorDBBadResponse.Error()})
		}

		return operations.NewPostControllerOK().WithPayload(&operations.PostControllerOKBody{Success: true})
	}
}

func DeleteControllerHandlerFunc(db *db.Queries) func(operations.DeleteControllerParams) middleware.Responder {
	return func(params operations.DeleteControllerParams) middleware.Responder {
		affectedRows, err := db.DeleteController(context.Background(), params.Controller.Serial)
		if err != nil {
			log.Println(err)
			return operations.NewDeleteControllerInternalServerError().WithPayload(&operations.DeleteControllerInternalServerErrorBody{Error: ErrorDBBadResponse.Error()})
		}
		if affectedRows == 0 {
			return operations.NewDeleteControllerBadRequest().WithPayload(&operations.DeleteControllerBadRequestBody{Error: ErrorControllerNotFound.Error()})
		}
		return operations.NewDeleteControllerOK().WithPayload(&operations.DeleteControllerOKBody{Success: true})
	}
}
