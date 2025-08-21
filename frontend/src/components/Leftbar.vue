<script setup>
import { onMounted, ref } from 'vue'
import { fetchFullProjectsByUserID } from '../api/projects'
import { me, logout, getToken } from '../api/auth'
import AuthModal from './AuthModal.vue'

const show = ref(false)
const user = ref(null)

async function bootstrapAuth() {
    try {
        user.value = await me()      // tente toujours
    } catch (e) {
        user.value = null            // silencieux si 401
    }
}

function onAuth(data) {
    user.value = {
        username: data.username,
        email: data.email,
        id: data.id
    }
    console.log('Utilisateur connecté/inscrit:', data)
}

async function onLogout() {
    await logout()
    user.value = null
}


//const userID = 1
const projects = ref([])

onMounted(async () => {
    await bootstrapAuth()
    try {
        projects.value = await fetchFullProjectsByUserID(user.value.id)
        projects.value = Array.isArray(projects.value) ? projects.value : []
        console.log("✅ Projets complets chargés :", projects.value)
        console.log("titre du premier projet :", projects.value[0]?.project.title)
    } catch (err) {
        console.error("❌ Erreur chargement projets :", err)
    }
})
</script>

<template>
    <nav class="left-nav">
        <a href="/"><img src="/img/logo_aveyrna.png" class="dashboard-logo" /></a>
        <h1>Aveyrna</h1>
        <p class="desc">Writing Toolkit</p>
        <div>
            <template v-if="user">
                <p>Bonjour <strong>{{ user.username }}</strong></p>
                <button @click="onLogout">Se déconnecter</button>
            </template>
            <template v-else>
                <button @click="show = true">Se connecter / S'inscrire</button>
            </template>

            <AuthModal v-model="show" @success="onAuth" />
        </div>
        <ul>
            <li>
                <details>
                    <summary
                        style="font-size: 1.4em; font-weight: bold; background: linear-gradient(#222222, rgba(58, 78, 255, 0.7), #222222); text-align: center;">
                        My stories</summary>
                    <ul>
                        <!-- List stories -->
                        <p style="font-size: 1.1em; color: skyblue; font-style: italic;">+ Add a story</p>
                        <li v-for="story in projects" :key="story.id">
                            <hr />
                            <details>
                                <summary class="story-summary">{{ story.project.title }}</summary>
                                <ul>
                                    <li>
                                        <details>
                                            <summary class="chapters-summary">CHAPTERS</summary>
                                            <ul>
                                                <!-- List story chapters -->
                                                <p style="font-size: 1.1em; color: orchid; font-style: italic;">+ Add a
                                                    chapter</p>
                                                <hr />
                                                <li v-for="chapter in story.chapters" :key="chapter.id"
                                                    class="chapters-item">
                                                    <details>
                                                        <summary>{{ chapter.title }}</summary>
                                                        <ul>
                                                            <!-- List chapter scenes -->
                                                            <p
                                                                style="font-size: 1.1em; color: lightcoral; font-style: italic;">
                                                                + Add a
                                                                scene</p>
                                                            <hr />
                                                            <li v-for="scene in story.scenes.filter(s => s.chapter_uuid === chapter.id)"
                                                                :key="scene.id" class="scenes-item">
                                                                - {{ scene.title }}
                                                            </li>
                                                        </ul>
                                                    </details>
                                                </li>
                                            </ul>
                                        </details>
                                    </li>
                                    <li>
                                        <details>
                                            <summary class="characters-summary">CHARACTERS</summary>
                                            <ul>
                                                <!-- List story characters -->
                                                <p style="font-size: 1.1em; color: skyblue; font-style: italic;">+ Add a
                                                    character</p>
                                                <hr />
                                                <li v-for="character in story.characters" :key="character.id"
                                                    class="characters-item">
                                                    - {{ character.name }}
                                                </li>
                                            </ul>
                                        </details>
                                    </li>
                                    <li>
                                        <details>
                                            <summary class="locations-summary">LOCATIONS</summary>
                                            <ul>
                                                <!-- List story locations -->
                                                <p style="font-size: 1.1em; color: lightgreen; font-style: italic;">+
                                                    Add a
                                                    location</p>
                                                <hr />
                                                <li v-for="location in story.locations" :key="location.id"
                                                    class="locations-item">
                                                    - {{ location.name }}
                                                </li>
                                            </ul>
                                        </details>
                                    </li>
                                    <li>
                                        <details>
                                            <summary class="factions-summary">FACTIONS</summary>
                                            <ul>
                                                <!-- List story factions -->
                                                <p
                                                    style="font-size: 1.1em; color: rgb(255, 248, 155); font-style: italic;">
                                                    + Add a
                                                    faction</p>
                                                <hr />
                                                <li v-for="faction in story.factions" :key="faction.id"
                                                    class="factions-item">
                                                    - {{ faction.name }}
                                                </li>
                                            </ul>
                                        </details>
                                    </li>
                                </ul>
                            </details>
                        </li>
                    </ul>
                </details>
            </li>
        </ul>
    </nav>
</template>

<style scoped>
h1 {
    margin-top: -10px;
    font-size: 2rem;
    background: linear-gradient(aqua, blue);
    -webkit-background-clip: text;
    background-clip: text;
    -webkit-text-fill-color: transparent;
}

.desc {
    margin-top: -30px;
    font-size: 1.25em;
    font-weight: bold;
    font-style: italic;
}

hr {
    border: 0;
    height: 2px;
    background: linear-gradient(to right, #222222, rgba(58, 78, 255, 0.7), #222222);
}

.left-nav {
    position: fixed;
    top: 0;
    left: 0;
    width: 250px;
    height: 100vh;
    background-color: #222;
    color: #fff;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding-top: 32px;
    box-shadow: 2px 0 8px rgba(0, 0, 0, 0.08);
    z-index: 100;
    text-align: start;
    overflow-y: auto;
    /* Enable vertical scrolling */
}

.dashboard-logo {
    width: 64px;
    height: auto;
    margin-bottom: 16px;
}

.left-nav h2 {
    font-size: 1.5rem;
    margin-bottom: 32px;
    font-weight: 700;
}

.left-nav ul {
    list-style: none;
    padding: 0;
    margin-top: 25px;
    width: 100%;
}

.left-nav li {
    margin-bottom: 20px;
    width: 100%;
    text-align: start;
}

.left-nav a {
    color: #b3e0ff;
    /* lighter blue for links */
    text-decoration: none;
    font-size: 1.1rem;
    transition: color 0.2s;
}

.left-nav a:hover {
    filter: drop-shadow(0 0 1em orchid);
    color: #ffb7fa;
    /* lighter orchid on hover */
}

.story-summary {
    font-size: 1.2em;
    font-weight: bold;
    background: linear-gradient(90deg, #222222 0%, #3a4eff 80%, #222222 100%);
    color: #fff;
    text-align: center;
    margin-top: 25px;
    margin-bottom: 25px;
}

.characters-summary {
    font-weight: bold;
    color: #7fdfff;
    /* brighter blue for characters */
    background: rgba(58, 78, 255, 0.15);
    text-align: center;
}

.characters-item {
    color: skyblue;
    margin-left: 25px;
}

.chapters-summary {
    font-weight: bold;
    color: #e0aaff;
    /* soft purple for chapters */
    background: rgba(218, 112, 214, 0.12);
    text-align: center;
}

.chapters-item {
    color: orchid;
    /* thistle for chapters */
    padding-left: 15px;
}

.scenes-summary {
    font-weight: bold;
    color: #ffb3b3;
    /* soft coral for scenes */
    background: rgba(240, 128, 128, 0.12);
}

.scenes-item {
    color: lightcoral;
    /* light coral for scenes */
    padding-left: 25px;
}

.locations-summary {
    font-weight: bold;
    color: #b3ffb3;
    /* light green for locations */
    background: rgba(144, 238, 144, 0.12);
    text-align: center;
}

.locations-item {
    color: lightgreen;
    /* light green for locations */
    padding-left: 25px;
}

.factions-summary {
    font-weight: bold;
    color: rgb(255, 248, 155);
    /* Dark yellow for factions */
    background: rgba(255, 134, 35, 0.5);
    text-align: center;
}

.factions-item {
    color: rgb(255, 248, 155);
    /* Dark yellow for factions */
    padding-left: 25px;
}
</style>