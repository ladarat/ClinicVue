
interface CreatePatientRequest {
  firstname: string,
  lastname: string,
  nickname: string,
  sex: string,
  career: string,
  phone_number: string,
  current_address: string,
  work_address: string,
  required_documents: string,
  congenital_disease: string,
  drug_allergy: string,
  emergency_contact: string,
  relationship: string,
  citizen_id: string
}

export default CreatePatientRequest