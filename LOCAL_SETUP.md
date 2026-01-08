# Running Ping Locally - Setup Guide

This guide will help you run Ping (customized Mattermost) on your local machine for development.

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go** (version 1.21 or later) - [Download](https://golang.org/dl/)
- **Node.js** (version 18.x or later) - [Download](https://nodejs.org/)
- **npm** (comes with Node.js)
- **Docker** and **Docker Compose** - [Download](https://www.docker.com/products/docker-desktop)
- **Make** - Usually pre-installed on macOS/Linux

## Quick Start

### 1. Start Docker Services (Database, MinIO, etc.)

First, start the required Docker containers (PostgreSQL, MinIO for file storage, etc.):

```bash
cd server
make start-docker
```

This will start:
- PostgreSQL database (port 5432)
- MinIO for file storage (port 9000)
- Inbucket for email testing (port 9001)
- Other development services

**Note:** Keep this terminal open or run it in the background. The services need to stay running.

### 2. Build and Run the Server

In a new terminal, navigate to the server directory and run:

```bash
cd server
make run-server
```

This will:
- Build the Go server
- Start the server on `http://localhost:8065`
- Set up necessary symlinks

### 3. Build and Run the Web Client

In another terminal, navigate to the webapp directory and run:

```bash
cd webapp
npm install
npm run build
npm run start
```

Or use the Makefile:

```bash
cd webapp
make run
```

This will:
- Install dependencies
- Build the React web app
- Start the development server (usually on port 9000 or configured port)

### 4. Access Ping

Open your browser and navigate to:
- **Web App**: `http://localhost:9000` (or the port shown in the terminal)
- **Server API**: `http://localhost:8065`

## Alternative: Run Everything Together

You can run both server and client with a single command:

```bash
cd server
make run
```

This runs both the server and web client together.

## First Time Setup

### Create Admin Account

1. Navigate to `http://localhost:8065` (or the web app URL)
2. You'll be prompted to create the first admin account
3. Fill in the required information:
   - Team name
   - Admin email
   - Admin username
   - Admin password

### Database Setup

The database will be automatically created and migrated when you first run the server. The default connection uses:
- **Host**: localhost
- **Port**: 5432
- **Database**: mattermost_test
- **User**: mmuser
- **Password**: mostest

These are configured in the Docker Compose setup.

## Development Workflow

### Making Changes

- **Backend (Go)**: Changes to Go files will require restarting the server (`make restart-server` or stop/start)
- **Frontend (React)**: Changes to React/TypeScript files will hot-reload automatically

### Restarting Services

```bash
# Restart server only
cd server
make restart-server

# Restart client only
cd webapp
make restart

# Restart both
cd server
make restart
```

### Stopping Services

```bash
# Stop server and client
cd server
make stop

# Stop Docker containers
cd server
make stop-docker
```

## Environment Variables

You can customize the setup using environment variables. Common ones:

```bash
# Server URL
export MM_SERVICESETTINGS_SITEURL=http://localhost:8065

# Database connection
export MM_SQLSETTINGS_DATASOURCE="postgres://mmuser:mostest@localhost:5432/mattermost_test?sslmode=disable"

# File storage
export MM_FILESETTINGS_DRIVERNAME=local
```

## Troubleshooting

### Port Already in Use

If you get port conflicts:

```bash
# Check what's using the port
lsof -i :8065  # Server port
lsof -i :9000  # Client port
lsof -i :5432  # PostgreSQL port

# Kill the process or change ports in config
```

### Docker Issues

```bash
# Restart Docker services
cd server
make stop-docker
make start-docker

# Or manually
docker compose down
docker compose up -d
```

### Database Connection Issues

```bash
# Check if PostgreSQL is running
docker ps | grep postgres

# Check database logs
docker logs mattermost-postgres
```

### Web Client Build Issues

```bash
# Clean and rebuild
cd webapp
rm -rf node_modules dist
npm install
npm run build
```

## Project Structure

```
ping-source/
├── server/          # Go backend server
│   ├── Makefile    # Server build/run commands
│   └── cmd/        # Server entry point
├── webapp/          # React frontend
│   ├── channels/    # Main web app code
│   └── package.json
└── docker-compose.yaml  # Docker services config
```

## Useful Make Commands

From the `server/` directory:

```bash
make run-server      # Run server only
make run-client      # Run web client only
make run             # Run both server and client
make stop            # Stop everything
make restart         # Restart everything
make start-docker    # Start Docker services
make stop-docker     # Stop Docker services
make build           # Build the server binary
make test            # Run tests
```

## Development Tips

1. **Hot Reload**: The web client supports hot reload for React changes
2. **Logs**: Server logs appear in the terminal where you ran `make run-server`
3. **Database**: Use a PostgreSQL client (like pgAdmin) to inspect the database
4. **File Storage**: Files are stored locally in `webapp/channels/dist/files/`
5. **Email Testing**: Use Inbucket at `http://localhost:9001` to view test emails

## Next Steps

- Check out the [API documentation](https://api.ping.com/)
- Explore the codebase structure
- Read the [contributing guide](./CONTRIBUTING.md)
- Join the development community

## Need Help?

- Check the [official documentation](https://docs.ping.com/)
- Review server logs for errors
- Check Docker container logs: `docker logs <container-name>`

Happy coding! 🚀

