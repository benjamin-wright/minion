const crds = require('@minion-ci/crd-lib');
const NAMESPACE = process.env.NAMESPACE;

module.exports = {
    getMonitorJobs: getMonitorJobs,
    deleteMonitorJobs: deleteMonitorJobs
};

async function getMonitorJobs() {
    const client = await crds.client.get();

    const jobs = await client.apis.batch.v1.namespace(NAMESPACE).jobs.get();

    console.log(jobs.body.items.map(j => j.metadata.name));

    return jobs.body.items.filter(job => job.metadata.name.includes('-monitor-'));
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
