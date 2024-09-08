-- name: CreateDoctor :one
INSERT INTO doctors(name,email,phone,password,gender,dob,img,shedule,degree,address,roles,category_id)
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
RETURNING *;

-- name: CreateDoctor :one
