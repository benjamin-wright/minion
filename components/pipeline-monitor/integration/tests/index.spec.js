const crds = require('@minion-ci/crd-lib');
const NAMESPACE = process.env.NAMESPACE;

describe('test', () => {
    beforeEach(async () => {
        await crds.pipelines.delete(NAMESPACE, 'my-pipeline');
    });

    it('works', async () => {
        await crds.pipelines.post(NAMESPACE, 'my-pipeline', [], []);
        const result = await crds.pipelines.get(NAMESPACE, 'my-pipeline');
        expect(result.metadata.name).toEqual('my-pipeline');
    });
});
