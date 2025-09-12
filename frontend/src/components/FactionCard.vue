<script setup>

const props = defineProps({
  faction: {
    type: Object,
    required: true
  }
})
</script>

<template>
  <div class="faction-card">
    <div class="faction-header">
      <h3 class="faction-name">{{ faction.name }}</h3>
      <h4 class="faction-type">{{ faction.type }}</h4>
    </div>
    <p class="faction-description" v-if="faction.description">{{ faction.description }}</p>
  </div>
</template>

<style scoped>
.faction-card {
  background: rgba(144, 238, 144, 0.5);
  color: white;
  display: flex;
  gap: 1rem;
  padding: 1rem;
  border-radius: 12px;
  width: 380px;
  flex-shrink: 0;
  text-align: left;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.4);
  transition: transform 0.2s;
}

.faction-header {
  margin-bottom: 8px;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

.faction-name {
  font-weight: bold;
  font-size: 1rem;
  color: white;
}

.faction-type {
  font-size: 1.2rem;
  margin-left: 25px;
}

.faction-description {
  font-style: italic;
  color: white;
  margin-bottom: 12px;
}
</style>
