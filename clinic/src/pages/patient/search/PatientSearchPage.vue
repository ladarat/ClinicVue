<template>
  <div class="patient-search container">
    <div class="text-center">ค้นหาผู้ป่วย</div>
    <hr />
    <div class="row">
      <div class="col-4">
        <div>ชื่อ-นามสกุล</div>
        <input v-model="searchFullName" type="text" />
      </div>
      <div class="col-4">
        <div>บัตรประชาชน</div>
        <input v-model="searchCitizenId" type="text" />
      </div>
      <div class="col-4">
        <button @click="searchPatient()" class="btn btn-primary">Search</button>
      </div>
    </div>
    <hr />
    <table class="table">
      <tr>
        <th>เลขบัตรประชาชน</th>
        <th>ชื่อ</th>
        <th>นามสกุล</th>
        <th>View</th>
      </tr>
      <template v-for="patient in getSearchedPatient">
        <tr class="justify-content-center" :key="patient.citizenId">
          <td>{{patient.citizenId}}</td>
          <td>{{patient.firstname}}</td>
          <td>{{patient.lastname}}</td>
          <td>
            <button @click="goToPatientDetail(patient)" class="btn btn-primary">View</button>
          </td>
        </tr>
      </template>
    </table>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';
import { Action, State, Getter, Mutation, namespace } from 'vuex-class';

import Patient from "../models/Patient";
import { PATIENT_NAMESPACE } from '../../../store/modules.namespace'

const PatientGetter = namespace(PATIENT_NAMESPACE, Getter);
const PatientAction = namespace(PATIENT_NAMESPACE, Action);

@Component
export default class PatientSearchPage extends Vue {
  @PatientAction searchPatient!: Function
  @PatientGetter getSearchedPatient!: Array<Patient>

  searchFullName: string = "";
  searchCitizenId: string = "";

  goToPatientDetail(patient: Patient) {
    this.$swal(patient.firstname);
  }

}
</script>