# Hippo Refactor
Refactoring hippo-webapp using a new tech stack:
- Frontend - Svelte/Tailwind
- Backend - Go
- Database - MySQL

# Development
- Server (`/server`)
    - Start server: `go run .`
    - Output logs: `tail -f server.log`

- UI (`/app`)
    - Run: `npm install` if necessary
    - Start dev: `npm run dev`

- Try to keep code under 100 lines by wrapping arguments/parameters

# Logs
- Logs are output to the root of the folder in `server.log` by default.

- Run `tail -f server.log` to continuously output server logs.

- If running from the docker, you can tail the logs by first running `docker exec -it {CONTAINER_ID} //bin/sh`

# Dependencies
- Go 
    - [Viper](https://github.com/spf13/viper) - `go get github.com/spf13/viper`
    - [Gorilla/Mux](https://github.com/gorilla/mux) - `go get -u github.com/gorilla/mux`

- UI
    - [Svelte](https://svelte.dev/) - https://svelte.dev/blog/the-easiest-way-to-get-started
        - TypeScript enabled
    - [Tailwind](https://tailwindcss.com/) - `npm install tailwindcss@npm:@tailwindcss/postcss7-compat postcss-cli@^7 autoprefixer@^9 concurrently cross-env --save-dev`