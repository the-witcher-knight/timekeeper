# Timekeeper

> The application is designed to streamline employee attendance tracking using blockchain technology. It consists of a Solidity-based smart contract and a robust backend service to ensure secure, efficient, and transparent attendance management

## How to run

Prerequisites:

- Install [Docker Desktop](https://www.docker.com/products/docker-desktop/) or Docker on your local machine.
- Install [make](https://www.gnu.org/software/make/) on your local machine. 

> Some commands in the Makefile might not work on Windows. Please refer to the Makefile to find the commands used in each step and type those commands into your CLI, such as (`cmd`)

Steps:

1. Run `make local-env` (you can copy value from `.env.sample` to `.env.local` also)
2. Run `make setup`
3. Open `timekeeper-ganache` logs on docker container and look for a privates key, copy & paste it to `.env.local`
4. Run `make deploy-contract` to deploy our smart contract to ganache \
You will see the response as below
```
{"level":"info","ts":1718988591.9327233,"caller":"logging/logger.go:48","msg":"Contract Deployed","contract.address":"0xC93B61C6B44A52C0d608dbEa0b359C3A3e6D0921","contract.tx_hash":"0x388be2240622de6c8dc1c3c256f386a34b6b192b1c814d34ff10b66f01ac9f2d"}
```

5. Copy the `contract.address` value and put it to `.env.local`.
6. Run `make serve` to run our API services.

> Notice: Every address have prefix `Ox`, please remove it before adding them to `.env.local`.

## Contributors

<img src="https://avatars.githubusercontent.com/u/44874068?v=4" style="width: 50px; height: auto" alt="the-witcher-knight">
