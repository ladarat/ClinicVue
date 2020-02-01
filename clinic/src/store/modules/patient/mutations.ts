import { MutationTree } from "vuex"
import { PatientState } from './state';
import { SET_PATIENT_LIST } from './mutations.type'
import Patient from '@/pages/patient/models/Patient';

export const mutations: MutationTree<PatientState> = {
  [SET_PATIENT_LIST](state, payload: Array<Patient>) {
      state.patientList = payload;
      state.error = false;
  }
};

export default mutations