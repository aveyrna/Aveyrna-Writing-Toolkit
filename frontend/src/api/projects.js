export async function deleteProjectByUUID(uuid) {
  const res = await fetch(`/api/projects/public/${uuid}`, { method: 'DELETE' })
  const text = await res.text()
  console.log("↪ Réponse brute (deleteProjectByUUID):", text)

  if (!res.ok) {
    throw new Error(`Erreur ${res.status} : ${res.statusText} → ${text}`)
  }

  // 204 attendu → pas de body
  return {}
}


export async function fetchFullProjectByUUID(uuid) {
  try {
    const res = await fetch(`/api/projects/public/${uuid}/full`)
    const text = await res.text()
    console.log("↪ Réponse brute :", text)

    if (!res.ok) {
      throw new Error(`Erreur ${res.status} : ${res.statusText} → ${text}`)
    }

    return JSON.parse(text)
  } catch (err) {
    console.error("⚠️ Erreur dans fetchFullProjectByUUID:", err)
    throw err
  }
}

export async function fetchFullProjectsByUserID(userID) {
  const res = await fetch(`/api/projects/user/${userID}/full`)
  if (!res.ok) {
    const text = await res.text()
    throw new Error(`Erreur ${res.status} → ${text}`)
  }
  return res.json()
}

export async function createProject({ title, description = '', story_model_id = null }) {
  try {
    const res = await fetch(`/api/projects`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      // Pas de user_id ici : il est déduit du cookie côté serveur
      body: JSON.stringify({ title, description, story_model_id }),
    })

    const text = await res.text()
    console.log("↪ Réponse brute (createProject):", text)

    if (!res.ok) {
      throw new Error(`Erreur ${res.status} : ${res.statusText} → ${text}`)
    }

    return JSON.parse(text) // { id, public_id, user_id, title, description, story_model_id, created_at }
  } catch (err) {
    console.error("⚠️ Erreur dans createProject:", err)
    throw err
  }
}

