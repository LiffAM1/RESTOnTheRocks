#!/bin/bash
_cur_dir=$(pwd)
CONTAINER_NAME=mariadb
VOLUME_ARG="-v $_cur_dir/db/data:/var/lib/mysql"

if [ "$(basename "$_cur_dir")" != "rotr" ]; then
	echo "Run this from the rotr/ dir"
	exit 1
fi
case "$1" in
	start)
		if docker container inspect $CONTAINER_NAME > /dev/null 2>&1; then
			echo "restarting existing '$CONTAINER_NAME' container..."
			docker start $CONTAINER_NAME
		else
			if [ "$2" != "--persist" ]; then
				VOLUME_ARG=""
			fi
			echo "creating new '$CONTAINER_NAME' container..."
			docker run --name $CONTAINER_NAME $VOLUME_ARG -p 3306:3306 -e MYSQL_ROOT_PASSWORD=local -e MYSQL_DATABASE=rotr -e MYSQL_USER=rotr -e MYSQL_PASSWORD=rotrpwd -d mariadb:10.4
		fi
	;;
	stop)
		docker stop $CONTAINER_NAME
	;;
	rm|delete)
		docker stop $CONTAINER_NAME
		docker rm $CONTAINER_NAME
	;;
	status)
		docker ps -a -f "name=$CONTAINER_NAME"
	;;
	clean)
		docker stop $CONTAINER_NAME
		docker rm $CONTAINER_NAME
		rm -rf db/data/*
	;;
	*)
	echo "Usage: $0 { start | stop | [rm|delete] | status | clean}"
	cat << END
	start		creates and starts or restarts an existing $CONTAINER_NAME
	  --persist	persists the database to 'db/data'. can be cleaned up with 'clean'
	stop		stops an existing $CONTAINER_NAME
	rm|delete	deletes a running or stopped $CONTAINER_NAME
	status		runs 'docker ps'
	clean		deletes $CONTAINER_NAME and force deletes the dev database.
				requires sudo.
END
	;;
esac


#docker run --name mariadb -v `pwd`/db/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=local -d mariadb:10.4