# miner-dapp-wasm
Distributed app example in Golang for blockchain deployed as a WebAssembly utilizing a js websocket.



# Example run

- First navigate to root directory of main.go and compile it to webassembly.
`GOARCH=wasm GOOS=js go build -o main.wasm main.go`
- Run `node tracker.js`, this to establish communication among application instances.
- Duplicate the application folder to later run another instance of the application.
- Run `http-server` on both root directories of main.wasm to be served through `localhost:port`.
```
miner-dapp-wasm beber$ http-server
Starting up http-server, serving ./
Available on:
  http://127.0.0.1:8080
  
miner-dapp-wasm-2 beber$ http-server
Starting up http-server, serving ./
Available on:
  http://127.0.0.1:8082
```
- Type both URLs on two tabs of any web browser to open two instances of the application.
- Pick Alice on first instance and pick Bob on second one, now both wallets are initiated.
- On Bob's instance: click `Reward`.
 - For a new block to be created mining is taking time.
 
 ![Bob-instance-reward1-bob.png](https://ucarecdn.com/bf175e07-6954-4f66-becb-d4fe929864aa/)
 
 - This mining problem is solved by Bob's miner giving a nonce value of 1453337, this is explained thoroughly on part 2.
 - A new block is created showing the transaction, his networth increases by 10 while Alice is still having no crypsys.
 
 ![Alice-instance-reward1-bob-blockchainView.png](https://ucarecdn.com/40e9ca29-5ab2-4b36-95c0-949f581a618f/)
 
 - Data presented by each block: 
 
 |                                            |
 | ------------------------------------------ |
 | crypsys from -> crypsys to: amount crypsys |
 | Hash of previous Block                     |
 | Nonce                                      |
 | Hash of this Block                         |
 

 - Alice also shares the same blockchain data despite that she was not involved in the last transaction.
 - Now Bob decides to buy bananas from Alice.
 
 ![Bob-instance-buyBananas1-blockchainView.png](https://ucarecdn.com/df440403-5a76-47aa-bfb8-2c3e165531a1/)
 
 - His networth becomes 0 because he paid all of his 10 crypsys for the bananas (unwise move Bob !).
 - After few random buy and reward operations Alice buys apples from Bob.
 ![Alice-instance-buyApples1.png](https://ucarecdn.com/2cfac6aa-d78f-4d0b-a208-86aa9d649615/)
 
 - This time Alice is alerted that the nonce was mined by Bob despite that she initiated that operation. In this case the chainfabric notified Alice's miner to stop mining once the nonce was solved by Bob.
 
  ![Alice-instance-buyApples1-blockchainView.png](https://ucarecdn.com/5f3cfe8f-3b89-49f8-83ec-29fbb8933084/)
  
  - Now her networth is 0 as all she had before invoking that buy was 20 crypsys which is the price for Bob's apples.
