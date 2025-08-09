запуск проекта:<br>
docker-compose.build.yml - файл для запуска проекта в build-версии<br>
docker-compose.yml - файл для запуска проекта в dev-версии (поддерживает hot-reload)

команда для вывода кошельков контейнера go-rest-api-db-1:<br>
docker exec -i go-rest-api-db-1 psql -U postgres -d wallet_db -c "SELECT id, address, balance FROM wallets;"
