#!/bin/bash
set -e;

# a default non-root role
MONGO_NON_ROOT_ROLE="${MONGO_NON_ROOT_ROLE:-readWrite}"

if [ -n "${MONGO_NON_ROOT_USERNAME:-}" ] && [ -n "${MONGO_NON_ROOT_PASSWORD:-}" ]; then
	"${mongo[@]}" "$MONGO_INITDB_DATABASE" <<-EOJS
		db= db.getSiblingDB($(_js_escape "$MONGO_APP_DATABASE"))
		db.createUser({
			user: $(_js_escape "$MONGO_NON_ROOT_USERNAME"),
			pwd: $(_js_escape "$MONGO_NON_ROOT_PASSWORD"),
			roles: [ { role: $(_js_escape "$MONGO_NON_ROOT_ROLE"), db: $(_js_escape "$MONGO_APP_DATABASE") } ]
			});
		
	EOJS
fi

mongorestore --host=localhost --port=27017 --username=${MONGO_NON_ROOT_USERNAME} --password=${MONGO_NON_ROOT_PASSWORD} --authenticationDatabase=${MONGO_APP_DATABASE} --db=${MONGO_APP_DATABASE} /opt/dump