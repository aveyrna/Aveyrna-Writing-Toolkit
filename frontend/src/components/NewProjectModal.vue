<script setup>
import { ref, watch } from 'vue'
import { createProject } from '../api/projects'

const props = defineProps({
  modelValue: { type: Boolean, default: false }
})
const emit = defineEmits(['update:modelValue', 'created'])
const dialog = ref(null)

const form = ref({ title: '', description: '' })
const error = ref('')

watch(
  () => props.modelValue,
  (v) => {
    if (v) dialog.value?.showModal()
    else dialog.value?.close()
  }
)

function close() {
  emit('update:modelValue', false)
}

async function onSubmit() {
  error.value = ''
  try {
    const project = await createProject({
      title: form.value.title,
      description: form.value.description
    })
    emit('created', project)
    close()
    form.value = { title: '', description: '' } // reset
  } catch (e) {
    error.value = e.message || 'Erreur serveur'
  }
}
</script>

<template>
  <dialog ref="dialog">
    <h2>Nouveau projet</h2>
    <form @submit.prevent="onSubmit">
      <label>Titre :
        <input v-model="form.title" type="text" required />
      </label>
      <label>Description :
        <textarea v-model="form.description"></textarea>
      </label>
      <button type="submit">Créer</button>
    </form>
    <p v-if="error" style="color:red">{{ error }}</p>
    <button @click="close">Fermer</button>
  </dialog>
</template>
