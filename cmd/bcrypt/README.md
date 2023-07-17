run to generate the hash:

go run cmd/bcrypt/bcrypt.go hash "secret" 

copy the hash and run the below, with the hash that was generated in the previous step

go run cmd/bcrypt/bcrypt.go compare "secret" '$2a$10$pnJ.VqfFeoi.p3iufDrkKuFnJr/zzBXZ7L/B6CzCpbo5VMWvzIYym'