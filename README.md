# 📡 RSS API – Web-Based Feed Aggregator

**RSS API** is a full-stack application that allows users to subscribe to RSS feeds from various sources and monitor these feeds through periodic automatic updates. A RESTful API connects a React-based frontend with a backend service written in Go.

---

## 🔧 Technical Specifications

- **Frontend:** React (located in the `rssapi-frontend` directory)
- **Backend:** Go (RESTful API using the `chi` router)
- **Database:** PostgreSQL
- **Background Tasks:** Automatic feed polling using Go routines and `time.Ticker`
- **RSS Parsing:** Real-time content fetching via the `gofeed` library
- **Authentication:** API Key-based authentication (via `Authorization` header)
- **Features:** User creation, feed addition, subscriptions, and personalized post listings
- **Message Queue:** Not used
- **Container Management:** Entire system can be deployed via `docker-compose.yml`

---

## 🔄 Automatic Data Refresh

📌 A background worker runs **every 30 seconds** to scan all user-added RSS feed URLs. If new content (posts) is found, it is stored in the database.

```go

handlers.StartFeedWorker(database, 30*time.Second)
```


## 🌟 Highlighted Features
⏱️ Feed Updates: Automatically scans RSS feeds every 30 seconds

🔐 Secure Access with API Key: Each user receives a unique key, used in API requests

📬 Subscription System: Users can subscribe to any feed they like

📰 Personalized Content: Users only see posts from feeds they are subscribed to

🐳 Easy Setup: Spin up frontend, backend, and database with a single Docker Compose command


## 🚀 Setup and Run

### 1. Clone the repository
git clone https://github.com/bayrambartu/rssapi.git
cd rssapi

### 2. Launch the Docker containers
docker-compose up --build

### 3. Access the services
Frontend: http://localhost:3001  
Backend API: http://localhost:3000
