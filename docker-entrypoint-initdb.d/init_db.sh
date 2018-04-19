#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 -d book_development --username dev -f db/migrate/000000000_bootstrap.up.sql
psql -v ON_ERROR_STOP=1 -d book_development --username dev -f db/migrate/152373965_alterBooks.up.sql
psql -v ON_ERROR_STOP=1 -d book_development --username dev -f db/migrate/1523753980_AddPublisherReference.up.sql
psql -v ON_ERROR_STOP=1 -d book_development --username dev -f  db/migrate/1523808410_AddColumnsToBooks.up.sql

# better
# psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" -f  db/migrate/1523808410_AddColumnsToBooks.up.sql