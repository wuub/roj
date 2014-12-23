# roj - opinionated container orchestration for growing infrastructures


## Concepts

* node - machine running *roj* daemon 
* roj cluster - any number of roj nodes connected to the same consul cluster
* app - a set of docker containers that can be launched multiple times 
* instance - an app running on a single node


## Preparing a Consul cluster 

__you may skip this paragraph if you already have a running consul cluster__

Roj uses consul for all of its features: node discovery, membership, access control, persistent metadata storage, change notifications and redeploys.

As such having a correctly configured consul cluster is a requirement. 

Fortunately running one is neither difficult nor heavy on the resources. As far as roj is concerned a single node consul cluster is perfectly fine. 

```bash
mkdir /var/consul
consul agent --server --bootstrap --data-dir=/var/consul -advertise=127.0.0.1
```

Obviously, such deployment suffers from all kinds of availability issues but running roj on a single VM/VPS is definitely **supported**.  Smooth migration from single-node to multi-node to multi-dc deployment is one of roj design aims. 


## Defining an app

You can create a single container app using the following command.

$ roj create web:v1 nginx:1.7.1 -p 80:8080 

It will create **v1** version of **web** app. This app will use **1.7.1** tag of **nginx** image available on public docker registry.  Once launched on a node, web:v1 app will be available on port 8080, forwarding all of the traffic to port 80 of nginx container.

## Listing all defined apps

Next we'll check that we have successfuly defined our web:v1 app. 

When you issue 

$ roj apps

you should see an output similiar to this one

[[TODO, show output of roj apps]]


## Adding a node to roj cluster

Creating apps is all 









## Inspiration and other projects

* helios by Spotify is another project with very similar design decisions. While the basic idea of roj was created independently, during it's development I often checked how helios solved some of common problems. When applicable I also tried to use identical/similar terms to easy migration from one to another if/when necessary. [TODO: how and why roj is different]
