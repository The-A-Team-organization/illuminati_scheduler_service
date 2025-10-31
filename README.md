# Illuminati Scheduler Service

This microservice periodically calls specific backend API endpoints via HTTP requests according to a schedule, using the [`robfig/cron`](https://github.com/robfig/cron) library.

---

## Functionality

Currently, the main task implemented:

- `CloseVotes` — sends a `POST` request to the backend at `http://host.docker.internal:8000/api/votes/vote_close`  
  with the parameter `DateOfEnd` in the format `YYYY-MM-DD HH:MM:SS`.

---

## TODO

Add other endpoints (SetInquisitor, NewPassword)
