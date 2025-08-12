<script setup>
import Navbar from '../components/Navbar.vue';
import Leftbar from '../components/Leftbar.vue';
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { fetchFullProjectByUUID } from '../api/projects'

const route = useRoute()
const uuid = route.params.uuid

console.log("UUID du projet :", uuid)

const project = ref(null)
const loading = ref(true)
const error = ref(null)

onMounted(async () => {
  try {
    project.value = await fetchFullProjectByUUID(uuid)
  } catch (err) {
    console.error("Erreur chargement projet :", err)
    error.value = err.message
  } finally {
    loading.value = false
  }
})

</script>

<template>
  <div>
    <Navbar />
    <Leftbar />
    <div class="container">
      <h1 class="page-title">Story</h1>
      <div v-if="loading">Chargement du projet...</div>
      <div v-else-if="error">Erreur : {{ error }}</div>
      <div v-else>
        <h1>{{ project.project.title }}</h1>
        <p>{{ project.project.description }}</p>

        <h2>Personnages</h2>
        <ul>
          <li v-for="c in project.characters" :key="c.id">{{ c.name }} — {{ c.role }}</li>
        </ul>

        <!-- Tu peux afficher les autres sections ici -->
        <h2>Locations</h2>
        <ul>
          <li v-for="l in project.locations" :key="l.id">{{ l.name }} — {{ l.description }}</li>
        </ul>

        <h2>Chapters</h2>
        <ul>
          <li v-for="ch in project.chapters" :key="ch.id">{{ ch.title }} — {{ ch.summary }}</li>
        </ul>
      </div>
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