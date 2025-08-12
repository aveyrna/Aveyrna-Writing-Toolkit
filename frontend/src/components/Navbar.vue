<script setup>
</script>

<template>
    <nav class="dashboard-nav">
        <!-- <a href="/"><img src="/src/assets/img/logo_only_aveyrna_online_story_writer_no_bg.png"
                class="dashboard-logo" /></a>
        <h1> <a href="/">Aveyrna</a></h1> -->
        <ul>
            <li><a href="/dashboard">Overview</a></li>
            <li><a href="/chapters">Chapters</a></li>
            <li><a href="/scenes">Scenes</a></li>
            <li><a href="/characters">Characters</a></li>
            <li><a href="/locations">Locations</a></li>
            <li><a href="/factions">Factions</a></li>
            <li><a href="/mystory/de2556d9-87a1-4df2-92eb-304f5393a1d2">My story</a></li>
        </ul>
    </nav>
</template>

<style scoped>
h1 {
    margin: 0;
    font-size: 2rem;
    background: linear-gradient(aqua, blue);
    -webkit-background-clip: text;
    background-clip: text;
    -webkit-text-fill-color: transparent;
}

ul {
    display: flex;
    margin-left: 200px;
}

li {
    list-style: none;
    padding: 5px;
    margin-left: 15px;
    background: linear-gradient(rgba(58, 78, 255, 0.2), black, rgba(58, 78, 255, 0.2));
    border-radius: 10px;
}

li a {
    color: skyblue;
    text-decoration: none;
    ;
}

li:hover {
    filter: drop-shadow(0 0 1em #646cffaa);
}

li:hover a {
    color: orchid;
}

.dashboard-nav {
    display: flex;
    align-items: center;
    justify-content: flex-start;
    padding: 0px 32px;
    background-color: rgba(12, 16, 51);
    border-bottom: 1px solid orchid;
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    z-index: 100;
}

.dashboard-logo {
    height: 40px;
    margin-right: 16px;
    margin-left: 220px;
}

</style>