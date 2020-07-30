const { Client, KubeConfig } = require('kubernetes-client');
const Request = require('kubernetes-client/backends/request');

let client;

async function get() {
    if (!client) {
        const kubeconfig = new KubeConfig();
        kubeconfig.loadFromCluster();
        const backend = new Request({ kubeconfig });
        client = await new Client({ backend }).loadSpec();
    }

    return client;
}

module.exports = {
    get: get
};
