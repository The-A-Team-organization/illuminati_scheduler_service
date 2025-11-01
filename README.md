# Illuminati Scheduler Service

This microservice periodically calls specific backend API endpoints via HTTP requests according to a schedule, using the [`robfig/cron`](https://github.com/robfig/cron) library.

---

## Functionality

Currently, the main task implemented:

- `CloseVotes` — sends a `PATCH` request to the backend at `http://host.docker.internal:8000/api/votes/vote_close`  
  with the parameter `date_of_end` in the format `YYYY-MM-DD HH:MM:SS`.

- `SetInquisitor` — sends a `PATCH` request to the backend at `http://backend:8000/api/votes/manage_inquisitor/`  
  to set the inquisitor.

- `UnsetInquisitor` — sends a `DELETE` request to the backend at `http://backend:8000/api/votes/manage_inquisitor/`  
  to unset the inquisitor.

- `BanArchitect` — sends a `DELETE` request to the backend at `http://backend:8000/api/ban_architect/`  
  with the parameter `architect_id` and `reason`.

- `NewEntryPassword` — sends a `POST` request to the backend at `http://backend:8000/api/entry_password/`
