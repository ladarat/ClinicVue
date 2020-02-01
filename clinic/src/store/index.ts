import Vue from 'vue'
import Vuex, { StoreOptions } from 'vuex';
import patient from './modules/patient';
import auth from './auth.module'

Vue.use(Vuex)

export interface RootState {
  version: string;
}

const store: StoreOptions<RootState> = {
  state: {
    version: '1.0.0'
  },
  modules: {
    patient,
    auth
  }
};

export default new Vuex.Store<RootState>(store);
