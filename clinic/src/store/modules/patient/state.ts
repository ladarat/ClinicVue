import Patient from '@/pages/patient/models/Patient';

export interface PatientState {
  patientList: Array<Patient>,
  error: boolean
}

export const state: PatientState = {
  patientList: [],
  error: false
};