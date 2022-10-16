<template>
  <div>
    <span @click="togglePopover">{{ priority }}</span>
    <div v-show="showPopover">
      <input type="number" v-model="priorityState">
      <button @click="updateJobPriority">Update</button>
      <button @click="togglePopover">Cancel</button>
      <span v-if="errorMessage">{{ errorMessage }}</span>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

import { useNotifs } from '@/stores/notifications';
import { apiClient } from '@/stores/api-query-count';
import { JobsApi, JobPriorityChange } from '@/manager-api';
import postcss from 'postcss';

const props = defineProps({
  jobId: String,
  priority: Number
});

// Init notification state
const notifs = useNotifs();

// Init internal state
const priorityState = ref(props.priority);
const showPopover = ref(false);
const errorMessage = ref('');

// Methods
function updateJobPriority() {
  const jobPriorityChange = new JobPriorityChange(priorityState.value);
  const jobsAPI = new JobsApi(apiClient);
  return jobsAPI.setJobPriority(props.jobId, jobPriorityChange)
    .then(() => {
      notifs.add(`Updated job priority to ${priorityState.value}`)
      showPopover.value = false;
      errorMessage.value = '';
    })
    .catch((err) => {
      errorMessage.value = err.body.message;
    });
}

function togglePopover() {
  // 'reset' the priorityState to match the actual priority
  priorityState.value = props.priority;
  showPopover.value = !showPopover.value;
}
</script>
