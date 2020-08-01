const jobs = require('../helpers/test-job');
const faker = require('faker');
const async = require('@minion-ci/async');

describe('should work', () => {
    beforeAll(async () => {
        await jobs.clean();
    });

    it('should pass', async () => {
        await jobs.createJob(faker.random.uuid(), '1');
        await async.sleep(1000);
    });
});
