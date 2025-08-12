<script setup>
import Leftbar from '../components/Leftbar.vue';

import { ref, onMounted } from 'vue'
import ProjectCard from '../components/ProjectCard.vue'
import { fetchFullProjectsByUserID } from '../api/projects'


const userID = 1
const projects = ref([])

onMounted(async () => {
    try {
        projects.value = await fetchFullProjectsByUserID(userID)
        console.log("✅ Projets complets chargés dans dashboard :", projects.value)
    } catch (err) {
        console.error("❌ Erreur chargement projets :", err)
    }
})
</script>

<template>
  <div>
    <Leftbar />
    <div class="container">
      <div class="p-6">
        <h1>Tableau de bord narratif</h1>
        <div v-if="projects.length">
          <ProjectCard v-for="project in projects" :key="project.id" :project="project" />
        </div>
        <div v-else>
          <p>Aucun projet trouvé.</p>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
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