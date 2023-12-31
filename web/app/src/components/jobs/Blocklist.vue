<template>
  <div v-if="isFetching" class="dl-no-data">
    <span>Fetching blocklist...</span>
  </div>
  <template v-else>
    <table class="blocklist" v-if="blocklist.length">
      <tr>
        <th>Worker</th>
        <th>Task Type</th>
        <th></th>
      </tr>
      <tr v-for="entry in blocklist">
        <td>
          <link-worker :worker="{ id: entry.worker_id, name: entry.worker_name }" />
        </td>
        <td>{{ entry.task_type }}</td>
        <td>
          <button
            class="btn in-table-row"
            @click="removeBlocklistEntry(entry)"
            title="Allow this worker to execute these task types">
            ❌
          </button>
        </td>
      </tr>
    </table>
    <div v-else class="dl-no-data">
      <span>This job has no blocked workers.</span>
    </div>
  </template>
  <p v-if="errorMsg" class="error">Error fetching blocklist: {{ errorMsg }}</p>
</template>

<script setup>
import { getAPIClient } from '@/api-client';
import { JobsApi } from '@/manager-api';
import LinkWorker from '@/components/LinkWorker.vue';
import { watch, onMounted, inject, ref, nextTick } from 'vue';

// jobID should be the job UUID string.
const props = defineProps(['jobID']);
const emit = defineEmits(['reshuffled']);

const jobsApi = new JobsApi(getAPIClient());
const isVisible = inject('isVisible');
const isFetching = ref(false);
const errorMsg = ref('');
const blocklist = ref([]);

function refreshBlocklist() {
  if (!isVisible.value) {
    return;
  }

  isFetching.value = true;
  jobsApi
    .fetchJobBlocklist(props.jobID)
    .then((newBlocklist) => {
      blocklist.value = newBlocklist;
    })
    .catch((error) => {
      errorMsg.value = error.message;
    })
    .finally(() => {
      isFetching.value = false;
    });
}

function removeBlocklistEntry(blocklistEntry) {
  jobsApi
    .removeJobBlocklist(props.jobID, { jobBlocklistEntry: [blocklistEntry] })
    .then(() => {
      blocklist.value = blocklist.value.filter(
        (entry) =>
          !(
            entry.worker_id == blocklistEntry.worker_id &&
            entry.task_type == blocklistEntry.task_type
          )
      );
    })
    .catch((error) => {
      console.log('Error removing entry from blocklist', error);
      refreshBlocklist();
    });
}

watch(() => props.jobID, refreshBlocklist);
watch(blocklist, () => {
  const emitter = () => {
    emit('reshuffled');
  };
  nextTick(() => {
    nextTick(emitter);
  });
});
watch(isVisible, refreshBlocklist);
onMounted(refreshBlocklist);
</script>

<style scoped>
table.blocklist {
  width: 100%;
  font-family: var(--font-family-mono);
  font-size: var(--font-size-sm);
  border-collapse: collapse;
}

table.blocklist td,
table.blocklist th {
  text-align: left;
  padding: calc(var(--spacer-sm) / 2) var(--spacer-sm);
}

table.blocklist th {
  color: var(--color-text-muted);
  font-weight: normal;
}

table.blocklist tr {
  background-color: var(--table-color-background-row);
}

table.blocklist tr:nth-child(odd) {
  background-color: var(--table-color-background-row-odd);
}

button.in-table-row {
  background-color: unset;
  border: unset;
  padding: 0;
}
</style>
