docker-compose -f visor.yml down
docker rm $(docker ps -a -q) 
docker image rm $(docker image ls -a -q)
docker volume rm $(docker volume ls -q)
sh visor.sh # $1 $2