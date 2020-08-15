const crds = require('@minion-ci/crd-lib');
const async = require('@minion-ci/async');
const k8s = require('../helpers/k8s');
const NAMESPACE = process.env.NAMESPACE;
const faker = require('faker');

describe('test', () => {
    beforeEach(async () => {
        const resources = await crds.resources.list(NAMESPACE);
        await Promise.all(resources.map(res => crds.resources.delete(NAMESPACE, res.metadata.name)));
        await k8s.deleteMonitorJobs();
    });

    it('works', async () => {
        const job1 = faker.name.firstName().toLowerCase();
        const job2 = faker.name.firstName().toLowerCase();

        await crds.resources.post(NAMESPACE, job1, 'docker.io/busybox', [{ name: 'FOO', value: 'BAR' }], []);
        await crds.resources.post(NAMESPACE, job2, 'docker.io/busybox', [{ name: 'FOO', value: 'BAR' }], []);

        const client = await crds.client.get();

        const result = await async.waitFor(() => client.apis.batch.v1beta1.namespace(NAMESPACE).cronjobs(`${job1}-monitor`).get());
        const podSpec = result.body.spec.jobTemplate.spec.template.spec;

        const actual = podSpec.initContainers.map(c => ({ image: c.image, env: c.env }));
        const expected = [
            {
                image: 'docker.io/busybox',
                env: [
                    { name: 'FOO', value: 'BAR' }
                ]
            }
        ];

        expect(actual).toEqual(expected);
    });
});
