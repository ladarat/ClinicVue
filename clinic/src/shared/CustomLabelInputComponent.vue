<template>
  <div>
    <label>{{ label }}</label>
    <input type="text" v-model="newInput" @input="handleInput"/>
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from "vue-property-decorator";

@Component
export default class CustomLabelInputComponent extends Vue {
  @Prop({ default: "label" })
  label!: string;

  @Prop({ default: '' })
  value!: string;

  newInput: string = this.value

  handleInput () {
      this.$emit('input', this.newInput)
  }
}
</script>
