# DNS bug

## Setup

```
docker-compose build
```

## Start

```
docker-compose up -d
```

Open a browser on `http://localhost`. Refresh. It should say `I see 1 backend container(s)`.

## Scale the backend up

```
docker-compose scale back=20
```

Open a browser on `http://localhost`. Refresh. It should say `I see 20 backend container(s)`.

## Scale the backend down

```
docker-compose scale back=1
```

Open a browser on `http://localhost`. Refresh. It should say `I see 1 backend container(s)`.

## Initialize swarm

```
docker swarm init
```

## Scale the backend up

```
docker-compose scale back=20
```

Open a browser on `http://localhost`. Refresh. It should say `I see 20 backend container(s)`.

## Scale the backend down

```
docker-compose scale back=1
```

Open a browser on `http://localhost`. Refresh. It should say `I see 1 backend container(s)`.
