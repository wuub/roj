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

## Checking your roj version

Install roj on your machine by placing roj binary somewhere on your PATH.  

Check if it's correctly installed by using 

$ roj version



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

[[ TODO: output of default roj deploy ]]

Roj deploy will output instance-id. It's a key that uniquely identifies app instance int the roj cluster.


This command will tell roj to deploy web:v1 application on node1. It's important to understand separation of concerns when it comes to deploying applications. What deploy command does is pretty simple, it validates that both app definition and node are correctly configured, and then modifies node1's metadata in consul KV store requesting that web:v1 app should be running there.

What it does not do is launch any containers or waits for the app to be available. Any image pulling, reconfiguration, launching containers, port allocation, and so on and so forth if full under control roj agent control running on node1. In most cases changes should be visible almost instantly, but under heavy contention or when node1 is partitioned from consul quorum, it might take it arbitrarily long to launch app instance.

To know that, we need to use a different command

## Listing app instances and checking their state

To check if our instance launched successfully look at the metadata.

$ roj instances node1

[[ TODO: output of default roj instances ]]

Look for "launched" state. Roj doesn't track if the app is running, relying instead on the docker daemon, but it does know if it was able to launch is successfully.


## Upgrading an app


App upgrades are the most important part of roj. While new apps are rarely deployed on a daily basis, it's quite common in modern organization to see dozens of upgrades per day per app.

Upgrade paths can be complex, with different requirements (rolling, staged, sharded deploys) and constraints (synchronized deploys, maintenance windows).

Downgrades and rollbacks are important, and can have very different time constraints from normal upgrades. 

While Roj does not attempt to solve those problems for you, it does try to make them more manageable by providing well defined building blocks and promises. If used correctly, you will be able to reason about the state of your system even after partitions or unexpected node outage/recovery. 

# Declarative nature of roj

Many of roj's constraints, promises and behaviors can be directly traced to it's declarative nature. Most actions described in the getting started tutorial operate on cluster metadata only, describing a state we want the cluster to be in, not necessarily performing any actions leading to this state. 

Each roj agent will periodically perform anti-entropy on the node it's managing. If for any reason current state differs from the one required by cluster metadata, agent will perform any actions required to synchronize actual state with the one required. Anti-entropy does not differentiate between upgrades, downgrades, deploys, undeploys or overriding operator manual actions. 

Declarative nature makes roj more indirect and a bit more tricky to understand in the happy-case. It also makes it much easier to reason about in the event of partial failure. It allows you to make changes to node's metadata when the machine itself is currently offline, guaranteeing anti-entropy as soon as it's back online.

[TODO: describe operation modes in other common failure scenarios]

# Immediate nature of roj

Although we do like 





## Anti-entropy
# Roj agent pull mode

## Garbage collection
# Containers
# Images

## Events




## Inspiration and other projects

* helios by Spotify is another project with very similar design decisions. While the basic idea of roj was created independently, during it's development I often checked how helios solved some of common problems. When applicable I also tried to use identical/similar terms to easy migration from one to another if/when necessary. [TODO: how and why roj is different]
