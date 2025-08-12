<script setup>

</script>

<template>
  <div class="login-form">
  <a href="/"><img src="/src/assets/img/logo_only_aveyrna_online_story_writer_no_bg.png" class="form-logo" /></a>
    <h1>Login</h1>
    <form class="form-group">
      <input v-model="emailLogin" type="email" class="form-control" placeholder="Email" required>
      <input v-model="passwordLogin" type="password" class="form-control" placeholder="Password" required>
      <button type="submit" class="btn btn-primary" @click="doLogin"> Login</button>
      <p v-if="emptyFields" style="color: red;">Please fill in all fields.</p>
      <p v-if="loginError" style="color: red;">{{ loginError }}</p>
      <p style="color: skyblue" >Don't have an account? <a href="/register" @click="registerActive = !registerActive, emptyFields = false">Sign up
          here</a>
      </p>
      <p><a href="#">Forgot your password?</a></p>
    </form>
  </div>
</template>

<style>

body {
  background-color: black;
}

.form-logo {
  height: 150px;
}

.login-form {
  background-color:rgba(58, 78, 255, 0.2);
  padding: 50px;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  max-width: 400px;
  margin: auto;
}

h1 {
  color: skyblue;
  padding-top: 0;
  margin-top: 0;
}

p {
  font-size: 1.2em;
  color: #35495e;
}

.form-group {
  max-width: 400px;
  margin: auto;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.form-control {
  margin-bottom: 10px;
  padding: 10px;
  background-color: rgba(255, 255, 255, 0.8);
  border: 1px solid #ccc;
  border-radius: 5px;
  font-size: 1em;
  color: #333;
}

.btn {
 margin-top: 50px;
  font-size: 1.7em;
  font-weight: bold;
  color: orchid;
  padding: 10px 20px 10px 20px;
}

.btn:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}

</style>