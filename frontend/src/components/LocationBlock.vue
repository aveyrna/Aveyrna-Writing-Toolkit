<script setup>
import LocationCard from './LocationCard.vue'

defineProps({ locations: Array })
</script>

<template>
  <div class="location-block">
    <h3 class="locations-title">Lieux ({{ locations.length }})</h3>
    <div class="scroll-row">
      <LocationCard v-for="location in locations" :key="location.id" :location="location" />
    </div>
  </div>
</template>

<style scoped>
.locations-title {
  font-size: 1.3rem;
  font-weight: bold;
  background: linear-gradient(black 0%, rgb(1, 112, 51) 50%, black 100%);
  border-bottom-left-radius: 25px;
  border-bottom-right-radius: 25px;
}

.scroll-row {
  display: flex;
  overflow-x: auto;
  gap: 1rem;
  max-width: 100%;
  padding-bottom: 8px;
  scrollbar-width: thin;
  scrollbar-color: #71fc9f #222;
}

.scroll-row::-webkit-scrollbar {
  height: 8px;
  background: #222;
  border-radius: 4px;
}

.scroll-row::-webkit-scrollbar-thumb {
  background: #71fc9f;
  border-radius: 4px;
}

.scroll-row::-webkit-scrollbar-thumb:hover {
  background: #215f36;
}

.location-card {
  flex-shrink: 0;
  width: 300px;
  padding: 1rem;
  border: 1px solid #eee;
  border-radius: 0.5rem;
  background: rgba(144, 238, 144, 0.5);
  transition: background 0.3s;
}

.location-card:hover {
  filter: drop-shadow(0 0 1em #646cffaa);
}
</style>