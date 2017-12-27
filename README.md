whoami
======

Simple HTTP docker service that prints it's container ID, container name and node name

    $ docker run -d -p 8000:8000 -v /var/run/docker.sock:/var/run/docker.sock -v /etc/hostname:/etc/hostname --name whoami -t karek/whoami
