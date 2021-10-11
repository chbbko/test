#run
docker-compose -f docker-compose.yml up -d
docker-compose up -d postgres && docker-compose up --build test-apis
