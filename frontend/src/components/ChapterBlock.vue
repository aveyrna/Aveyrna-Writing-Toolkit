<script setup>
import ChapterCard from './ChapterCard.vue'

defineProps({
  chapters: { type: Array, default: () => [] },
  scenes: { type: Array, default: () => [] }
})
</script>

<template>
  <div class="chapter-block">
    <h2 class="block-title">Chapitres ({{ chapters.length }})</h2>
    <div class="scroll-row">
      <ChapterCard v-for="chapter in chapters" :key="chapter.id" :chapter="chapter" :scenes="scenes" />
    </div>
  </div>
</template>


<style scoped>
.block-title {
  font-size: 1.3rem;
  font-weight: bold;
  background: linear-gradient(black 0%, rgb(150, 16, 145) 50%, black 100%);
  border-bottom-left-radius: 25px;
  border-bottom-right-radius: 25px;
}

.scroll-row {
  display: flex;
  overflow-x: auto;
  gap: 1rem;
  padding-bottom: 8px;
  /* espace pour ne pas coller au texte */
  scrollbar-width: thin;
  scrollbar-color: #da70d6 #222;
  max-width: 100%;
}

.scroll-row::-webkit-scrollbar {
  height: 8px;
  background: #222;
  border-radius: 4px;
}

.scroll-row::-webkit-scrollbar-thumb {
  background: #da70d6;
  border-radius: 4px;
}

.scroll-row::-webkit-scrollbar-thumb:hover {
  background: #b266b2;
}

.chapter-card {
  flex-shrink: 0;
  min-width: 300px;
  max-width: 400px;
  padding: 1rem;
  border: 1px solid #eee;
  border-radius: 0.5rem;
  background: rgba(218, 112, 214, 0.25);
  transition: background 0.3s;
}

.chapter-card:hover {
  filter: drop-shadow(0 0 1em orchid);
}
</style>
