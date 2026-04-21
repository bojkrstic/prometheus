# prometheus

## Pokretanje lokalno

Pokretanje cele platforme:

```bash
docker compose up --build
```

Servisi:

- App: `http://localhost:8080`
- Prometheus: `http://localhost:9090`
- Grafana: `http://localhost:3000`
- Alertmanager: `http://localhost:9093`

Grafana pristup:

- username: `admin`
- password: `admin`

## Šta se prati

Prometheus prati:

- lokalnu Go aplikaciju na `app:8080`
- `https://www.cerevac.rs/`
- `https://www.tecta.rs/`

Za eksterni monitoring koristi se `blackbox_exporter`.

## Alerti

Podešeni alerti:

- `WebsiteDown`
- `WebsiteSlowResponse`
- `WebsiteUnexpectedStatusCode`

Alerti se šalju na Telegram preko `Alertmanager`.

## Deployment na server

### 1. Instaliraj Docker i Docker Compose

Proveri da su dostupni:

```bash
docker --version
docker compose version
```

### 2. Prekopiraj projekat na server

Na primer:

```bash
scp -r prometheus user@server:/opt/prometheus-monitoring
```

Ili preko `git clone` ako repo postoji na Git serveru.

### 3. Podesi Telegram token

U fajlu `alertmanager.yml` upiši važeći Telegram bot token.

Ako je prethodni token bio deljen ili izložen, obavezno ga opozovi u `@BotFather` i generiši novi.

### 4. Pokreni servise

Iz root foldera projekta:

```bash
docker compose up --build -d
```

### 5. Provera da sve radi

Status kontejnera:

```bash
docker compose ps
```

Logovi:

```bash
docker compose logs prometheus --tail=100
docker compose logs alertmanager --tail=100
```

Provera dostupnosti interfejsa:

- Prometheus: `http://SERVER_IP:9090`
- Grafana: `http://SERVER_IP:3000`
- Alertmanager: `http://SERVER_IP:9093`

### 6. Automatsko podizanje posle restarta servera

Po želji možeš dodati restart politiku u `docker-compose.yml`, na primer:

```yaml
restart: unless-stopped
```

To je preporučeno za produkcioni server.

## Napomene

- Server mora imati izlaz na internet da bi mogao da proverava eksterne sajtove.
- Ako koristiš firewall, otvori portove `3000`, `9090` i `9093` po potrebi.
- Ako ne želiš javno izlaganje ovih portova, postavi reverse proxy ili ograniči pristup firewall pravilima.



• Checklist

  - Prebaci projekat na server, npr. u /opt/prometheus-monitoring
  - Proveri da su docker i docker compose instalirani
  - Potvrdi da je novi Telegram token upisan u alertmanager.yml:1
  - Pokreni servis: docker compose up --build -d
  - Proveri stanje: docker compose ps