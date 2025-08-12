<template>
  <div v-if="scenes?.length">
    <h4 class="font-semibold">Scènes ({{ scenes.length }})</h4>
    <ul class="list-disc pl-5">
      <li v-for="scene in scenes" :key="scene.id">
        {{ scene.title }} – {{ scene.phase }} ({{ scene.characters?.length || 0 }} perso)
      </li>
    </ul>
  </div>
  <div v-else>
    <p class="text-sm italic text-gray-500">Aucune scène dans ce chapitre.</p>
  </div>
</template>

<script setup>
defineProps({ scenes: Array })
</script>
