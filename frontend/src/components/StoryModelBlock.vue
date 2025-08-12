<template>
  <div v-if="model">
    <h3 class="text-lg font-semibold mb-2">Modèle narratif</h3>
    <div class="border rounded p-3">
      <div class="font-bold text-blue-700">
        {{ model.name || 'Modèle non défini' }}
      </div>
      <div class="text-sm text-gray-500 mt-1">
        {{ model.description || 'Aucune description du modèle' }}
      </div>

      <div v-if="model.phases?.length" class="mt-3">
        <h4 class="font-semibold mb-1">Phases :</h4>
        <ul class="list-disc pl-5 space-y-1">
          <li
            v-for="phase in model.phases"
            :key="phase.id"
            class="flex items-center gap-2"
          >
            <span
              class="inline-block w-3 h-3 rounded-full"
              :style="{ backgroundColor: phase.color || '#ccc' }"
            />
            <span>{{ phase.name }}</span>
          </li>
        </ul>
      </div>
    </div>
  </div>

  <div v-else>
    <h3 class="text-lg font-semibold mb-2">Modèle narratif</h3>
    <div class="border rounded p-3 text-sm text-gray-500 italic">
      Aucun modèle défini pour ce projet.
    </div>
  </div>
</template>

<script setup>
defineProps({ model: Object })
</script>
