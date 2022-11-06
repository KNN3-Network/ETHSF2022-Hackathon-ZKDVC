# ETHSF2022-Hackathon-zkDVC
ETH SF 2022 Hackathon project - zkDVC on Polygon ID 

zkDVC is a Threshold-based Credit Issuance Template Module on Polygon ID. The primary goal of this initiative is to allow on-chain creators to leverage users' behavioral data to evaluate the record and issue credits score accordingly. It enables a new way of credit issuance and cross-platform action/status verification.

zkDVC focuses on 3 technical goals in this platform:
1.  Template-based claim schema extension allows multiple threshold-type claims to be packed into one offer.
2.  Support continuous tracking for transparent revocation.
3.  Multi-addresses data aggregating and credit issuing with anonymity (To be continued)


#### Project Brief
*   Issuer server, which is used to manage claims, in which go language is used to build the project framework, iden3 is used to manage the issuer’s identity and credentials, and zk is used to construct proofs and verification. Claims are organized according to the iden3 protocol, and each user includes three Merkle trees, which are stored in postgresql.
*   The back-end part of wallet (wallet-end) is designed with Go + iden3. Go is used as the project framework to provide an api server for wallet-front; the iden3 protocol is used to manage the holder’s claims and identity. The specific data is stored in the iden3 protocol according to the in postgresql. Wallet-end will also call a Badge contract (ERC-721), provide proof of its own claims and obtain Badge.
*   The front-end of wallet (wallet-front) is implemented by react, and its functions are obtained from wallet-end.
*   By using the Dynamic Verifiable Credentials (DVC) powered by KNN3, we can verify user accounts based on real-time data on the chain. DVC makes this project possible, and in the future, we can even aggregate and authenticate different assets of different accounts of the same user through DVC.


### Introducory deck
https://drive.google.com/file/d/116G7nltJTGa8T1Dg7i92Fdgy02W8M2Ew/view?usp=share_link


