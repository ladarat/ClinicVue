<template>
  <div>
    Create Patient
    <custom-label-input label="ชื่อ-นามสกุล" v-model="name" />
    <custom-select label="เพศ" :options="genderOptions" v-model="gender" />
    {{ name }}
    <br>
    {{ gender }}
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