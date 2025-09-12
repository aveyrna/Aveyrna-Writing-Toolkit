<script setup>

const props = defineProps({
  location: {
    type: Object,
    required: true
  }
})
</script>

<template>
  <div class="location-card">
    <div class="location-header">
      <h3 class="location-name">{{ location.name }}</h3>
      <h4 class="location-type">{{ location.type }}</h4>
    </div>
    <p class="location-description" v-if="location.description">{{ location.description }}</p>
  </div>
</template>

<style scoped>
.location-card {
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

.location-header {
  margin-bottom: 8px;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

.location-name {
  font-weight: bold;
  font-size: 1rem;
  color: white;
}

.location-type {
  font-size: 1.2rem;
  margin-left: 25px;
}

.location-description {
  font-style: italic;
  color: white;
  margin-bottom: 12px;
}
</style>
