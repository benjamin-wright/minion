const crds = require('@minion-ci/crd-lib');
const NAMESPACE = process.env.NAMESPACE;

describe('test', () => {
    beforeEach(async () => {
        await crds.resources.tryDelete(NAMESPACE, 'my-resource');
        await crds.resources.tryDelete(NAMESPACE, 'my-resource-2');
    });

    it('works', async () => {
        await crds.resources.post(NAMESPACE, 'my-resource', 'my-image', [{ name: 'FOO', value: 'BAR' }], []);
        await crds.resources.post(NAMESPACE, 'my-resource-2', 'my-image', [{ name: 'FOO', value: 'BAR' }], []);
    });
});
