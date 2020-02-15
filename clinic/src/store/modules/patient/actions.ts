
import { ActionTree } from 'vuex';
import { RootState } from '../..';
import { PatientState } from './state';
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
    // const patient1: Patient = {
    //   citizenId: '11111111111',
    //   firstname: 'boom',
    //   lastname: 'boom'
    // };
    // commit(SET_PATIENT_LIST, [patient1])
  },
  async createPatient(context: any, patient: Patient) {
    const createPatientRequest = patientToCreatePatientRequest(patient)
    await PatientService.createPatient(createPatientRequest)
  }
};

export default actions