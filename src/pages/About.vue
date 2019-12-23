<template>
  <div class="about">
    <h1>This is an about page</h1>

    {{username}}
  </div>
</template>

<script>
export default {
  props: {
    username: String,
  },
};
</script>
