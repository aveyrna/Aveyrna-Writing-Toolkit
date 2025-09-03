<script setup>
defineProps({ characters: { type: Array, default: () => [] } })
import CharacterCard from './CharacterCard.vue'
</script>

<template>
  <div class="character-block">
    <h3 class="characters-title">Personnages ({{ characters.length }})</h3>
    <div class="scroll-row">
      <CharacterCard v-for="character in characters" :key="character.id" :character="character" />
    </div>
  </div>
</template>

<style scoped>
.characters-title {
  font-size: 1.3rem;
  font-weight: bold;
  background: linear-gradient(black 0%, rgb(29, 16, 150) 50%, black 100%);
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
  scrollbar-color: #7fdfff #222;
}

.scroll-row::-webkit-scrollbar {
  height: 8px;
  background: #222;
  border-radius: 4px;
}

.scroll-row::-webkit-scrollbar-thumb {
  background: #7fdfff;
  border-radius: 4px;
}

.scroll-row::-webkit-scrollbar-thumb:hover {
  background: #ffb7fa;
}

.character-card {
  flex-shrink: 0;
  padding: 1rem;
  border: 1px solid #eee;
  border-radius: 0.5rem;
  background: rgba(58, 78, 255, 0.25);
  transition: background 0.3s;
}

.character-card:hover {
  filter: drop-shadow(0 0 1em #646cffaa);
}
</style>