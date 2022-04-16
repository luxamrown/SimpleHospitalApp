package manager

import "enigmacamp.com/rumah_sakit/repository"

type RepoManager interface {
	PatientRepo() repository.PatientRepo
	ClinicRepo() repository.ClinicRepo
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) PatientRepo() repository.PatientRepo {
	return repository.NewPatientRepoImpl(r.infra.SqlDb())
}

func (r *repoManager) ClinicRepo() repository.ClinicRepo {
	return repository.NewClinicRepoImpl(r.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
