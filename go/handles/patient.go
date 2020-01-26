package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	services "github.com/ladarat/ClinicVue/go/services"
)

// Now is time now
var Now = time.Now()

// PatientRequest struct 
type PatientRequest struct {
	ID                uint   `json:"id"`
	Firstname         string `json:"firstname"`
	Lastname          string `json:"lastname"`
	Nickname          string `json:"nickname"`
	Sex               string `json:"sex"`
	Career            string `json:"career"`
	PhoneNumber       string `json:"phone_number"`
	CurrentAddress    string `json:"current_address"`
	WorkAddress       string `json:"work_address"`
	RequiredDocuments string `json:"required_documents"`
	CongenitalDisease string `json:"congenital_disease"`
	DrugAllergy       string `json:"drug_allergy"`
	EmergencyContact  string `json:"emergency_contact"`
	Relationship      string `json:"relationship"`
}

// PatientResponse struct
type PatientResponse struct {
	ID                uint   `json:"id"`
	Firstname         string `json:"firstname"`
	Lastname          string `json:"lastname"`
	Nickname          string `json:"nickname"`
	Sex               string `json:"sex"`
	Career            string `json:"career"`
	PhoneNumber       string `json:"phone_number"`
	CurrentAddress    string `json:"current_address"`
	WorkAddress       string `json:"work_address"`
	RequiredDocuments string `json:"required_documents"`
	CongenitalDisease string `json:"congenital_disease"`
	DrugAllergy       string `json:"drug_allergy"`
	EmergencyContact  string `json:"emergency_contact"`
	Relationship      string `json:"relationship"`
}

// CreatePatient by POST /patient
func CreatePatient(ps services.PatientService) echo.HandlerFunc {
	return func(c echo.Context) error {
		version, err := toPatient(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		err = ps.Create(version)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		return c.JSON(http.StatusOK, toPatientJSON(version))
	}
}

func toPatient(c echo.Context) (*services.Patient, error) {
	var patientReq PatientRequest
	if err := c.Bind(&patientReq); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	var patient services.Patient

	patient.ID = patientReq.ID
	patient.Nickname = patientReq.Nickname
	patient.Firstname = patientReq.Firstname
	patient.Lastname = patientReq.Lastname
	patient.Sex = patientReq.Sex
	patient.Career = patientReq.Career
	patient.PhoneNumber = patientReq.PhoneNumber
	patient.CurrentAddress = patientReq.CurrentAddress
	patient.WorkAddress = patientReq.WorkAddress
	patient.RequiredDocuments = patientReq.RequiredDocuments
	patient.DrugAllergy = patientReq.DrugAllergy
	patient.EmergencyContact = patientReq.EmergencyContact
	patient.Relationship = patientReq.Relationship

	return &patient, nil
}

func toPatientJSON(p *services.Patient) PatientResponse {
	var json PatientResponse
	json.ID = p.ID
	json.Nickname = p.Nickname
	json.Firstname = p.Firstname
	json.Lastname = p.Lastname
	json.Sex = p.Sex
	json.Career = p.Career
	json.PhoneNumber = p.PhoneNumber
	json.CurrentAddress = p.CurrentAddress
	json.WorkAddress = p.WorkAddress
	json.RequiredDocuments = p.RequiredDocuments
	json.DrugAllergy = p.DrugAllergy
	json.EmergencyContact = p.EmergencyContact
	json.Relationship = p.Relationship
	return json
}
