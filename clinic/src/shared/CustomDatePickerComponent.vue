<template>
  <div>
    <label>{{ label }}</label>
    <input
      type="date"
      id="start"
      name="trip-start"
      min="1920-01-01"
      v-model="newInput"
      @input="handleInput"
    />
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from "vue-property-decorator";

@Component
export default class CustomDatePickerComponent extends Vue {
  @Prop({ default: "label" })
  label!: string;

  @Prop({ default: "" })
  value!: string;

  newInput: string = this.value;

  handleInput() {
    this.$emit("input", this.newInput);
  }
}
</script>
