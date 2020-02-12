# docker run -it --name=client1 --network=host -d c
# docker run -it --name=client2 --network=host -d c
docker run -it --name=server1 -p 7777:7777 -d s
docker run -it --name=server2 -p 8888:7777 -d s
docker run -it --name=server3 -p 9999:7777 -d s  

# docker stop client1 &
# docker stop client2 &
docker stop server1 &
docker stop server2 &
docker stop server3 &



xfce4-terminal -e 'docker start -a server1' &
xfce4-terminal -e 'docker start -a server2' &
xfce4-terminal -e 'docker start -a server3' &
xfce4-terminal -e 'docker start -a client' &