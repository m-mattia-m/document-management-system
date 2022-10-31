# Document-Management-System

## Getting started

1. Start container: `docker compose -f docker-compose.yml -f docker-compose.dev.yml up -d`
2. .env-File

    ```env
    APP_ENV= # PROD, TST, ABT, DEV, ...
    APP_PORT=
    APP_HOST=
    
    DB_HOST=
    DB_PORT=
    DB_ROOT_USERNAME=
    DB_ROOT_PASSWORD=
    DB_NAME=
    DB_USERNAME=
    DB_PASSWORD=
    
    ADMIN_USERNAME=
    ADMIN_PASSWORD=
    ```
