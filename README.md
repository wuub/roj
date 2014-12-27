# roj - opinionated container orchestration for growing infrastructures


## Concepts

* node - machine running *roj* daemon 
* roj cluster - any number of roj nodes connected to the same consul cluster
* app - a set of docker containers that can be launched multiple times 
* app instance - an app running on a single node


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

It will create **v1** version of **web** app. This app will use **nginx** image available on public docker registry, specifically tag **1.7.1**.  Once launched on a node, web:v1 app will be available on port 8080, forwarding all of the traffic to port 80 of nginx container.

## Listing all defined apps

Next we'll check that we have correctly defined our web:v1 app. 

When you use following command:

$ roj apps

you should see an output similiar to this one

[[TODO, show output of roj apps]]


## Adding a node to roj cluster

Creating apps is all nice and fine, but to be able to do any real work we need to launch them. Right now all we have is an app __definition__. To use it, we require a place where it will run, what we need is a **roj node**.

Lets see if we have one:

$ roj nodes

[[ TODO: empty output of roj nodes]]

Since we didn't launch any nodes yet, our roj cluster is predictably empty.

Fortunately it's very easy to fix. Open a **new terminal** and issue following command launching **roj node daemon**

$ roj agent

[[ TODO: output of default roj agent ]]

If everything goes well, roj will connect to local consul agent, register itself, display some helpful output and **block** waiting for orders to execute.

Leave the agent running and go back to the previous terminal. Try running   

$ roj nodes 

You'll notice that our cluster is no longer empty. Great job! Now we can start deploying our apps.


## Deploying an app

Once we are sure our cluster has both app definition and some nodes, we can begin doing some real work. Specifically deploying our beautiful apps to handle some real traffic or do real calculations. Lets do this now:

$ roj deploy web:v1 node1

This command will tell roj to deploy web:v1 application on node1. It's important to understand separation of concerns when it comes to deploying applications. What deploy command does is pretty simple, it validates that both app definition and node are correctly configured, and then modifies node1's metadata in consul KV store requesting that web:v1 app should be running there.

What it does not do is launch any containers or waits for the app to be available. 

Any image pulling and launching containers is fully under control of roj agent running on node1.


## Inspiration and other projects

* helios by Spotify is another project with very similar design decisions. While the basic idea of roj was created independently, during it's development I often checked how helios solved some of common problems. When applicable I also tried to use identical/similar terms to easy migration from one to another if/when necessary. [TODO: how and why roj is different]
