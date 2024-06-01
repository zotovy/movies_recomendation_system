#!/bin/zsh

#docker exec -i postgres_movies bash < ./create_tables.sh
#docker exec -i postgres_movies bash < ./load_datasets.sh
#docker exec -i postgres_movies bash < ./setup_fts.sh
docker exec -i postgres_movies bash < ./setup_indexes.sh