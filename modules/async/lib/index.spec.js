const async = require('./index');

describe('sleep', () => {
    beforeAll(() => {
        jest.useFakeTimers();
    });

    it('should not return before timeout elapses', async () => {
        const promise = async.sleep(500);
        promise.then(() => { promise.done = true; });

        await Promise.resolve();

        return expect(promise.done).not.toBe(true);
    });

    it('should return after the timeout elapses', async () => {
        const promise = async.sleep(500);
        promise.then(() => { promise.done = true; });

        jest.runAllTimers();
        await Promise.resolve();

        return expect(promise.done).toBe(true);
    });
});
