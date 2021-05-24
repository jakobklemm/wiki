---
title: Traefik Docker (Compose) + Loadbalancer + TLS
description: 
published: true
date: 2021-05-24T18:47:15.071Z
tags: 
editor: markdown
dateCreated: 2021-05-24T18:46:02.943Z
---

# Compose
Important: acme.json must have permission level 600.
``` yml
version: "3.2"
services:
  router:
    image: traefik:v2.3
    container_name: "traefik"
    networks:
      - "web"
    ports:
      - 80:80
      - 443:443
      - 8080:8080
    command:
      - "--api=true"
      - "--api.dashboard=true"
      - "--log.level=INFO"
      - "--providers.docker=true"
      - "--providers.docker.exposedbydefault=false"
      - "--providers.docker.network=web"
      - "--entrypoints.web.address=:80"
      - "--entrypoints.websecure.address=:443"
      - "--providers.file.filename=/etc/traefik/dynamic.yml"
      - "--entrypoints.web.http.redirections.entryPoint.to=:443"
      - "--entrypoints.web.http.redirections.entryPoint.scheme=https"
      - "--certificatesresolvers.myresolver.acme.httpchallenge=true"
      - "--certificatesresolvers.myresolver.acme.httpchallenge.entrypoint=web"
      - "--certificatesresolvers.myresolver.acme.email=web@jeykey.net"
      - "--certificatesresolvers.myresolver.acme.storage=/etc/traefik/acme.json"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - $PWD/acme.json:/etc/traefik/acme.json
      - $PWD/dynamic.yml:/etc/traefik/dynamic.yml
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.api.service=api@internal"
  karinklemm:
    image: nginx:karinklemm
    container_name: "karinklemm"
    networks:
      - "web"
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.karinklemm.rule=Host(`karinklemm.ch`)"
      - "traefik.http.routers.karinklemm.entrypoints=web,websecure"
      - "traefik.http.routers.karinklemm.tls=true"
      - "traefik.http.routers.karinklemm.tls.certresolver=myresolver"
  jeykey:
    image: nginx:jeykey
    container_name: "jeykey"
    networks:
      - "web"
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.jeykey.rule=Host(`jeykey.net`)"
      - "traefik.http.routers.jeykey.entrypoints=web,websecure"
      - "traefik.http.routers.jeykey.tls=true"
      - "traefik.http.routers.jeykey.tls.certresolver=myresolver"
networks:
  web:
    external: true
```
# Dynamic
``` yml
http:
  routers:
    orion:
      entryPoints:
        - "web"
        - "websecure"
      rule: "Host(`orion.jeykey.net`)"
      service: orion
      tls:
        certResolver: myresolver
    wiki:
      entryPoints:
        - "web"
        - "websecure"
      rule: "Host(`wiki.jeykey.net`)"
      service: wiki
      tls:
        certResolver: myresolver
  services:
   orion:
      loadBalancer:
        servers:
        - url: "http://192.168.1.27:3001"
   wiki:
      loadBalancer:
        servers:
        - url: "http://192.168.1.27:3002"

```