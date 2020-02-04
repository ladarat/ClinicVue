<template>
  <div class="container mx-auto py-4">
    <div class="text-center mb-4">Create Patient</div>
    <custom-label-input label="ชื่อ-นามสกุล" v-model="name" />
    <custom-select label="เพศ" :options="genderOptions" v-model="gender" />
    {{ name }}
    <br />
    {{ gender }}
    <div class="row px-3 mt-2">
      <div class="col-12 mb-2">
        <custom-label-input label="ชื่อ-นามสกุล" v-model="name" />
      </div>
      <div class="col-5 mb-2">
        <custom-label-input label="ชื่อเล่น" v-model="name" />
      </div>
      <div class="col-7 mb-2">
        <custom-select label="เพศ" :options="genderOptions" v-model="gender" />
      </div>
      <div class="col-12 mb-2">วันเกิด</div>
      <div class="col-8 mb-2">
        <custom-label-input label="บัตรประชาชนลขที่" v-model="name" />
      </div>
      <div class="col-8 mb-2">
        <custom-label-input label="อาชีพ" v-model="name" />
      </div>
      <div class="col-8 mb-2">
        <custom-label-input label="โทร" v-model="name" />
      </div>
      <div class="col-12 mb-2">
        <custom-label-input label="ที่อยู่ปัจจุบัน" v-model="name" />
      </div>
      <div class="col-12 mb-2">
        <custom-label-input label="ที่อยู่ที่ทำงาน" v-model="name" />
      </div>
      <div class="col-12 mb-2">เอกสารที่ต้องการ</div>
      <div class="col-12 mb-2">
        <custom-label-input label="โรคประจำตัว" v-model="name" />
      </div>
      <div class="col-12 mb-2">
        <custom-label-input label="แพ้ยา" v-model="name" />
      </div>
    </div>
    <hr />
    <div class="text-center mb-4">ติดต่อฉุกเฉิน</div>
    <div class="row px-3">
      <div class="col-12 mb-2">
        <custom-label-input label="ผู้ติดต่อฉุกเฉิน" v-model="name" />
      </div>
      <div class="col-8 mb-2">
        <custom-label-input label="ความสัมพันธ์" v-model="name" />
      </div>
      <div class="col-8 mb-2">
        <custom-label-input label="โทร" v-model="name" />
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

@Component({
  components: {
    "custom-label-input": CustomLabelInputComponent,
    "custom-select": CustomSelectComponent
  }
})
export default class CreatePatientPage extends Vue {
  name: string = "firstname";
  gender: string = "";

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
}
</script>