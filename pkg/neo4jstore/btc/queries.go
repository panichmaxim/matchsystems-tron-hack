package btc

const FindWalletByWidQuery = `MATCH (w:Wallet {wid: $wid}) RETURN w`
const FindTransactionsInBlockByHashQuery = `MATCH (b:Block {hash: $hash}) WITH b MATCH (t:Transaction)-[:BELONGS_TO]->(b) RETURN t SKIP $skip LIMIT $limit`
const FindTransactionsInBlockByHashCountQuery = `MATCH (b:Block {hash: $hash}) WITH b MATCH (t:Transaction)-[:BELONGS_TO]->(b) RETURN count(t) AS t`
