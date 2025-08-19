// src/api/auth.js
export function setToken(t) { localStorage.setItem('auth_token', t) }
export function getToken() { return localStorage.getItem('auth_token') || '' }
export function clearToken() { localStorage.removeItem('auth_token') }

const BASE = import.meta.env.VITE_API_BASE_URL || ''

async function http(path, opts = {}) {
    const res = await fetch(BASE + path, {
        method: opts.method || 'GET',
        headers: {
            'Content-Type': 'application/json',
            ...(opts.token ? { Authorization: `Bearer ${opts.token}` } : {})
        },
        body: opts.body ? JSON.stringify(opts.body) : undefined
    })
    const data = await res.json().catch(() => ({}))
    if (!res.ok) throw new Error(data?.error || data?.message || `HTTP ${res.status}`)
    return data
}

export async function apiLogout() {
    // côté serveur: 204, pas de JSON requis
    try { await fetch(BASE + '/api/auth/logout', { method: 'POST', credentials: 'include' }) }
    catch { /* ignore */ }
    clearToken()
}

export async function apiRegister({ username, email, password }) {
    const data = await http('/api/auth/register', { method: 'POST', body: { username, email, password } })
    if (data?.token) setToken(data.token)
    return data
}
export async function apiLogin({ email, password }) {
    const data = await http('/api/auth/login', { method: 'POST', body: { email, password } })
    if (data?.token) setToken(data.token)
    return data
}
export async function apiMe() {
    const t = getToken(); if (!t) return null
    try { return await http('/api/auth/me', { token: t }) }
    catch { clearToken(); return null }
}
