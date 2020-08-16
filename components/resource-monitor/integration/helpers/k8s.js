const crds = require('@minion-ci/crd-lib');
const NAMESPACE = process.env.NAMESPACE;

module.exports = {
    getMonitorJob: getMonitorJob,
    deleteMonitorJobs: deleteMonitorJobs
};

async function getMonitorJob(name) {
    const client = await crds.client.get();
    const result = await client.apis.batch.v1beta1.namespace(NAMESPACE).cronjob(name).get();

    if (result.statusCode !== 200) {
        throw new Error(`Expected 200 status code, recieved ${result.statusCode}`);
    }

    return result.body;
}

async function deleteMonitorJobs() {
    const client = await crds.client.get();

    const jobs = await client.apis.batch.v1.namespace(NAMESPACE).jobs.get();
    const names = jobs.body.items.filter(job => job.metadata.name.includes('-monitor-')).filter(job => !job.metadata.name.includes('-monitor-int-tests')).map(job => job.metadata.name);

    await Promise.all(names.map(name => deleteMonitorJob(client, name)));
}

async function deleteMonitorJob(client, name) {
    await client.apis.batch.v1.namespace(NAMESPACE).jobs(name).delete();
    await client.api.v1.namespace(NAMESPACE).pods.delete({ qs: { labelSelector: `job-name=${name}` } });
}
