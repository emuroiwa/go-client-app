# Go Client Management App

A simple web application built with Go that allows users to create, list, and delete clients. This project demonstrates separation of concerns, in-memory data storage, form validation, and a layered architecture following SOLID principles.

---

## Features

- Add a client with name and email.
- Delete a client by ID.
- List all clients on the homepage.
- In-memory storage for quick prototyping.
- Email and input validation.
- Follows separation of concerns (Handlers, Services, Models, Validators).

---

## Project Structure

```bash
go-client-app/
│
├── handlers/ 
├── models/ 
├── services/ 
├── validators/
├── templates/ 
├── middleware/ 
├── utils/ 
├── tests/ 
├── main.go
└── README.md

```
---

## How to Run

1. **Clone the Repository**

   ```bash
   git clone
   cd go-client-app
   go run main.go

   go to http://localhost:8080  


Run tests 

go test ./tests


## Key Design Decisions
- Separation of Concerns: The application is broken into handlers, services, and models to maintain clean boundaries between layers.

- SOLID Principles: Each service has a single responsibility, and dependencies are injected via interfaces for better testability.


- Validation: Email and name validation are handled via a dedicated validators package to keep logic clean and reusable.