<script setup>
defineProps({ model: Object })
</script>

<template>
  <div v-if="model">
    <h3 class="model-title">Modèle narratif</h3>
    <div>
      <div v-if="model.story_model_id == 1" class="model-text">
        Save the Cat!
      </div>
      <div v-if="model.phases?.length">
        <h4>Phases :</h4>
        <ul>
          <li v-for="phase in model.phases" :key="phase.id">
            <span :style="{ backgroundColor: phase.color || '#ccc' }" />
            <span>{{ phase.name }}</span>
          </li>
        </ul>
      </div>
    </div>
  </div>

  <div v-else>
    <h3>Modèle narratif</h3>
    <div>
      Aucun modèle défini pour ce projet.
    </div>
  </div>
</template>

<style scoped>
.model-text {
  margin-top: -20px;
}
</style>
