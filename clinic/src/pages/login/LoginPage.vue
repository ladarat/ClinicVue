<template>
  <div class="login-page">
    <div class="container">
      <div>{{loginTitle}}</div>
      <hr>
      <div>{{usernameTitle}}</div>
      <input type="text" class="form-control">
      <div>{{passwordTitle}}</div>
      <input type="password" class="form-control">
      <button class="btn btn-primary mt-3">
        {{signInBtnText}}
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default class LoginPage extends Vue {
  loginTitle: string = 'เข้าสู่ระบบ'
  usernameTitle: string = 'User Name:'
  passwordTitle: string = 'Password:'
  signInBtnText: string = 'Sign In'
}
</script>

<style lang="scss" scoped>
.login-page {
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>