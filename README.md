whoami
======

Run the container:

    $ docker run -d -p 8000:8000 -v /var/run/docker.sock:/var/run/docker.sock -v /etc/hostname:/etc/hostname --name whoami -t karek/whoami
