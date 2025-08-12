import { createRouter, createWebHistory } from 'vue-router';

const routes = [
    {
        path: '/',
        name: 'Home',
        component: () => import('./views/Home.vue') // Lazy-loaded
    },

    {
        path: '/about',
        name: 'About',
        component: () => import('./views/About.vue') // Lazy-loaded
    },

    {
        path: '/login',
        name: 'Login',
        component: () => import('./views/Login.vue') // Lazy-loaded
    },

    {
        path: '/register',
        name: 'Register',
        component: () => import('./views/Register.vue') // Lazy-loaded
    },

    {
        path: '/profile',
        name: 'Profile',
        component: () => import('./views/Profile.vue') // Lazy-loaded
    },

    {
        path: '/dashboard',
        name: 'Dashboard',
        component: () => import('./views/Dashboard.vue') // Lazy-loaded
    },

    {
        path: '/scenes',
        name: 'Scenes',
        component: () => import('./views/Scenes.vue') // Lazy-loaded
    },

    {
        path: '/characters',
        name: 'Characters',
        component: () => import('./views/Characters.vue') // Lazy-loaded
    },

    {
        path: '/locations',
        name: 'Locations',
        component: () => import('./views/Locations.vue') // Lazy-loaded
    },

    {
        path: '/mystory/:uuid',
        name: 'MyStory',
        component: () => import('./views/Story.vue') // Lazy-loaded
    }
];

const router = createRouter({
    history: createWebHistory(),
    routes
});

export default router;
