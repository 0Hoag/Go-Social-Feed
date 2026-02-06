export interface TokenMetadata {
    symbol: string;
    description: string;
    website?: string;
    explorer?: string;
    totalSupply?: string;
    circulatingSupply?: string;
    rank?: string;
}

export const TOKEN_METADATA: Record<string, TokenMetadata> = {
    "BTC": {
        symbol: "BTC",
        description: "Bitcoin is the world's first cryptocurrency, a digital asset that uses public-key cryptography to record, sign and send transactions over the Bitcoin blockchain - all done without the oversight of a central authority.",
        website: "https://bitcoin.org",
        explorer: "https://mempool.space",
        rank: "#1",
        totalSupply: "21M",
        circulatingSupply: "19.6M"
    },
    "ETH": {
        symbol: "ETH",
        description: "Ethereum is a decentralized open-source blockchain system that features its own cryptocurrency, Ether. ETH works as a platform for numerous other cryptocurrencies, as well as for the execution of decentralized smart contracts.",
        website: "https://ethereum.org",
        explorer: "https://etherscan.io",
        rank: "#2",
        totalSupply: "120M",
        circulatingSupply: "120M"
    },
    "SOL": {
        symbol: "SOL",
        description: "Solana is a highly functional open source project that banks on blockchain technology's permissionless nature to provide decentralized finance (DeFi) solutions.",
        website: "https://solana.com",
        explorer: "https://solscan.io",
        rank: "#4",
        totalSupply: "570M",
        circulatingSupply: "440M"
    },
    "BNB": {
        symbol: "BNB",
        description: "BNB is the cryptocurrency coin that powers the BNB Chain ecosystem. As one of the world's most popular utility tokens, not only can you buy or sell BNB like any other cryptocurrency, but BNB comes with a wide range of applications and benefits.",
        website: "https://www.bnbchain.org",
        explorer: "https://bscscan.com",
        rank: "#5",
        totalSupply: "153M",
        circulatingSupply: "153M"
    },
    "XRP": {
        symbol: "XRP",
        description: "XRP is the native cryptocurrency of the XRP Ledger, which is an open-source, permissionless and decentralized blockchain technology. The XRP Ledger utilizes the Federated Consensus mechanism to settle transactions.",
        website: "https://xrpl.org",
        explorer: "https://xrpscan.com",
        rank: "#6",
        totalSupply: "100B",
        circulatingSupply: "54.8B"
    },
    "DOGE": {
        symbol: "DOGE",
        description: "Dogecoin is an open-source peer-to-peer digital currency, favored by Shiba Inus worldwide. At its heart, Dogecoin is the accidental crypto movement that makes people smile!",
        website: "https://dogecoin.com",
        explorer: "https://dogechain.info",
        rank: "#8",
        totalSupply: "143B",
        circulatingSupply: "143B"
    },
    "PEPE": {
        symbol: "PEPE",
        description: "Pepe is a deflationary meme coin launched on Ethereum. The crypto was created as a tribute to the Pepe the Frog internet meme, created by Matt Furie, which gained popularity in the early 2000s.",
        website: "https://www.pepe.vip",
        explorer: "https://etherscan.io/token/0x6982508145454Ce325dDbE47a25d4ec3d2311933",
        rank: "#24",
        totalSupply: "420.69T",
        circulatingSupply: "420.69T"
    },
    "WIF": {
        symbol: "WIF",
        description: "dogwifhat (WIF) is a meme coin on the Solana blockchain. It features a Shiba Inu dog wearing a pink knitted hat. The token was launched in late 2023 and quickly gained popularity within the crypto community.",
        website: "https://dogwifcoin.org",
        explorer: "https://solscan.io/token/EKpQGSJtjMFqKZ9KQanSqYXRcF8fBopzLHYxdM65zcjm",
        rank: "#40",
        totalSupply: "998.9M",
        circulatingSupply: "998.9M"
    }
    // Add generic fallback
};

export const getMetadata = (symbol: string): TokenMetadata => {
    // Basic mapping for common pairs
    const baseSymbol = symbol.replace("USDT", "").toUpperCase();
    return TOKEN_METADATA[baseSymbol] || {
        symbol: baseSymbol,
        description: `${baseSymbol} is a cryptocurrency token. Detailed information is currently being updated.`,
        rank: "N/A",
        totalSupply: "N/A",
        circulatingSupply: "N/A"
    };
};
