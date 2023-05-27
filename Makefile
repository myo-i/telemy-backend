.PHONY: clean up

clean:
	docker-compose down 
	docker volume rm mysql_volume

up:
	docker-compose up -d

all: clean up