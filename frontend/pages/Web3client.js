import Web3 from "web3";
import pub from "../../build/contracts/pub.json";

let selectedAccount;

export const init = async () => {
        // console.log(web3);
        let provider = window.ethereum;
        if(typeof provider !== "undefined"){
            provider.request({method: "eth_requestAccounts"})
            .then((accounts) => {
                // console.log(accounts);
                selectedAccount = accounts[0];
                console.log(selectedAccount);
            })
            .catch((err) => {
                console.log(err);
            });

            provider.on('accountsChanged', (accounts) => {
                selectedAccount = accounts[0];
                console.log(accounts);
            });
        }
        const web3 = new Web3(provider);
        const networkId = await web3.eth.net.getId();
        const nftContract = new web3.eth.Contract(pub.abi, nftContract.networks[networkId].address);
        
}