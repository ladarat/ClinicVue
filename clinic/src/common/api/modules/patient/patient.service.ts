import ApiService from '@/common/api/api.service'
import CreatePatientRequest from './models/CreatePatientRequest'
import CreatePatientResponse from './models/CreatePatientResponse'

const PatientService = {
  createPatient(createPatientRequest: CreatePatientRequest) {
    return ApiService.post<CreatePatientResponse>('patient', createPatientRequest)
  }
}

export default PatientService