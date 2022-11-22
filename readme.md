EndPoints

POST "/service/createCity" - add city
Body: {"Id":0,"Name":"","Contry":"","Сarbases":[],"Sort":0 }

POST "/service/createCarbase" - add carbase
Body: {"Id":0,"CityId":0,"Name":"","Vehicle":""}

GET "/service/getcity/{id}" - get city by id





City

{"Id":0,"Name":"1","Contry":"","Сarbases":[],"Sort":1 }
{"Id":0,"Name":"2","Contry":"","Сarbases":[],"Sort":2 }
{"Id":0,"Name":"3","Contry":"","Сarbases":[],"Sort":3}

{"Id":0,"Name":"1","Contry":"","Сarbases":[]}
{"Id":0,"Name":"2","Contry":"","Сarbases":[]}
{"Id":0,"Name":"3","Contry":"","Сarbases":[]}

Carbase

{"Id":0,"CityId":1,"Name":"q","Vehicle":""}
{"Id":0,"CityId":1,"Name":"w","Vehicle":""}
{"Id":0,"CityId":1,"Name":"e","Vehicle":""}
{"Id":0,"CityId":2,"Name":"q","Vehicle":""}
{"Id":0,"CityId":2,"Name":"w","Vehicle":""}
