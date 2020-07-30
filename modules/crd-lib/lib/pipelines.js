const client = require('./client');

module.exports = {
    get,
    list,
    post,
    delete: deletePipeline
};

async function list(namespace) {
    const cli = await client.get();
    const result = await cli
        .apis['minion.ponglehub.co.uk']
        .v1alpha1
        .namespace(namespace)
        .pipeline
        .get();

    return result.items;
}

async function get(namespace, name) {
    const cli = await client.get();
    const result = await cli
        .apis['minion.ponglehub.co.uk']
        .v1alpha1
        .namespace(namespace)
        .pipeline(name)
        .get();

    return result.body;
}

async function post(namespace, name, resources, steps) {
    const manifest = {
        apiVersion: 'minion.ponglehub.co.uk/v1alpha1',
        kind: 'Pipeline',
        metadata: {
            name,
            namespace
        },
        spec: {
            resources,
            steps
        }
    };

    const cli = await client.get();
    const result = await cli
        .apis['minion.ponglehub.co.uk']
        .v1alpha1
        .namespace(namespace)
        .pipeline
        .post({
            body: manifest
        });

    return result;
}

async function deletePipeline(namespace, name) {
    const cli = await client.get();
    const result = await cli
        .apis['minion.ponglehub.co.uk']
        .v1alpha1
        .namespace(namespace)
        .pipeline(name)
        .delete();

    return result;
}
