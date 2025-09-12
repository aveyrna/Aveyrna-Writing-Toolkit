<script setup>
import { computed } from 'vue'

const props = defineProps({
  chapter: {
    type: Object,
    required: true
  },
  scenes: {
    type: Array,
    required: true,
    default: () => []
  }
})

const filteredScenes = computed(() =>
  props.scenes.filter(scene => scene.chapter_uuid === props.chapter.id)
)
</script>

<template>
  <div class="chapter-card">
    <div class="chapter-header">
      <h3 class="chapter-number">Chapitre {{ chapter.order_index }}:</h3>
      <h2 class="chapter-title">{{ chapter.title }}</h2>
    </div>
    <p class="chapter-summary" v-if="chapter.summary">{{ chapter.summary }}</p>

    <div class="scenes" v-if="filteredScenes.length">
      <h4 class="scene-section-title">Scènes</h4>
      <div class="scene" v-for="scene in filteredScenes" :key="scene.id">
        <div class="scene-title">{{ scene.title }}</div>
        <div class="scene-summary">{{ scene.summary }}</div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.chapter-card {
  background: #f9f9f9;
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

.chapter-header {
  margin-bottom: 8px;
  display: flex;
  flex-direction:row;
  justify-content: space-between;
}

.chapter-number {
  font-weight: bold;
  font-size: 1rem;
  color: #838383;
}

.chapter-title {
  font-size: 1.2rem;
  margin-left: 25px;
}

.chapter-summary {
  font-style: italic;
  color: #666;
  margin-bottom: 12px;
}

.scenes {
  margin-top: 12px;
  margin-left: -200px;
}

.scene-section-title {
  font-weight: bold;
  margin-bottom: 8px;
  font-size: 1rem;
}

.scene {
  padding: 8px;
  border-left: 3px solid #ccc;
  margin-bottom: 10px;
}

.scene-title {
  font-weight: bold;
}

.scene-summary {
  font-size: 0.95rem;
  color: white;
  font-style: italic;
}
</style>
