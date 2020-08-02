const jobs = require('../helpers/test-job');
const faker = require('faker');
const async = require('@minion-ci/async');
const crds = require('@minion-ci/crd-lib');

const NAMESPACE = process.env.NAMESPACE;

describe('should work', () => {
    beforeAll(async () => {
        await jobs.clean();
    });

    it('should pass', async () => {
        const version = faker.random.number();
        const pipeline = faker.internet.avatar();
        const resource = faker.internet.domainName();
        await jobs.createJob(faker.random.uuid(), version, pipeline, resource);

        const versionObj = await async.waitFor(async () => {
            return await crds.versions.list(NAMESPACE, `pipeline=${pipeline},resource=${resource},version=${version}`);
        });

        expect(versionObj).toEqual({});
    });
});
