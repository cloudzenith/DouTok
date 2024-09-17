#!/bin/sh

nohup /app/promtail -config.file /app/promtail.yaml > /app/logs/promtail.log 2>&1 &

/app/serve