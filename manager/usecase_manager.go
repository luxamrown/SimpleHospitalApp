package manager

import (
	"enigmacamp.com/rumah_sakit/usecase"
)

type UseCaseManager interface {
	AddPatientUseCase() usecase.AddPatientUseCase
	SetDonePatientUseCase() usecase.SetDonePatientUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) AddPatientUseCase() usecase.AddPatientUseCase {
	return usecase.NewAddPatientUseCase(u.repo.PatientRepo(), u.repo.ClinicRepo())
}

func (u *useCaseManager) SetDonePatientUseCase() usecase.SetDonePatientUseCase {
	return usecase.NewSetDoneUseCase(u.repo.PatientRepo(), u.repo.ClinicRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repo: repo,
	}
}
