<template>
  <div class="home">
    <div>{{ aaa }}</div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

export default class Home extends Vue {
  aaa: string = "a";
  test = () => {
    console.log(">>>");
  };
  mounted() {
    this.test();
    console.log("aaa");
  }
}
</script>
