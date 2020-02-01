
import { Module } from 'vuex'
import { state, PatientState } from './state'
import { actions } from './actions'
import { mutations } from './mutations'
import { getters } from './getters'
import { RootState } from '../..'


const namespaced: boolean = true;

export const patientStoreModule: Module<PatientState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations
};

export default patientStoreModule;
