<template>
  <div class="container mx-auto py-4">
    <div class="text-center mb-4">Create Patient</div>
    <div class="row px-3 mt-2">
      <div class="col-12 mb-2">
        <custom-label-input label="ชื่อ-นามสกุล" v-model="name" />
      </div>
      <div class="col-5 mb-2">
        <custom-label-input label="ชื่อเล่น" v-model="nickname" />
      </div>
      <div class="col-7 mb-2">
        <custom-select label="เพศ" :options="genderOptions" v-model="gender" />
      </div>
      <div class="col-12 mb-2">
        <custom-date-picker label="วันเกิด" v-model="dateOfBirth" />
      </div>
      <div class="col-8 mb-2">
        <custom-label-input label="บัตรประชาชนลขที่" v-model="citizenId" />
      </div>
      <div class="col-8 mb-2">
        <custom-label-input label="อาชีพ" v-model="job" />
      </div>
      <div class="col-8 mb-2">
        <custom-label-input label="โทร" v-model="phoneNumber" />
      </div>
      <div class="col-12 mb-2">
        <custom-label-input label="ที่อยู่ปัจจุบัน" v-model="currentLocation" />
      </div>
      <div class="col-12 mb-2">
        <custom-label-input label="ที่อยู่ที่ทำงาน" v-model="officeLocation" />
      </div>
      <div class="col-12 mb-2">
        <custom-radio-select label="เอกสารที่ต้องการ" :options="documentOptions" v-model="confirmDocumentType" />
      </div>
      <div class="col-12 mb-2">
        <custom-label-input label="โรคประจำตัว" v-model="congenitalDisease" />
      </div>
      <div class="col-12 mb-2">
        <custom-label-input label="แพ้ยา" v-model="allergicMedicine" />
      </div>
    </div>
    <hr />
    <div class="text-center mb-4">ติดต่อฉุกเฉิน</div>
    <div class="row px-3">
      <div class="col-12 mb-2">
        <custom-label-input label="ผู้ติดต่อฉุกเฉิน" v-model="urgentContactName" />
      </div>
      <div class="col-8 mb-2">
        <custom-label-input label="ความสัมพันธ์" v-model="urgentContactRelation" />
      </div>
      <div class="col-8 mb-2">
        <custom-label-input label="โทร" v-model="urgentContactPhoneNumber" />
      </div>
    </div>
    <div class="text-center">
      <button class="btn btn-primary"> เพิ่ม </button>
    </div>
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from "vue-property-decorator";
import CustomLabelInputComponent from "@/shared/CustomLabelInputComponent.vue";
import CustomSelectComponent from "@/shared/CustomSelectComponent.vue";
import CustomDatePickerComponent from "@/shared/CustomDatePickerComponent.vue";
import CustomRadioSelectComponent from "@/shared/CustomRadioSelectComponent.vue";

@Component({
  components: {
    "custom-label-input": CustomLabelInputComponent,
    "custom-select": CustomSelectComponent,
    "custom-date-picker": CustomDatePickerComponent,
    "custom-radio-select": CustomRadioSelectComponent
  }
})
export default class CreatePatientPage extends Vue {
  name: string = ""
  nickname: string = ""
  gender: string = "male"
  dateOfBirth: string = ""
  citizenId: string = ""
  job: string = ""
  phoneNumber: string = ""
  currentLocation: string = ""
  officeLocation: string = ""
  confirmDocumentType: string = ""
  allergicMedicine: string = ""
  congenitalDisease: string = ""
  urgentContactName: string = ""
  urgentContactRelation: string = ""
  urgentContactPhoneNumber: string = ""

  genderOptions: Array<{ text: string; value: string }> = [
    {
      text: "male",
      value: "male"
    },
    {
      text: "female",
      value: "female"
    }
  ];

  documentOptions: Array<{ text: string; value: string }> = [
    {
      text: "ใบรับรองแพทย์",
      value: "1"
    },
    {
      text: "ใบประกันสังคม",
      value: "2"
    }
  ];
}
</script>