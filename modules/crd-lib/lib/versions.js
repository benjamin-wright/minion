const client = require('./client');

module.exports = {
    get,
    list,
    delete: deleteResource,
    tryDelete
};

async function list(namespace, selectors) {
    let query;

    if (selectors) {
        query = {
            qs: {
                labelSelectors: selectors
            }
        };
    }

    const cli = await client.get();
    const result = await cli
        .apis['minion.ponglehub.co.uk']
        .v1alpha1
        .namespace(namespace)
        .version
        .get(query);

    return result.items;
}

async function get(namespace, name) {
    const cli = await client.get();
    const result = await cli
        .apis['minion.ponglehub.co.uk']
        .v1alpha1
        .namespace(namespace)
        .version(name)
        .get();

    return result.body;
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
        .version(name)
        .delete();

    return result;
}
