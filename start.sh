#!/bin/sh

alembic upgrade head

uv run src/main.py