#Clone Repos
#cd ..
#git clone https://github.com/cesarsasi/visor-mapas-backend
#git clone https://github.com/cesarsasi/visor-mapas-frontend

#GO TO $1 BRANCH Front, $2 backend

# cd ../visor-mapas-frontend
# git pull
# git checkout $1

# cd ../visor-mapas-backend
# git pull
# git checkout $2

cd ..

#Apply configurations
# cp ./visor-mapas-scripts/.envconfig.yml ./visor-mapas-backend/config/config.yml
# cp ./visor-mapas-scripts/.envfrontprod ./visor-mapas-frontend/.env.production
# cp ./visor-mapas-scripts/.envfrontdev ./visor-mapas-frontend/.env.development

#Build images
pwd
ls

# docker build ./visor-mapas-scripts -t geoserver
# docker build ./visor-mapas-backend -t visor-backend
# docker build ./visor-mapas-frontend -t visorfrontend

# #Clean repos
# #rm -rf visor-mapas-backend
# #rm -rf visor-mapas-frontend

# #Compose up. Add -d at the end to hide docker logs
# #Geoserver
# docker-compose -f ./visor-mapas-scripts/geoserver.yml up -d 
# #Visor
docker-compose --env-file ./visor-mapas-scripts/.env -f ./visor-mapas-scripts/visor.yml up -d --build
