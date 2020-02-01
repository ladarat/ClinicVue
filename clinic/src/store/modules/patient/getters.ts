import { GetterTree } from 'vuex'
import { PatientState } from './state'
import { RootState } from '../../'
import { Patient } from '@/pages/patient/models/Patient'

export const getters: GetterTree<PatientState, RootState> = {
  getSearchedPatient(state): Array<Patient> {
    return state.patientList
  }
};

