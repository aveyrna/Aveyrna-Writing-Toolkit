<script setup> </script>
<template>
  <div>
    <img src="/img/logo_aveyrna_online_story_writer_no_bg.png" class="logo" />
    <p class="desc">Online Writing Toolkit</p>

    <p class="about">Your all-in-one workspace to structure, plan, and organize your stories.</p>

    <router-link to="/dashboard">
      <button class="startButton">Start</button>
    </router-link>
  </div>
</template>


<style scoped>

body {
  background-color: black;
}

.logo {
  height: 25em;
  will-change: filter;
  transition: filter 300ms;
}


.desc {
  color: skyblue;
  margin-top: -10px;
  font-size: 1.5em;
}

.about {
  font-size: 1.6em;
  margin-top: 100px;
  color: skyblue;
}

.startButton {
  margin-top: 50px;
  font-size: 1.7em;
  font-weight: bold;
  color: orchid;
  padding: 10px 20px 10px 20px;
}

.startButton:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}

</style>