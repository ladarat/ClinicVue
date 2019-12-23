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
    <About :username="requestLogin.username"></About>
  </div>
</template>

<script lang="ts">
import About from '@/pages/About.vue';
import { RequestLogin } from '../models/requestLogin';
import loginService from '../services/login.service';

const logoImg = require('../assets/logo.png');

export default {
  name: 'login',
  data() {
    return {
      logoImg,
      requestLogin: {} as RequestLogin,
    };
  },
  components: {
    About,
  },
  methods: {
    login() {
      console.log(this);
      // loginService.login();
      // await loginService.login(self.requestLogin);

      // .then((res) => {
      //   if (res.status === 200) {
      //     alert('success');
      //     self.$router.push({ name: 'home' });
      //   } else {
      //     alert(res.status);
      //   }
      // }).catch((err) => {
      //   console.log(err);
      //   alert('error');
      // });
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
