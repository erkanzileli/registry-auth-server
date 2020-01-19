# registry-auth-server

This is a PoC of [Registry Token Auth Spec](https://docs.docker.com/registry/spec/auth/token/) document. This app now works on hard-coded user and permission list.

I want to create a configurable application. Configurations that I think

- Database: I think SQLite is an enough choice. But multiple Database compatibility may necessary.
- We may not want to open User API(User REST API). So all User CRUD operations may done over this CLI. It is just possibility.

## How it works?

You can look [Registry Token Auth Spec](https://docs.docker.com/registry/spec/auth/token/) for full information.
Registry application needs certificates when you want to use this with Token Authentication. For this reason you should have certificates. You can create dummy certificate like this

    sudo openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout server.key -out server.pem

But I already created and put it these to `ssl` directory.

Also I have an example Registry configuration file. In the root directory of registry-auth-server config.yml

You can run a registry instance like this. I give my all network interfaces to this container. This is easyway. If your 5000 port is unavailable then you should change 5000 port to another port which you want in config.yml file.

For Linux

    docker run \
    --detach \
    --name=registry \
    --network=host \
    -v `pwd`/ssl:/ssl \
    -v `pwd`/config.yml:/etc/docker/registry/config.yml \
    registry:2

For Mac

    docker run \
    --detach \
    --name=registry \
    --publish 5000:5000 \
    -v `pwd`/ssl:/ssl \
    -v `pwd`/config.yml:/etc/docker/registry/config.yml \
    registry:2

## Let's try

Login as `admin`

    docker login localhost:5000 -u admin -p 123qweasd

It works.

Before trying pull or push and image this registry you need to create an image which named as localhost:5000/`repository`:`tag`

Tag hello-world image

    docker pull hello-world:latest && docker tag hello-world:latest localhost:5000/hello-world:latest

Push

    docker push localhost:5000/hello-world:latest

Pull

    docker pull localhost:5000/hello-world:latest

It works. Try another user.

Login as `user`

    docker login localhost:5000 -u user -p password

Pull

    docker pull localhost:5000/hello-world:latest

It works.

Push

    docker push localhost:5000/hello-world:latest

It fails because `user` user has not push permission to hello-world repository.
