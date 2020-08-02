const jobs = require('../helpers/test-job');
const faker = require('faker');
const async = require('@minion-ci/async');
const crds = require('@minion-ci/crd-lib');

const NAMESPACE = process.env.NAMESPACE;

describe('should work', () => {
    beforeAll(async () => {
        await jobs.clean();
    });

    it('should complete successfully', async () => {
        const jobName = faker.random.uuid();
        const version = faker.random.number();
        const resource = faker.name.firstName().toLowerCase();
        await jobs.createJob(jobName, version, resource);

        await async.waitFor(async () => {
            const status = await jobs.getStatus(jobName);

            if (!status) {
                throw new Error('Job status missing');
            }

            if (!status.succeeded && !status.error) {
                throw new Error('Job not finished');
            }
        });
    });

    it('should create a new version CRD', async () => {
        const version = faker.random.number();
        const resource = faker.name.firstName().toLowerCase();
        await jobs.createJob(faker.random.uuid(), version, resource);

        const versionObj = await async.waitFor(async () => {
            const versions = await crds.versions.list(NAMESPACE, `resource=${resource}`);
            if (versions.length === 0) {
                throw new Error(`Didn't find version for resource ${resource}`);
            }

            if (versions.length > 1) {
                throw new Error(`Expected one version for resource ${resource}, found ${versions.length}`);
            }

            return versions[0];
        });

        expect(versionObj.spec).toEqual({
            resource,
            version: version.toString()
        });
    });
});
