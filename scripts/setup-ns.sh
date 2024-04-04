#!/bin/bash

ip netns add client
ip netns add server

ip netns exec client ip link add eth0 type veth peer name eth0 netns server

ip netns exec client ip link set eth0 up
ip netns exec server ip link set eth0 up

ip netns exec client ip addr add 192.168.100.1/30 dev eth0
ip netns exec server ip addr add 192.168.100.2/30 dev eth0