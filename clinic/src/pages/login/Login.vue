<template>
  <div class="login center">
    <h5>Login</h5>
    <div class="form-group">
      <label>Username</label>
      <input class="form-control" placeholder="Username" />
    </div>
    <div class="form-group">
      <label>Password</label>
      <input type="password" class="form-control" placeholder="Password" />
    </div>
    <button class="btn btn-primary" @click="onLogin();">Submit</button>
  </div>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      username: "",
      password: ""
    };
  },
  methods: {
    onLogin() {
      this.$swal.fire("Success", "", "success");
      this.$router.push({ name: 'menu' });
    }
  }
};
</script>

<style lang="css" scoped>
.login {
  width: 100%;
  max-width: 40%;
  padding: 5%;
  margin: auto;
}
</style>