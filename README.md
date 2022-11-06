# ETHSF2022-Hackathon-zkDVC
ETH SF 2022 Hackathon project - zkDVC on Polygon ID 

<img width="954" alt="Screen Shot 2022-11-06 at 8 28 34 AM" src="https://user-images.githubusercontent.com/7462849/200195179-4c285ac4-e198-4d9d-94c3-f5f52c34c04d.png">



### zkDVC - ZK Dynamic Verifiable Credential Service

zkDVC is a Threshold-based Credit Issuance Template Module on Polygon ID. The primary goal of this initiative is to allow on-chain creators to leverage users' behavioral data to evaluate the record and issue credits score accordingly. It enables a new way of credit issuance and cross-platform action/status verification.

![addresses](https://user-images.githubusercontent.com/7462849/200195099-ab53d251-559c-437e-8b40-36d0988a5fb8.png)


zkDVC focuses on 3 technical goals in this platform:
1.  Template-based claim schema extension allows multiple threshold-type claims to be packed into one offer.
2.  Support continuous tracking for transparent revocation.
3.  Multi-addresses data aggregating and credit issuing with anonymity (To be continued)



#### Project Brief
*   Issuer server, which is used to manage claims, in which go language is used to build the project framework, iden3 is used to manage the issuer’s identity and credentials, and zk is used to construct proofs and verification. Claims are organized according to the iden3 protocol, and each user includes three Merkle trees, which are stored in postgresql.
*   The back-end part of wallet (wallet-end) is designed with Go + iden3. Go is used as the project framework to provide an api server for wallet-front; the iden3 protocol is used to manage the holder’s claims and identity. The specific data is stored in the iden3 protocol according to the in postgresql. Wallet-end will also call a Badge contract (ERC-721), provide proof of its own claims and obtain Badge.
*   The front-end of wallet (wallet-front) is implemented by react, and its functions are obtained from wallet-end.
*   By using the Dynamic Verifiable Credentials (DVC) powered by KNN3, we can verify user accounts based on real-time data on the chain. DVC makes this project possible, and in the future, we can even aggregate and authenticate different assets of different accounts of the same user through DVC.

<img width="1097" alt="Screen Shot 2022-11-06 at 1 01 53 AM" src="https://user-images.githubusercontent.com/7462849/200195124-93e15e5e-8ae1-4e37-8612-ec65af8c3540.png">


### Introducory deck
https://drive.google.com/file/d/116G7nltJTGa8T1Dg7i92Fdgy02W8M2Ew/view?usp=share_link


## Demo
### Claim template anage Platform
http://34.212.245.164:4000/home/claim

### zk Wallet for scan-to-claim
http://34.212.245.164:3000/


### Test wallet address
0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266

### NFT List
- BasementDwellers (DWEL) （已使用）https://etherscan.io/address/0x9a95ecee5161b888ffe9abd3d920c5d38e8539da：0x9a95ecee5161b888ffe9abd3d920c5d38e8539da
- Slang Skulls - World Skull Art （已使用） https://opensea.io/collection/slangskullsworldskullart：0x495f947276749ce646f68ac8c248420045cb7b5e
- Bear Market Buds (BMBUDS)  （已使用）https://etherscan.io/address/0x48a0501d67eb0dcccc725e33fa43e08de045816f：0x48a0501d67eb0dcccc725e33fa43e08de045816f
- Gutter Punks - Everai (GPE) https://etherscan.io/address/0xa84145eb5f78d082a295e394ba5d906e9b8ff1e3：0xa84145eb5f78d082a295e394ba5d906e9b8ff1e3
- SimpleGraveStone https://etherscan.io/address/0xdf4464a2379ee1bd48a95cfebeb4e07c52f79da2：0xdf4464a2379ee1bd48a95cfebeb4e07c52f79da2


### Demo video
https://drive.google.com/file/d/1KJzF5SZM96RwrCNrQsHJ0XSdAogyLKcQ/view?usp=sharing


### More information - ETH Global Gallery
https://ethglobal.com/showcase/zkdvc-on-polygonid-x2ie1

