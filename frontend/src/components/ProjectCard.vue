<script setup>
defineProps({ project: Object })
import ChapterBlock from './ChapterBlock.vue'
import CharacterBlock from './CharacterBlock.vue'
import LocationBlock from './LocationBlock.vue'
import FactionBlock from './FactionBlock.vue'
import StoryModelBlock from './StoryModelBlock.vue'
</script>

<template>
  <div class="border rounded-xl p-4 mb-6 shadow">
    <h2 class="project-title">{{ project.project.title }}</h2>
    <p class="project-description">{{ project.project.description }}</p>

    <div class="mt-4 space-y-4">
      <StoryModelBlock :model="project.story_model" />
      <ChapterBlock :chapters="project.chapters" :scenes="project.scenes" />
      <CharacterBlock :characters="project.characters" />
      <LocationBlock :locations="project.locations" />
      <FactionBlock :factions="project.factions" />
    </div>
  </div>
</template>

<style scoped>
.project-title {
  font-size: 2rem;
  font-weight: bold;
  background: linear-gradient(black 0%, #3a4eff 50%, black 100%);
  border-bottom-left-radius: 25px;
  border-bottom-right-radius: 25px;
}

.project-description {
  font-size: 1.25rem;
  font-weight: bold;
}
</style>
