create database lanjukang;

migrate -path db/migration -database "postgresql://din:postgres@localhost:5432/lanjukang?sslmode=disable" -verbose up 

jika terjadi masalah pada kode 
- fix masalahnya
- migrate -path db/migration -database "postgresql://din:postgres@localhost:5432/lanjukang?sslmode=disable" -verbose force angka(misalnya 1)
- migrate -path db/migration -database "postgresql://din:postgres@localhost:5432/lanjukang?sslmode=disable" -verbose up

