# ✏️ Aveyrna Writing Toolkit

Outil web pour structurer, organiser et gérer des projets narratifs.  
Conçu pour les auteurs et game designers, avec une approche orientée **story models** et gestion fine des personnages, lieux, chapitres et scènes.

---

## 📄 Présentation

Aveyrna est une application web full-stack permettant :
- La création et gestion de **projets narratifs**
- L'organisation en **chapitres, scènes, personnages, lieux**
- L'utilisation de **modèles narratifs** (Save the Cat, Hero's Journey, etc.)
- Une **interface moderne** avec Vue 3
- Un **backend performant** en Go + PostgreSQL

---

## 🏗 Stack technique

### Frontend
- [Vue 3](https://vuejs.org/)
- [Vite](https://vitejs.dev/)
- [Pinia](https://pinia.vuejs.org/) pour le store
- [Axios](https://axios-http.com/) pour les appels API

### Backend
- [Go](https://go.dev/) 1.22+
- [chi](https://github.com/go-chi/chi) pour le routing
- [pgx](https://github.com/jackc/pgx) pour la connexion PostgreSQL
- Architecture REST avec routes modulaires

### Base de données
- [PostgreSQL 16](https://www.postgresql.org/)
- Hébergée sur [OVH Web Cloud Database](https://www.ovhcloud.com/fr/web-cloud-databases/)

---

## ⚙️ Installation locale (dev)

### 1. Cloner le projet
git clone https://github.com/toncompte/aveyrna.git
cd aveyrna

### 2. Backend Go
#### Installer les dépendances
cd backend
go mod tidy

#### Configurer l'environnement
Créer un fichier `.env` dans `backend` :

DATABASE_URL=postgres://`<user>:<password>@<host>:<port>/<dbname>?sslmode=require`

(voir Gdrive)

#### Lancer le backend
go run .
API dispo sur http://localhost:8080.

### 3. Frontend Vue 3
#### Installer les dépendances
cd frontend
npm install

#### Lancer le frontend
npm run dev
Interface dispo sur http://localhost:5173.

---

## 🚀 Connexion à la DB OVH

1. Créer une **instance PostgreSQL Web Cloud Database** sur OVH.
2. Créer la base (`aveyrna_db`) et l’utilisateur (`aveyrna`).
3. Whitelister l’IP publique de la machine qui se connecte (ou celle du VPN avec IP dédiée).
4. Restaurer le dump local `.backup` :
pg_restore \
  --host=<ovh_host> \
  --port=<ovh_port> \
  --username=<user> \
  --dbname=<dbname> \
  --no-owner \
  --no-privileges \
  --clean --if-exists \
  --disable-triggers \
  --verbose \
  /chemin/vers/aveyrna.backup

5. Vérifier la connexion :
psql "host=<ovh_host> port=<ovh_port> dbname=<dbname> user=<user> password=<password> sslmode=require" -c "SELECT version();"

---

## 🔄 Déploiement

### Option VPS OVH
- Héberger **frontend** et **backend** sur le même VPS avec Nginx en reverse proxy
- Connexion sécurisée à la DB OVH via IP whitelistée ou VPN IP dédiée

### Option hébergement séparé
- Frontend sur hébergement mutualisé OVH
- Backend sur VPS ou autre serveur
- DB OVH accessible publiquement avec IP filtrée

---

## 🛠 Commandes utiles

### Lancer le backend avec DATABASE_URL en env
$Env:DATABASE_URL="postgres://user:pass@host:port/dbname?sslmode=require"; go run .

### Lancer le frontend
npm run dev

---

## 📜 Licence
Projet privé – © Aubry Varen
