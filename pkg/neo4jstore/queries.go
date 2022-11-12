package neo4jstore

const FindAddressByHashQuery = `MATCH (a:Address {address: $address}) RETURN a`

const FindTransactionsInBlockQuery = `MATCH (b:Block {height: $height}) WITH b MATCH (t:Transaction)-[:BELONGS_TO]->(b) RETURN t SKIP $skip LIMIT $limit`
const FindTransactionsInBlockCountQuery = `MATCH (b:Block {height: $height}) WITH b MATCH (t:Transaction)-[:BELONGS_TO]->(b) RETURN count(t) AS t`

const FindBlockByTransactionQuery = `MATCH (t:Transaction {txid: $txid}) WITH t MATCH (t)-[:BELONGS_TO]->(b:Block) RETURN b`
const FindBlockByHeightQuery = `MATCH (b:Block {height: $height}) RETURN b`
const FindBlockByHashQuery = `MATCH (b:Block {hash: $hash}) RETURN b`
const FindTransactionByHashQuery = `MATCH (t:Transaction {txid: $txid}) RETURN t`

const FindContactByAddressQuery = `MATCH (a:Address {address: $address}) WITH a MATCH (a)-[:FIGURED_IN]->(m:Mention)-[:INFO]->(c:Contact) RETURN c`
const FindMentionsByAddressQuery = `MATCH (a:Address {address: $address}) WITH a MATCH (a)-[:FIGURED_IN]->(m:Mention) RETURN m SKIP $skip LIMIT $limit`
const FindMentionsByAddressCountQuery = `MATCH (a:Address {address: $address}) WITH a MATCH (a)-[:FIGURED_IN]->(m:Mention) RETURN count(m) AS t`

const FindTransactionsByAddressQuery = `MATCH (a:Address {address: $address}) WITH a MATCH (a)--(t:Transaction) RETURN DISTINCT t SKIP $skip LIMIT $limit`
const FindTransactionsByAddressCountQuery = `MATCH (a:Address {address: $address}) WITH a MATCH (a)--(t:Transaction) RETURN count(DISTINCT t) AS t`

const FindOutcomingTransactionsQuery = `MATCH (t:Transaction {txid: $txid}) WITH t MATCH (a:Address)<-[:RECEIVES]-(t) RETURN a SKIP $skip LIMIT $limit`
const FindOutcomingTransactionsCountQuery = `MATCH (t:Transaction {txid: $txid}) WITH t MATCH (a:Address)<-[:RECEIVES]-(t) RETURN count(a) AS t`

const FindIncomingTransactionsQuery = `MATCH (t:Transaction {txid: $txid}) WITH t MATCH (a:Address)-[:SENDS]->(t) RETURN a SKIP $skip LIMIT $limit`
const FindIncomingTransactionsCountQuery = `MATCH (t:Transaction {txid: $txid}) WITH t MATCH (a:Address)-[:SENDS]->(t) RETURN count(a) AS t`

const FindWalletAddressesQuery = `MATCH (a:Address)-[:CS]->(w:Wallet {wid: $wid}) RETURN a SKIP $skip LIMIT $limit`

const FindWalletForAddressQuery = `MATCH (a:Address {address: $address}) WITH a MATCH (a)--(w:Wallet) RETURN w`
