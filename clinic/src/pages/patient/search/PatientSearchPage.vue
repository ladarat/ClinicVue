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
        <button @click="search(searchPatient)" class="btn btn-primary">Search</button>
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
      <template v-for="patient in patientList">
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
import { Vue, Component, Prop } from "vue-property-decorator";
import Patient from "../models/Patient";

@Component
export default class PatientSearchPage extends Vue {
  patientList: Array<Patient> = [];
  searchFullName: string = ''
  searchCitizenId: string = ''

  get searchPatient(): Patient {
    const fullNameSplited = this.searchFullName.split(" ")
    const patient: Patient = {
      firstname: fullNameSplited[0],
      lastname: fullNameSplited[1],
      citizenId: this.searchCitizenId
    }
    return patient
  }

  mockPatient() {
    const patient1: Patient = {
      citizenId: "11111111111",
      firstname: "boom",
      lastname: "boom"
    };
    this.patientList.push(patient1);
    const patient2: Patient = {
      citizenId: "22222222222",
      firstname: "bas",
      lastname: "bas"
    };
    this.patientList.push(patient2);
    const patient3: Patient = {
      citizenId: "33333333333",
      firstname: "boss",
      lastname: "boss"
    };
    this.patientList.push(patient3);
  }

  goToPatientDetail(patient: Patient) {
    this.$swal(patient.firstname);
  }

  search(patient: Patient) {
    this.mockPatient()
  }
}
</script>