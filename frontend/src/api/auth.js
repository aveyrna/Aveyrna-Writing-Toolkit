// src/api/auth.js
// ------------------------------------------------------------------
// Wrapper fetch avec credentials inclus + gestion JSON/erreurs.
// Cohabite avec ton backend actuel : cookie HttpOnly pour la session
// ET token "legacy" (Bearer) pour /auth/me tant qu'on ne l'a pas migré.
// ------------------------------------------------------------------

const API_BASE = '/api' // ↔ utilise le proxy Vite en dev

let accessToken = null; // on garde le token en mémoire (pas de localStorage)

/** Permet d'injecter/vider le token "legacy" si besoin */
export function setToken(tok) { accessToken = tok || null; }
export function getToken() { return accessToken; }

async function apiFetch(path, { headers, ...init } = {}) {
    const res = await fetch(`${API_BASE}${path}`, {
        credentials: 'include',                                 // ⬅️ IMPORTANT
        headers: { 'Content-Type': 'application/json', ...(headers || {}) },
        ...init,
    });

    // Tente de parser le JSON même en cas d'erreur HTTP pour renvoyer des infos utiles
    const contentType = res.headers.get('content-type') || '';
    const isJSON = contentType.includes('application/json');
    const data = isJSON ? await res.json().catch(() => ({})) : await res.text().catch(() => '');

    if (!res.ok) {
        const message =
            (data && (data.error || data.message)) ||
            (typeof data === 'string' && data) ||
            `HTTP ${res.status}`;
        const err = new Error(message);
        err.status = res.status;
        err.data = data;
        throw err;
    }

    return data;
}

// -------------------------
// Auth endpoints
// -------------------------

/** POST /auth/register */
export async function register({ username, email, password }) {
    return apiFetch('/auth/register', {
        method: 'POST',
        body: JSON.stringify({ username, email, password }),
    });
}

/** POST /auth/login
 *  - Pose le cookie HttpOnly (via Set-Cookie)
 *  - Récupère aussi le token "legacy" renvoyé par l’API (pour /auth/me actuel)
 */
export async function login({ email, password }) {
    const data = await apiFetch('/auth/login', {
        method: 'POST',
        body: JSON.stringify({ email, password }),
    });

    // Tant que /auth/me exige un Bearer, on stocke le token en mémoire :
    if (data && data.token) setToken(data.token);

    return data; // { id, username, email, token }
}

/** GET /auth/me
 *  - Backend actuel : exige Bearer → on l'ajoute si on l’a.
 *  - Dès que tu migreras /me pour lire le cookie, on pourra retirer l’Authorization.
 */
export async function me() {
    const headers = {};
    const tok = getToken();
    if (tok) headers['Authorization'] = `Bearer ${tok}`;
    return apiFetch('/auth/me', { headers });
}

/** POST /auth/logout
 *  - Supprime côté serveur la session + vide le cookie (backend)
 *  - Vide aussi le token en mémoire côté front
 */
export async function logout() {
    try {
        await apiFetch('/auth/logout', { method: 'POST' });
    } finally {
        setToken(null);
    }
    window.location.reload(); // Recharger pour mettre à jour l'état utilisateur
}
