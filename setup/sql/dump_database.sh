#!/usr/bin/env bash
pg_dump --create --schema-only -U postgres primetrust > "primetrust-$(date +%F).sql"
