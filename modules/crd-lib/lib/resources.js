const client = require('./client');

module.exports = {
    get,
    list,
    post,
    put,
    delete: deleteResource,
    tryDelete
};

async function list(namespace) {
    const cli = await client.get();
    const result = await cli
        .apis['minion.ponglehub.co.uk']
        .v1alpha1
        .namespace(namespace)
        .resource
        .get();

    return result.body.items;
}

async function get(namespace, name) {
    const cli = await client.get();
    const result = await cli
        .apis['minion.ponglehub.co.uk']
        .v1alpha1
        .namespace(namespace)
        .resource(name)
        .get();

    return result.body;
}

async function put(namespace, name, image, env, secrets) {
    const cli = await client.get();
    const existing = await get(namespace, name);
    const manifest = {
        apiVersion: 'minion.ponglehub.co.uk/v1alpha1',
        kind: 'Resource',
        metadata: {
            name,
            namespace,
            resourceVersion: existing.metadata.resourceVersion
        },
        spec: {
            image,
            env,
            secrets
        }
    };

    const result = await cli
        .apis['minion.ponglehub.co.uk']
        .v1alpha1
        .namespace(namespace)
        .resource(name)
        .put({
            body: manifest
        });

    if (result.statusCode !== 200) {
        throw new Error(`Failed updating resource: expected 200, got ${result.statusCode}`);
    }
}

async function post(namespace, name, image, env, secrets) {
    const manifest = {
        apiVersion: 'minion.ponglehub.co.uk/v1alpha1',
        kind: 'Resource',
        metadata: {
            name,
            namespace
        },
        spec: {
            image,
            env,
            secrets
        }
    };

    const cli = await client.get();
    const result = await cli
        .apis['minion.ponglehub.co.uk']
        .v1alpha1
        .namespace(namespace)
        .resource
        .post({
            body: manifest
        });

    return result;
}

async function tryDelete(namespace, name) {
    try {
        return await deleteResource(namespace, name);
    } catch (err) {
        if (err.statusCode === 404) {
            return {
                statusCode: 404,
                body: null
            };
        }

        throw err;
    }
}

async function deleteResource(namespace, name) {
    const cli = await client.get();
    const result = await cli
        .apis['minion.ponglehub.co.uk']
        .v1alpha1
        .namespace(namespace)
        .resource(name)
        .delete();

    return result;
}
