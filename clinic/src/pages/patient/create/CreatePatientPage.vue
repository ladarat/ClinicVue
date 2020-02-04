<template>
  <div>
    Create Patient
    <custom-label-input label="ชื่อ-นามสกุล" v-model="name" />
    <custom-select label="เพศ" :options="genderOptions" v-model="gender" />
    {{ name }}
    <br>
    {{ gender }}
    <div class="row">
      <div class="col-12">
        <custom-label-input label="ชื่อ-นามสกุล" v-model="name" />
      </div>
      <div class="col-6">
        <custom-label-input label="ชื่อเล่น" v-model="name" />
      </div>
      <div class="col-6">
        <custom-select label="เพศ" :options="genderOptions" v-model="gender" />
      </div>
      <div class="col-12">
        วันเกิด
      </div>
      <div class="col-8">
        <custom-label-input label="บัตรประชาชนลขที่" v-model="name" />
      </div>
      <div class="col-8">
        <custom-label-input label="อาชีพ" v-model="name" />
      </div>
      <div class="col-8">
        <custom-label-input label="โทร" v-model="name" />
      </div>
      <div class="col-12">
        <custom-label-input label="ที่อยู่ปัจจุบัน" v-model="name" />
      </div>
      <div class="col-12">
        <custom-label-input label="ที่อยู่ที่ทำงาน" v-model="name" />
      </div>
      <div class="col-12">
        เอกสารที่ต้องการ
      </div>
      <div class="col-12">
        <custom-label-input label="โรคประจำตัว" v-model="name" />
      </div>
      <div class="col-12">
        <custom-label-input label="แพ้ยา" v-model="name" />
      </div>
    </div>
   <hr>
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