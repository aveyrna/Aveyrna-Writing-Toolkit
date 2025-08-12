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
