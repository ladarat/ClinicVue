<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <router-link :to="{name:'login'}">
      <button>login</button>
    </router-link>

  </div>
</template>

<script>
export default {
  name: 'home',
};
</script>
