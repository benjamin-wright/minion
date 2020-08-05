const crds = require('@minion-ci/crd-lib');
const NAMESPACE = process.env.NAMESPACE;
const faker = require('faker');

describe('test', () => {
    beforeEach(async () => {
        const resources = await crds.resources.list(NAMESPACE);
        await Promise.all(resources.map(res => crds.resources.delete(NAMESPACE, res.metadata.name)));
    });

    it('works', async () => {
        const job1 = faker.name.firstName().toLowerCase();
        const job2 = faker.name.firstName().toLowerCase();

        await crds.resources.post(NAMESPACE, job1, 'my-image', [{ name: 'FOO', value: 'BAR' }], []);
        await crds.resources.post(NAMESPACE, job2, 'my-image', [{ name: 'FOO', value: 'BAR' }], []);
    });
});
