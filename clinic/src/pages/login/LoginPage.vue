<template>
  <div class="login-page">
    <div class="container">
      <div id="login-title">{{loginTitle}}</div>
      <hr />
      <div>{{usernameTitle}}</div>
      <input id="username-input" type="text" class="form-control" v-model="username" />
      <div>{{passwordTitle}}</div>
      <input id="password-input" type="password" class="form-control" v-model="password" />
      <button id="login-button" class="btn btn-primary mt-3" @click="login()">{{signInBtnText}}</button>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import UserService from "../../common/api/modules/user/user.service";
import LoginRequest from "../../common/api/modules/user/models/LoginRequest";
import { LOGIN } from "@/store/modules/auth/action.type";

@Component
export default class LoginPage extends Vue {
  loginTitle: string = "เข้าสู่ระบบ";
  usernameTitle: string = "User Name:";
  passwordTitle: string = "Password:";
  signInBtnText: string = "Sign In";
  username: string = "";
  password: string = "";
  login() {
    let request: LoginRequest = {
      username: this.username,
      password: this.password
    };
    this.$store
      .dispatch(LOGIN, request)
      .then(() => this.$router.push({ name: "menu" })
      )
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