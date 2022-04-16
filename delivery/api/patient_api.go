package api

import (
	"net/http"

	"enigmacamp.com/rumah_sakit/delivery/appreq"
	"enigmacamp.com/rumah_sakit/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PatientApi struct {
	addPatientUseCase usecase.AddPatientUseCase
	setDoneUsecase    usecase.SetDonePatientUseCase
}

func (p *PatientApi) AddPatient() gin.HandlerFunc {
	return func(c *gin.Context) {
		var patientReq appreq.PatientRequest
		patientId := uuid.New().String()
		if err := c.ShouldBindJSON(&patientReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot bind json"})
		}
		newPatient, err := p.addPatientUseCase.Insert(patientId, patientReq.PatientName, patientReq.EarlySymptoms, patientReq.ClinicId, false)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot add patient"})
		}
		c.JSON(http.StatusOK, newPatient)
	}
}

func (p *PatientApi) SetDone() gin.HandlerFunc {
	return func(c *gin.Context) {
		var patientReq appreq.PatientRequestDone
		if err := c.ShouldBindJSON(&patientReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot bind json"})
		}
		patient, err := p.setDoneUsecase.SetDone(patientReq.PatientId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot change status"})
		}
		c.JSON(http.StatusOK, patient)
	}
}

func NewPatientApi(patientRoute *gin.RouterGroup, addPatientUsecase usecase.AddPatientUseCase, setDonePatientUseCase usecase.SetDonePatientUseCase) {
	api := PatientApi{
		addPatientUseCase: addPatientUsecase,
		setDoneUsecase:    setDonePatientUseCase,
	}
	patientRoute.POST("/add", api.AddPatient())
	patientRoute.POST("/setdone", api.SetDone())
}
