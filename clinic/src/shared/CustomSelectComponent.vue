<template>
  <div>
    <label>{{ label }}</label>
    <select v-model="newInput" @change="handleInput">
      <option v-for="(item,index) in options" :key="index" :value="item.value">
          {{item.text}}
      </option>
    </select>
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from "vue-property-decorator";

@Component
export default class CustomSelectComponent extends Vue {
  @Prop({ default: "label" })
  label!: string;

  @Prop({ default: "" })
  value!: string;

  @Prop()
  options!: Array<{text: string, value: string}>

  newInput: string = this.value

  handleInput () {
      this.$emit('input', this.newInput)
  }
}
</script>