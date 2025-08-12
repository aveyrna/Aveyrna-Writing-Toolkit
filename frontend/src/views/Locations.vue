<script setup>
  import Navbar from '../components/Navbar.vue';
  import Leftbar from '../components/Leftbar.vue';
</script>

<template>
  <div>
    <Navbar />
    <Leftbar />
    <div class="container">
      <h1 class="page-title">Locations</h1>
      <p>This will be the Locations page.</p>
    </div>
  </div>
</template>

<style>

body {
  background-color: black;
}

h1 {
  color: orchid;
}

p {
  font-size: 1.2em;
  color: skyblue;
}

.logo {
  height: 150px;
}

.desc {
  font-size: 1.5em;
  color: skyblue;
  margin-bottom: 20px;
}
</style>