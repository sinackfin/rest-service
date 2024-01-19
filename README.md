docker-compose up -d
go run cmd/api/main.go
curl -XPOST -d '{"name":"Vasil"}' 'http://127.0.0.1:8080/api/person' -- add person
curl  'http://127.0.0.1:8080/api/person?id=b817dcc6-b6b8-11ee-a4fc-9783a89cbc9d' -- get person by uuid
curl -XDELETE 'http://127.0.0.1:8080/api/person/b817dcc6-b6b8-11ee-a4fc-9783a89cbc9d' -- delete person
curl -XPATCH  -d '{"name":"Vitalya"}' 'http://127.0.0.1:8080/api/person/b817dcc6-b6b8-11ee-a4fc-9783a89cbc9d' -- update person
