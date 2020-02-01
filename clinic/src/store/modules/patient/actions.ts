
import { ActionTree } from 'vuex';
import { RootState } from '../..';
import { PatientState } from './state';
import { SET_PATIENT_LIST } from './mutations.type';
import Patient from '@/pages/patient/models/Patient';


export const actions: ActionTree<PatientState, RootState> = {
  searchPatient({ commit }) {
    const patient1: Patient = {
      citizenId: '11111111111',
      firstname: 'boom',
      lastname: 'boom'
    };
    commit(SET_PATIENT_LIST, [patient1])
  }
};

export default actions