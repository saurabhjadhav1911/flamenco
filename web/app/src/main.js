import { createApp } from 'vue';
import { createPinia } from 'pinia';
import { DateTime } from 'luxon';

import App from '@/App.vue';
import SetupAssistant from '@/SetupAssistant.vue';
import autoreload from '@/autoreloader';
import router from '@/router/index';
import setupAssistantRouter from '@/router/setup-assistant';
import { MetaApi } from '@/manager-api';
import { newBareAPIClient } from '@/api-client';
import * as urls from '@/urls';

// Ensure Tabulator can find `luxon`, which it needs for sorting by
// date/time/datetime.
window.DateTime = DateTime;
// plain removes any Vue reactivity.
window.plain = (x) => JSON.parse(JSON.stringify(x));
// objectEmpty returns whether the object is empty or not.
window.objectEmpty = (o) => !o || Object.entries(o).length == 0;

// Automatically reload the window after a period of inactivity from the user.
autoreload();

const pinia = createPinia();

function normalMode() {
  const app = createApp(App);
  app.use(pinia);
  app.use(router);
  app.mount('#app');
}

function setupAssistantMode() {
  console.log('Flamenco Setup Assistant is starting');
  const app = createApp(SetupAssistant);
  app.use(pinia);
  app.use(setupAssistantRouter);
  app.mount('#app');
}

/* This cannot use the client from '@/stores/api-query-count', as that would
 * require Pinia, which is unavailable until the app is actually started. And to
 * know which app to start, this API call needs to return data. */
const apiClient = newBareAPIClient();
const metaAPI = new MetaApi(apiClient);
metaAPI
  .getConfiguration()
  .then((config) => {
    if (config.isFirstRun) setupAssistantMode();
    else normalMode();
  })
  .catch((error) => {
    console.warn('Error getting Manager configuration:', error);
  });
