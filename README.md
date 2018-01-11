# go-storm
Golang PSQL ORM for Price Probe.

# Run with Docker Compose:

 - Install and Start Docker: `yaourt -S docker`, `sudo systemctl start docker`
 - Install Docker Compose: `yaourt -S docker-compose`
 - Run `docker-compose up`.

## Other useful commands

### Docker

 - Run with no output `docker-compose up -d`
 - Stop with `docker-compose down`
 - Stop and remove images with `docker-compose down --rmi all`
 - Open Shell on one container with `docker exec -i -t container_id /bin/bash`

### Postgresql

  - Dump: `sudo -u postgres pg_dump -a -h xx.xxx.xxx.xxx -p xxxx -U postgres -d priceprobe > path_to_store_backup.sql`
  - Restore: `sudo -u postgres psql -h xx.xxx.xxx.xxx -p xxxx -U postgres -d priceprobe < path_to_backup.sql`

# Run with Goland

  - Start PostgresSQL: `sudo systemctl start postgresql`
  - Create a Database called `priceprobe`
  - Edit your Run configuration program's arguments by passing `-environment=development`