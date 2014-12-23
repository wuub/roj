# roj - opinionated Docker orchestration for growing infrastructures


## Concepts

* node - machine running *docker*, *consul* and *roj* daemons 
* app - a set of docker containers that can be launched multiple times 
* instance - an app running on a single node


## Preparing a Consul cluster 

__you may skip this paragraph if you already have a functional consul cluster__

Roj uses consul for all of its features: node discovery, membership, access control, persistent metadata storage, change notifications and redeploys.

As such having a correctly configured consul cluster is a requirement. 

Fortunately running one is neither difficult nor heavy on the resources. As far as roj is concerned a single node consul cluster is perfectly fine. 

```bash
mkdir /var/consul
consul agent --server --bootstrap --data-dir=/var/consul -advertise=127.0.0.1
```

Obviously, such deployment suffers from all kinds of availability issues but running roj on a single VM/VPS is definitely **supported**.  Smooth migration from single-node to multi-node to multi-dc deployment is one of roj design aims. 


## Defining an app

You can create a single constainer app using following command.

$ roj create web:v1 nginx:1.7.1 -p 80:8080 

This will define **v1** version of **web** application. This app will use 1.7.1 tag (version) of **nginx** image available on public docker registry.










## Inspiration and other projects

* helios by Spotify is another project with very similar design decisions. While the basic idea of roj was created independently, during it's development I often checked how helios solved some of common problems. When applicable I also tried to use identical/similar terms to easy migration from one to another if/when necessary. [TODO: how and why roj is different]
