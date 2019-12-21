<template>
  <div>
    <img ref="logo" :src="logoImg" />
    <div class="login-container">
      <div>เข้าสู่ระบบ</div>
      <div>
        username
        <input type="text" v-model="requestLogin.username" />
      </div>
      <div>
        password
        <input type="password" v-model="requestLogin.password" />
      </div>
      <button @click="login();">LOGIN</button>
      {{requestLogin.username}} + {{requestLogin.password}}
    </div>
  </div>
</template>

<script lang="ts">
import { RequestLogin } from '../models/requestLogin';

const logoImg = require('../assets/logo.png');

export default {
  name: 'login',
  data() {
    return {
      logoImg,
      requestLogin: {} as RequestLogin,
    };
  },
  methods: {
    login() {
      alert(this.requestLogin.username);
    },
  },
};
</script>

<style lang="scss" scoped>
.login-container {
  border: 1px solid lightgreen;
  border-radius: 10px;
}
</style>
