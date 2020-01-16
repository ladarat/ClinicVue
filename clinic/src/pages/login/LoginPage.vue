<template>
  <div>
    <div>{{loginTitle}}</div>
    <div>{{usernameTitle}}</div>
    <div>{{passwordTitle}}</div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default class LoginPage extends Vue {
  loginTitle: string = 'เข้าสู่ระบบ'
  usernameTitle: string = 'User Name:'
  passwordTitle: string = 'Password:'
}
</script>