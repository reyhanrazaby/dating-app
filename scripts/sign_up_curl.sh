curl --location 'localhost:4545/sign-up' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "reyhan@mail.com",
    "password": "1234",
    "gender": "M",
    "full_name": "Reyhan Razaby",
    "date_birth": "22-12-1994",
    "location_lat": -6.8059399341534075,
    "location_lng": 106.95161161606342,
    "bio": "Aku anak sehat"
}'