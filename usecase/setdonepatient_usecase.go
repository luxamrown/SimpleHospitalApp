package usecase

import (
	"enigmacamp.com/rumah_sakit/model"
	"enigmacamp.com/rumah_sakit/repository"
)

type SetDonePatientUseCase interface {
	SetDone(patientId string) (model.Patient, error)
}

type setDonePatientUseCase struct {
	patientRepo repository.PatientRepo
	clinicRepo  repository.ClinicRepo
}

func (s *setDonePatientUseCase) SetDone(patientId string) (model.Patient, error) {
	var clinicId string
	patient, err := s.patientRepo.PatientStatusSetDone(patientId)
	if err != nil {
		return model.Patient{}, err
	}
	clinicId = s.patientRepo.GetClinicIdByPatient(patientId)
	err = s.clinicRepo.RemoveQueue(clinicId)
	if err != nil {
		return model.Patient{}, err
	}
	return patient, nil
}

func NewSetDoneUseCase(patientRepo repository.PatientRepo, clinicRepo repository.ClinicRepo) SetDonePatientUseCase {
	return &setDonePatientUseCase{
		patientRepo: patientRepo,
		clinicRepo:  clinicRepo,
	}
}
