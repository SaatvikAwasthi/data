package handler

import (
	"log"

	"github.com/gin-gonic/gin"

	"tester/app/api/contract"
	"tester/app/operation"
	"tester/app/operation/factory"
	query2 "tester/app/operation/query"
)

const (
	idQuery        = "id"
	createdAtQuery = "createdAt"
)

type dataHandler struct {
	operation contract.OperationHandlerFactory
}

func NewDataHandler(op contract.OperationHandlerFactory) *dataHandler {
	return &dataHandler{
		operation: op,
	}
}

func (dh *dataHandler) GetData(c *gin.Context) {
	if id, exists := c.GetQuery(idQuery); exists {
		dh.getDataById(c, id)
		return
	}
	if createdAt, exists := c.GetQuery(createdAtQuery); exists {
		dh.getDataByCreatedAt(c, createdAt)
		return
	}
	dh.getAllData(c)
}

func (dh *dataHandler) AddData(c *gin.Context) {
	var data operation.AddRawDataResponse
	var err error

	data, err = dh.operation.CommandHandler(factory.AddRawDataCommandHandler).(contract.AddRawData).Handle(c.Request.Context())
	if err != nil {
		log.Printf("Error adding data: %v\n", err)
		c.JSON(500, gin.H{"error": "Failed to add data"})
		return
	}

	c.JSON(201, data)
}

func (dh *dataHandler) getAllData(c *gin.Context) {
	var data operation.RawDataResponse
	var err error

	data, err = dh.operation.QueryHandler(factory.GetAllRawDataQueryHandler).(contract.GetAllRawData).Handle(c.Request.Context())
	if err != nil {
		log.Printf("Error fetching data: %v\n", err)
		c.JSON(500, gin.H{"error": "Failed to fetch data"})
		return
	}

	c.JSON(200, data)
}

func (dh *dataHandler) getDataById(c *gin.Context, id string) {
	var data operation.RawDataResponse
	var err error

	req := query2.GetRawDataQueryRequest{
		Key:   query2.Id,
		Value: id,
	}

	data, err = dh.operation.QueryHandler(factory.GetRawDataQueryHandler).(query2.GetRawData).Handle(c.Request.Context(), req)
	if err != nil {
		log.Printf("Error fetching data: %v\n", err)
		c.JSON(500, gin.H{"error": "Failed to fetch data"})
		return
	}

	c.JSON(200, data)
}

func (dh *dataHandler) getDataByCreatedAt(c *gin.Context, createdAt string) {
	var data operation.RawDataResponse
	var err error

	req := query2.GetRawDataQueryRequest{
		Key:   query2.CreatedAt,
		Value: createdAt,
	}

	data, err = dh.operation.QueryHandler(factory.GetRawDataQueryHandler).(query2.GetRawData).Handle(c.Request.Context(), req)
	if err != nil {
		log.Printf("Error fetching data: %v\n", err)
		c.JSON(500, gin.H{"error": "Failed to fetch data"})
		return
	}

	c.JSON(200, data)
}
