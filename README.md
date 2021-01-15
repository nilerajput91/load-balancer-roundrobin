# load-balancer-roundrobin
A simple round robin load balancer implemented in Go

Here we supply a list of routes to servers that generate a load balancer that uses round robin selection to route requests to each server.

Example

go run any number of servers from /server
go run loadbalancer/loadbalancer.go
Now when you head over to localhost:8080. Based on the servers started it should show a new port every time you refresh the page.

Methodology Use for Loadbalancer

It works by looping through all available servers using round robin methodology and then routes the request to the appropriate server.

In a way using the load balancer here is fault tolerant. In the case a server goes down, the load balancer routes the request to the next available server. Logs are shown in loadbalancer.go to show the routing mechanisms.

Health checks are done in real time during every route to determine which servers are alive or not. Only in the case when all the given servers are down, it returns a http.StatusServiceUnavailable status back to the client.
