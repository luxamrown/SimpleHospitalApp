package appreq

type PatientRequest struct {
	PatientName   string `json:"patient_name"`
	EarlySymptoms string `json:"early_symptoms"`
	ClinicId      string `json:"clinic_id"`
}

type PatientRequestDone struct {
	PatientId string `json:"patient_id"`
}
