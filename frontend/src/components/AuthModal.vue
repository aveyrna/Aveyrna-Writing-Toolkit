<script setup>
import { ref, watch } from 'vue'
import { apiLogin, apiRegister } from '../api/auth'

const props = defineProps({ modelValue: { type: Boolean, default: false } })
const emit = defineEmits(['update:modelValue', 'success'])
const dialog = ref(null)

const mode = ref('login')
const form = ref({ username: '', email: '', password: '' })
const error = ref('')
const success = ref('')

watch(() => props.modelValue, v => v ? dialog.value?.showModal() : dialog.value?.close())
function close() { emit('update:modelValue', false) }
function toggleMode() { mode.value = mode.value === 'login' ? 'register' : 'login'; error.value = ''; success.value = '' }

async function onSubmit() {
    error.value = ''; success.value = ''
    try {
        const data = mode.value === 'login'
            ? await apiLogin({ email: form.value.email, password: form.value.password })
            : await apiRegister({ username: form.value.username, email: form.value.email, password: form.value.password })
        emit('success', data)
        close()
    } catch (e) { error.value = e.message || 'Erreur' }
}
</script>


<template>
    <dialog ref="dialog">
        <h2>{{ mode === 'login' ? 'Connexion' : 'Inscription' }}</h2>
        <form @submit.prevent="onSubmit">
            <div v-if="mode === 'register'">
                <label>Username :
                    <input v-model="form.username" type="text" required />
                </label>
            </div>
            <label>Email :
                <input v-model="form.email" type="email" required />
            </label>
            <label>Mot de passe :
                <input v-model="form.password" type="password" required />
            </label>
            <button type="submit">{{ mode === 'login' ? 'Se connecter' : "S'inscrire" }}</button>
        </form>
        <p v-if="error" style="color:red">{{ error }}</p>
        <button @click="toggleMode">
            {{ mode === 'login' ? "Pas de compte ? S'inscrire" : "Déjà inscrit ? Se connecter" }}
        </button>
        <button @click="close">Fermer</button>
    </dialog>
</template>
