docker stop $(docker ps | grep 8000 | awk '{print $1}')
