<script setup>
  import Navbar from '../components/Navbar.vue';
</script>

<template>
  <div>
     <Navbar />
     <a href="/"><img src="/src/assets/img/logo_only_aveyrna_online_story_writer_no_bg.png" class="logo" /></a>
    <h1>About</h1>
    <p>This is the about page.</p>
  </div>
</template>

<style>
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