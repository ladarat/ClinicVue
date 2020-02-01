import Vue from 'vue'
import Vuex, { StoreOptions } from 'vuex';
import patient from './modules/patient';

Vue.use(Vuex)

export interface RootState {
  version: string;
}

const store: StoreOptions<RootState> = {
  state: {
    version: '1.0.0'
  },
  modules: {
    patient
  }
};

export default new Vuex.Store<RootState>(store);
