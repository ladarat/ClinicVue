<template>
    <div>
        Create Patient

        <!-- <custom-label-input label="ชื่อ-นามสกุล" v-model="name"/> -->
        <custom-label-input label="ชื่อ-นามสกุล" @input="handleInput2"/>
        {{ name }}
    </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from "vue-property-decorator";
import CustomLabelInputComponent from '@/shared/CustomLabelInputComponent.vue';

@Component({components: {'custom-label-input': CustomLabelInputComponent}})
export default class CreatePatientPage extends Vue {
  name: string = 'firstname'
  handleInput2(e) {
      this.name = e
  }
}
</script>