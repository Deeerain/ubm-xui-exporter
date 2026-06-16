# UBM 3xui Exporter

A lightweight Go exporter for monitoring the 3xui panel with Prometheus.

## Overview

This exporter collects metrics from the 3xui API and exposes them in a Prometheus-compatible format.
It is designed to help integrate 3xui panel metrics into Prometheus-based monitoring and observability stacks.

## Features

- Fetches metrics from the 3xui API
- Exposes metrics at an HTTP endpoint for Prometheus scraping
- Includes built-in logging and configuration support
- Provides a sample `prometheus.yml` for quick setup

## Requirements

- Go 1.20+ (or compatible Go version)
- Access to the 3xui API endpoint
- Prometheus for scraping the exporter endpoint


Then configure Prometheus to scrape the exporter endpoint, for example:

```yaml
scrape_configs:
  - job_name: ubm_3xui
    static_configs:
      - targets: [localhost:8080]
```

## Docker Compose

If you build the image on GitHub and want to run it with `docker-compose`, use a configuration like this:

```yaml
services:
  ubm-xui-exporter:
    image: ghcr.io/deeerain/ubm-xui-exporter:latest
    container_name: ubm-xui-exporter
    ports:
      - 8080:8080
    environment:
      - LISTEN_ADDRESS=${LISTEN_ADDRESS:-:8080}
      - METRICS_PATH=${METRICS_PATH:-/metrics}
      - XUI_BASE_URL=${XUI_BASE_URL:-http://localhost:5353}
      - XUI_SECRET_PATH=${XUI_SECRET_PATH:-/}
      - XUI_ACCESS_TOKEN=${XUI_ACCESS_TOKEN:-abcd}
      - LOG_LEVEL=${LOG_LEVEL:-debug}
    restart: unless-stopped
```