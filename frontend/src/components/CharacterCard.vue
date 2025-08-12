<script setup>
const props = defineProps({
  character: {
    type: Object,
    required: true
  }
})
</script>

<template>
  <div class="character-card">
    <img
      :src="`/img/characters/${character.avatar_url}`"
      alt="Avatar"
      class="avatar"
    />
    <div class="info">
      <div>
        <h3 class="name">{{ character.name }}</h3>
        <p class="role" v-if="character.role">{{ character.role }}</p>
      </div>
      <div class="details">
        <p class="objectif"><strong>Objectif :</strong> {{ character.objective || '—' }}</p>
        <p class="conflit"><strong>Conflit interne :</strong><br />{{ character.internal_conflict || '—' }}</p>
        <p class="arc"><strong>Arc narratif :</strong> {{ character.arc || '—' }}</p>
      </div>
    </div>
  </div>
</template>



<style scoped>
.character-card {
  background: #1e1e1e;
  color: white;
  display: flex;
  gap: 1rem;
  padding: 1rem;
  border-radius: 12px;
  width: 380px;
  flex-shrink: 0;
  text-align: left;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.4);
  transition: transform 0.2s;
}

.character-card:hover {
  transform: translateY(-4px);
}

.avatar {
  width: 100px;
  height: 100px;
  object-fit: cover;
  border-radius: 8px;
  flex-shrink: 0;
}

.info {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  flex-grow: 1;
}

.name {
  font-size: 1.1rem;
  font-weight: bold;
  margin: 0 0 0.25rem;
}

.role {
  font-size: 0.9rem;
  color: #ccc;
  font-style: italic;
  margin-bottom: 0.5rem;
  margin-top: -5px;
}

.details {
  font-size: 0.85rem;
}

.objectif {
  margin-top: 0;
}

.conflit {
  margin-left: -125px;
}

.arc {
  margin-left: -125px;
}
</style>
