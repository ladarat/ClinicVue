
import { Module, MutationTree, GetterTree } from 'vuex'
import actions from './actions'
import { RootState } from '../..'

import { SET_PATIENT_LIST } from './mutations.type'
import Patient from '@/pages/patient/models/Patient';

export interface PatientState {
  patientList: Array<Patient>,
  error: boolean
}

export const state: PatientState = {
  patientList: [],
  error: false
};


export const mutations: MutationTree<PatientState> = {
  [SET_PATIENT_LIST](state, payload: Array<Patient>) {
      state.patientList = payload;
      state.error = false;
  }
};

export const getters: GetterTree<PatientState, RootState> = {
  getSearchedPatient(state): Array<Patient> {
    return state.patientList;
  }
};

const namespaced: boolean = true;

export const patientStoreModule: Module<PatientState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations
};

export default patientStoreModule;
