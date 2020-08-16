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

    it('should create a resource monitor cronjob', async () => {
        const resourceName = faker.name.firstName().toLowerCase();

        await crds.resources.post(NAMESPACE, resourceName, 'docker.io/busybox', [{ name: 'FOO', value: 'BAR' }], []);

        const cronjob = await async.waitFor(() => k8s.getMonitorJob(`${resourceName}-monitor`));
        const podSpec = cronjob.spec.jobTemplate.spec.template.spec;

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

    it('should update cronjob when resource is updated', async () => {
        const resourceName = faker.name.firstName().toLowerCase();

        await crds.resources.post(NAMESPACE, resourceName, 'docker.io/busybox:1.31', [{ name: 'FOO', value: 'BAR' }], []);
        await crds.resources.put(NAMESPACE, resourceName, 'docker.io/busybox:1.32', [{ name: 'FOO', value: 'BAZ' }], []);

        await async.waitFor(async () => {
            const cronjob = await async.waitFor(() => k8s.getMonitorJob(`${resourceName}-monitor`));
            const podSpec = cronjob.spec.jobTemplate.spec.template.spec;

            const actual = podSpec.initContainers.map(c => ({ image: c.image, env: c.env }));
            const expected = [
                {
                    image: 'docker.io/busybox:1.32',
                    env: [
                        { name: 'FOO', value: 'BAZ' }
                    ]
                }
            ];

            expect(actual).toEqual(expected);
        });
    });
});
