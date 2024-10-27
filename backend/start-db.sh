docker run --name postgres \
    -e POSTGRES_USER=admin \
    -e POSTGRES_PASSWORD=password \
    -e POSTGRES_DB=hive \
    -v postgres_data:/var/lib/postgresql/data \
    -p 5432:5432 \
    -d postgres;