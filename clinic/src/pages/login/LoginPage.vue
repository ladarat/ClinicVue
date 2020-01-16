<template>
  <div class="login-page">
    <div class="container">
      <div>{{loginTitle}}</div>
      <hr />
      <div>{{usernameTitle}}</div>
      <input type="text" class="form-control" v-model="username" />
      <div>{{passwordTitle}}</div>
      <input type="password" class="form-control" v-model="password" />
      <button class="btn btn-primary mt-3" @click="login()">
        {{signInBtnText}}
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import UserService from "../../common/api/modules/user/user.service";
import LoginRequest from "../../common/api/modules/user/models/LoginRequest";

export default class LoginPage extends Vue {
  loginTitle: string = "เข้าสู่ระบบ";
  usernameTitle: string = "User Name:";
  passwordTitle: string = "Password:";
  signInBtnText: string = "Sign In";
  username: string = '';
  password: string = '';
  login() {
    let request: LoginRequest = {
      username: this.username,
      password: this.password
    }
    UserService.login(request).then(() => {

    })
  }
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