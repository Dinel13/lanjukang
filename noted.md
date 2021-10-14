create database lanjukang;

migrate -path db/migration -database "postgresql://din:postgres@localhost:5432/lanjukang?sslmode=disable" -verbose up 

jika terjadi masalah pada kode 
- fix masalahnya
- migrate -path db/migration -database "postgresql://din:postgres@localhost:5432/lanjukang?sslmode=disable" -verbose force angka(misalnya 1)
- migrate -path db/migration -database "postgresql://din:postgres@localhost:5432/lanjukang?sslmode=disable" -verbose up

key is of invalid type karena pake secret key string harusnya []byte

token tidak mau diparse karena interface{}
claims["id"].(float64) untuk ubahh interfcae{} menjadi float64
lalu ubah float64 menjadi int dengna int(claims["id"].(float64)) 

