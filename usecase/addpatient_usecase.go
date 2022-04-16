package usecase

import (
	"enigmacamp.com/rumah_sakit/model"
	"enigmacamp.com/rumah_sakit/repository"
)

type AddPatientUseCase interface {
	Insert(patientId, patientName, patientClinic, earlySymptoms string, isDone bool) (model.Patient, error)
}

type addPatientUseCase struct {
	patientRepo repository.PatientRepo
	clinicRepo  repository.ClinicRepo
}

func (a *addPatientUseCase) Insert(patientId, patientName, earlySymptoms, patientClinic string, isDone bool) (model.Patient, error) {
	newPatient := model.NewPatient(patientId, patientName, earlySymptoms, patientClinic, isDone)
	patient, err := a.patientRepo.InsertPatient(newPatient)
	if err != nil {
		return model.Patient{}, err
	}
	err = a.clinicRepo.AddQueue(patientClinic)
	if err != nil {
		return model.Patient{}, err
	}
	return patient, nil
}

func NewAddPatientUseCase(patientRepo repository.PatientRepo, clinicRepo repository.ClinicRepo) AddPatientUseCase {
	return &addPatientUseCase{
		patientRepo: patientRepo,
		clinicRepo:  clinicRepo,
	}
}
