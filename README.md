# zafirabot
## Server
### Requirements to run
- `cd server && go mod tidy`
- install [air](https://github.com/air-verse/air)
- run `make server`
## client
- `npm install`
- `npx vite --host`
## nginx
- add `rtmp.conf` to `nginx.conf`
- copy `robot.conf, client.conf, server.conf` to `sites-available`
## rpi
- enable services in systemd
