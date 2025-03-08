# LeetCursor

A web application for tracking and comparing LeetCode progress across users.

## Project Structure

competitive-coding-leaderboard/
├── backend/
│ ├── cmd/
│ │ └── server/
│ │ └── main.go
│ ├── internal/
│ │ ├── auth/
│ │ ├── config/
│ │ ├── handlers/
│ │ ├── middleware/
│ │ ├── models/
│ │ ├── repository/
│ │ ├── services/
│ │ └── worker/
│ ├── Dockerfile
│ └── go.mod
├── frontend/
│ ├── src/
│ │ ├── components/
│ │ ├── pages/
│ │ ├── services/
│ │ ├── store/
│ │ └── types/
│ ├── Dockerfile
│ └── package.json
├── docker-compose.yml
└── Makefile

## Getting Started

### Prerequisites

- Node.js (v16 or higher)
- Go (v1.19 or higher)
- Docker (optional)

### Frontend Setup

1. Navigate to the frontend directory:

2. Install dependencies:

```bash
npm install
```

3. Create a `.env` file:

```
REACT_APP_API_URL=http://localhost:8080
```

4. Start the development server:

```bash
npm start
```

### Backend Setup

1. Navigate to the backend directory:

```bash
cd backend
```

2. Run the server:

```bash
go run cmd/server/main.go
```

### Docker Setup (Optional)

Build and run the containers:

```bash
docker-compose up --build
```

## Features

- User Authentication
- LeetCode Profile Integration
- Leaderboard System
- Progress Tracking
- Profile Management

## Tech Stack

### Frontend

- React
- TypeScript
- Tailwind CSS
- React Router
- React Query

### Backend

- Go
- [Add your backend frameworks/libraries]

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

[Add your license here]
