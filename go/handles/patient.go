package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	services "github.com/ladarat/ClinicVue/go/services"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Now is time now
var Now = time.Now()

// PatientRequest struct
type PatientRequest struct {
	ID                primitive.ObjectID `json:"id"`
	Firstname         string             `json:"firstname"`
	Lastname          string             `json:"lastname"`
	Nickname          string             `json:"nickname"`
	Sex               string             `json:"sex"`
	Career            string             `json:"career"`
	PhoneNumber       string             `json:"phone_number"`
	CurrentAddress    string             `json:"current_address"`
	WorkAddress       string             `json:"work_address"`
	RequiredDocuments string             `json:"required_documents"`
	CongenitalDisease string             `json:"congenital_disease"`
	DrugAllergy       string             `json:"drug_allergy"`
	EmergencyContact  string             `json:"emergency_contact"`
	Relationship      string             `json:"relationship"`
}

// PatientResponse struct
type PatientResponse struct {
	ID                primitive.ObjectID `json:"id"`
	Firstname         string             `json:"firstname"`
	Lastname          string             `json:"lastname"`
	Nickname          string             `json:"nickname"`
	Sex               string             `json:"sex"`
	Career            string             `json:"career"`
	PhoneNumber       string             `json:"phone_number"`
	CurrentAddress    string             `json:"current_address"`
	WorkAddress       string             `json:"work_address"`
	RequiredDocuments string             `json:"required_documents"`
	CongenitalDisease string             `json:"congenital_disease"`
	DrugAllergy       string             `json:"drug_allergy"`
	EmergencyContact  string             `json:"emergency_contact"`
	Relationship      string             `json:"relationship"`
	CreatedAt         time.Time          `json:"created_at,omitempty"`
	UpdatedAt         time.Time          `json:"updated_at,omitempty"`
}

// CreatePatient by POST /patient
func CreatePatient(ps services.PatientService) echo.HandlerFunc {
	return func(c echo.Context) error {
		patient, err := toPatient(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusExpectationFailed, err.Error())
		}
		_, err = ps.Create(patient)
		if err != nil {
			return echo.NewHTTPError(http.StatusExpectationFailed, err.Error())
		}

		return c.JSON(http.StatusCreated, toPatientJSON(patient))
	}
}

// GetPatientAll by GET /patient
func GetPatientAll(ps services.PatientService) echo.HandlerFunc {
	return func(c echo.Context) error {
		patients, err := ps.GetAll()
		if err != nil {
			return echo.NewHTTPError(http.StatusExpectationFailed, err.Error())
		}

		jsons := []PatientResponse{}
		for _, p := range patients {
			jsons = append(jsons, toPatientJSON(&p))
		}

		return c.JSON(http.StatusOK, jsons)
	}
}

// GetPatientByID by GET /patient/:id
func GetPatientByID(ps services.PatientService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		patient, err := ps.GetByID(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusExpectationFailed, err.Error())
		}
		return c.JSON(http.StatusOK, toPatientJSON(patient))
	}
}

// UpdatePatient by PUT /patient
func UpdatePatient(ps services.PatientService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		patient, err := toPatient(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusExpectationFailed, err.Error())
		}

		p, err := ps.Update(patient, id)
		if err != nil {
			return echo.NewHTTPError(http.StatusExpectationFailed, err.Error())
		}

		return c.JSON(http.StatusOK, toPatientJSON(p))
	}
}

// DeletePatient by DELETE /patient/:id
func DeletePatient(ps services.PatientService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		err := ps.Delete(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusExpectationFailed, err.Error())
		}
		return c.NoContent(http.StatusNoContent)
	}
}

func toPatient(c echo.Context) (*services.Patient, error) {
	var patientReq PatientRequest
	if err := c.Bind(&patientReq); err != nil {
		return nil, echo.NewHTTPError(http.StatusExpectationFailed, err.Error())
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
	json := PatientResponse{}
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
	json.CreatedAt = p.CreatedAt
	json.UpdatedAt = p.UpdatedAt
	return json
}
