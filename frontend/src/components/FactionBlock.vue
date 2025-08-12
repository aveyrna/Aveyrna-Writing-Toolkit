<script setup>
defineProps({ factions: Array })
</script>

<template>
  <div class="faction-block">
    <h3 class="factions-title">Factions ({{ factions.length }})</h3>
    <div class="scroll-row">
      <div
        v-for="faction in factions"
        :key="faction.id"
        class="faction-card"
      >
        <div class="font-bold">{{ faction.name }}</div>
        <div class="text-sm text-gray-500 mb-2">
          {{ faction.description || 'Aucune description spécifiée' }}
        </div>
        <div v-if="faction.members?.length" class="text-sm">
          Membres :
          <ul class="list-disc ml-5">
            <li v-for="member in faction.members" :key="member.id">{{ member.name }}</li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

.factions-title {
  font-size: 1.3rem;
  font-weight: bold;
  background: linear-gradient(black 0%, rgb(124, 122, 13) 50%, black 100%);
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
  scrollbar-color: #e5ff86 #222;
}

.scroll-row::-webkit-scrollbar {
  height: 8px;
  background: #222;
  border-radius: 4px;
}

.scroll-row::-webkit-scrollbar-thumb {
  background: #e5ff86;
  border-radius: 4px;
}

.scroll-row::-webkit-scrollbar-thumb:hover {
  background: #b8cc3f;
}

.faction-card {
  flex-shrink: 0;
  width: 300px;
  padding: 1rem;
  border: 1px solid #eee;
  border-radius: 0.5rem;
  background: rgba(255, 134, 35, 0.8);
  transition: background 0.3s;
}

.faction-card:hover {
  filter: drop-shadow(0 0 1em #646cffaa);
}
</style>
