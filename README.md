# Fantasy Cricket Application

## Overview

This application is a fantasy cricket platform where users can build teams and score points based on the performance of real players. The application includes:
- A pool of players to choose from
- Constraints on team selection
- An administrator interface to enter points scored by each player
- User interface to check total points after a match

## Installation

### Prerequisites
- Go 1.16 or later
- Git

### Clone the Repository
```sh
git clone https://github.com/rishabhjain02/fantasyCricket.git
cd fantasyCricket
```

## Running the Application
```sh
go build
./fantasyCricket
```

## Project Structure
```sh
fantasyCricket/
├── entities/
│   ├── player.go        # Player entity definition
│   └── team.go          # Team entity definition
├── delivery/
│   └── http.go  # HTTP handlers for API endpoints
├── repository/
│   ├── playerRepository.go  # In-memory repository for players
│   └── teamRepository.go    # In-memory repository for teams
├── usecase/
│   └── teamService.go   # Business logic for creating teams and calculating points
├── main.go                # Entry point of the application
```

## Endpoints

### 1. Create Team
Endpoint
```sh
POST /create-team
```
Parameters
* user_id: User ID (integer)
* player_ids: All player ids

Example
```sh
curl -X POST -d "user_id=1&player_ids=1&player_ids=2&player_ids=3&player_ids=4&player_ids=5&player_ids=6&player_ids=7&player_ids=8&player_ids=9&player_ids=10&player_ids=11" http://localhost:8080/create-team
```

### 2. Get Team Points
Endpoint
```sh
GET /team-points
```
Parameters
* user_id: User ID (integer)

Example
```sh
curl -X GET "http://localhost:8080/team-points?user_id=1"
```

### 3. Update Player Points
Endpoint
```sh
POST /update-player-points/{player_id}
```
Parameters
* player_id: Player ID (integer)
* points: Points scored by the player (integer)

Example
```sh
curl -X POST -d "points=40" http://localhost:8080/update-player-points/1
```

## Assumptions
1. **Player Pool:** Assumed to be pre-populated. There is no flow for adding players.
2. **User Identification:** Users are identified by user ID. No authentication or user management system is implemented.
3. **Points Calculation:** The administrator updates Points directly without validation.
4. **Budget Constraint:** Each team has a budget constraint of 100 credits, enforced during team creation.
5. **Team Constraints:** Each team must have 11 players with at least 1 wicket-keeper, 3 batsmen, and 3 bowlers.

## Trade-offs
1. **In-Memory Storage:** For simplicity and demonstration purposes, the application uses in-memory storage for players and teams. For production use, persistent storage (e.g., a database) can be used.
2. **Error Handling:** Basic error handling is implemented. More comprehensive error handling and logging can be implemented for a production system.
3. **Scalability:** To scale the application in the future, concurrent access, distributed storage, and load balancing can be implemented.
4. **User Interface:** A simple CLI-based interaction is implemented. A full-fledged web or mobile UI would enhance user experience.

## Conclusion
This documentation provides a clear guide to setting up, running, and understanding the fantasy cricket application.
The design follows clean architecture principles, separating concerns across different layers. 
While the current implementation focuses on core functionalities, it is structured for extensibility and future enhancements.
