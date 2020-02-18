
import { ActionTree } from 'vuex';
import { RootState } from '../..';
import { PatientState } from './state';
import { CREATE_PATIENT } from './action.type';
import { SET_PATIENT_LIST } from './mutations.type';
import Patient from '@/pages/patient/models/Patient';

import PatientService from '@/common/api/modules/patient/patient.service'
import CreatePatientRequest from '@/common/api/modules/patient/models/CreatePatientRequest'

const patientToCreatePatientRequest = (patient: Patient): CreatePatientRequest => {
  return {
    firstname: patient.firstname,
    lastname: patient.lastname,
    nickname: patient.nickname,
    sex: patient.sex,
    career: patient.carreer,
    phone_number: patient.phoneNumber,
    current_address: patient.currentAddress,
    work_address: patient.workAddress,
    required_documents: patient.requiredDocuments,
    congenital_disease: patient.congenitalDisease,
    drug_allergy: patient.drugAllergy,
    emergency_contact: patient.emergencyContact,
    relationship: patient.relationship,
    citizen_id: patient.citizenId
  }
}

export const actions: ActionTree<PatientState, RootState> = {
  searchPatient({ commit }) {
    // commit(SET_PATIENT_LIST, [patient1])
  },
  [CREATE_PATIENT](context: any, patient: Patient) {
    const createPatientRequest = patientToCreatePatientRequest(patient)
    return new Promise((resolve) => {
      PatientService.createPatient(createPatientRequest)
        .then((res) => {
          resolve(res.data)
        })
        .catch((err) => {
          console.log(err)
        })
    })
  }
};

export default actions