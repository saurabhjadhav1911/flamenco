<template>
  <div class="popover-container">
    <span @click="togglePopover" class="popover-toggle" title="Set priority for this job">
      {{ priority }}
    </span>
    <div v-show="showPopover" class="popover">
      <div class="popover-header">
        <span>Job Priority</span>
        <button @click="togglePopover">&#10006;</button>
      </div>
      <div class="popover-form">
        <input type="number" v-model="priorityState">
        <button @click="updateJobPriority" class="btn-primary">Set</button>
      </div>
      <div class="input-help-text">
        Range 1-100.
      </div>
      <div class="popover-error" v-if="errorMessage">
        <span>{{ errorMessage }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

import { useNotifs } from '@/stores/notifications';
import { getAPIClient } from "@/api-client";
import { JobsApi, JobPriorityChange } from '@/manager-api';

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
  const jobsAPI = new JobsApi(getAPIClient());
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

<style scoped>
.popover-toggle {
  cursor: pointer;
  display: block;
}

.popover-toggle:hover {
  color: var(--color-accent);
}

.popover-container {
  position: relative;
}

.popover {
  background-color: var(--color-background-popover);
  border-radius: var(--border-radius);
  bottom: 25px;
  box-shadow: var(--box-shadow-float);
  display: flex;
  flex-direction: column;
  left: -5rem;
  max-width: 20rem;
  min-width: 10rem;
  padding: var(--spacer-sm) var(--spacer);
  position: absolute;
  z-index: 1;
}

.popover-header {
  align-items: baseline;
  display: flex;
  font-weight: bold;
  justify-content: space-between;
}

/* Close/cancel popover button. */
.popover-header button {
  background-color: transparent;
  border: none;
  margin-left: var(--spacer-lg);
  margin-right: -5px;
}

.popover-form {
  display: flex;
  margin: var(--spacer-sm) 0;
}

/* Number input field. */
.popover-form input {
  margin-right: var(--spacer-sm);
  width: 10ch;
  padding: 0 var(--spacer-sm);
}

/* Save/Set button. */
.popover-form button {
  flex: 1
}

.input-help-text {
  margin-left: 0;
}

.popover-error {
  color: var(--color-danger);
}
</style>
