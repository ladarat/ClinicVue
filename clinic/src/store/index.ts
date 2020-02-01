import Vue from 'vue'
import Vuex, { StoreOptions } from 'vuex';
import patient from './patient.modules';
import { PATIENT_NAMESPACE } from './modules.namespace';

Vue.use(Vuex)

export interface RootState {
  version: string;
}

const store: StoreOptions<RootState> = {
  state: {
    version: '1.0.0'
  },
  modules: {
    'patient': patient
  }
};

export default new Vuex.Store<RootState>(store);
