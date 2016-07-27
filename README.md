openssl genrsa -out app.rsa 1024

openssl rsa -in app.rsa -pubout > app.rsa.pub


mongod: error while loading shared libraries: libssl.so.10: cannot open shared object file: No such file or directory

http://localhost:8080/users/register

{
	"data" : {
		"firstname" : "Rodolfo",
		"lastname" : "Guzman Huerta",
		"email" : "santo@santo.com",
		"password" : "asdfasdf"
	}
}


http://localhost:8080/users/login

{
	"data" : {
		"email" : "santo@santo.com",
		"password" : "asdfasdf"
	}
}


docker run -d -p 27017:27017 -h santo --name mongo-container mongo

sudo docker exec -it 479 mongo

docker run -d -p 8080:8080 --name go-api --link mongo-container docker-image-gosanto
