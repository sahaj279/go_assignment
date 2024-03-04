# HLD

## Problem

Develop a multi-threaded go program where one thread reads the data from the database say, details of an Item from a mysql table. This thread builds an in-memory object, stores it in a collection. Simultaneously another thread should fetch already created Item objects from this collection and calculate the tax as per rules detailed in assignment#1 update the tax value in appropriate Item attribute and store it in a different collection. Finally print out the item details to console as detailed in assignment #1.


## Design
![Your paragraph text](https://github.com/sahaj279/go_assignment/assets/88133213/4d5765e9-5605-44ac-ae5b-ff8975a55842)

### Requirements
- An in memory collection 1 to store data from mysql
- A channel to store items from collection1 and pass it to other thread
- Go thread 1 that will read data from mysql and store it in a channel
- Go thread 2 that will read data from channel

### Flow
- Thread 1 reads data from mysql to store it in collection 1 and then in channel
- Simultaneously thread 2 reads from channel and calculates tax on it and prints invoice

