
## Install Task

- [Task](https://taskfile.dev/)

## How to Run

sudo task setup-ns
sudo task run-server
sudo task curl <--- Run this in another terminal

## Clean up

sudo task cleanup-ns

## How to sniff the traffic

sudo ip netns exec client tshark -i eth0
sudo ip netns exec server tshark -i eth0