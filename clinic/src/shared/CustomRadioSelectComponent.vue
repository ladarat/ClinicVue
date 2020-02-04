<template>
  <div>
    <label>{{ label }}</label>
    <template v-for="(item, index) in options">
      <input
        type="radio"
        name="name"
        :key="index"
        :value="item.value"
        v-model="newInput"
        @change="handleInput"
      />
      {{ item.text }}
    </template>
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from "vue-property-decorator";

@Component
export default class CustomRadioSelectComponent extends Vue {
  @Prop({ default: "label" })
  label!: string;

  @Prop({ default: "" })
  value!: string;

  @Prop()
  options!: Array<{ text: string; value: string }>;

  newInput: string = this.value;

  handleInput() {
    this.$emit("input", this.newInput);
  }
}
</script>
