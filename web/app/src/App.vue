<template>
  <header>
    <router-link :to="{ name: 'index' }" class="navbar-brand">{{ flamencoName }}</router-link>
    <nav>
      <ul>
        <li>
          <router-link :to="{ name: 'jobs' }">Jobs</router-link>
        </li>
        <li>
          <router-link :to="{ name: 'workers' }">Workers</router-link>
        </li>
        <li>
          <router-link :to="{ name: 'tags' }">Tags</router-link>
        </li>
        <li>
          <router-link :to="{ name: 'last-rendered' }">Last Rendered</router-link>
        </li>
      </ul>
    </nav>
    <api-spinner />
    <span class="app-version">
      <a :href="backendURL('/flamenco-addon.zip')">add-on</a>
      | <a :href="backendURL('/api/v3/swagger-ui/')">API</a> | version: {{ flamencoVersion }}
    </span>
  </header>
  <router-view></router-view>
</template>

<script>
import * as API from '@/manager-api';
import { getAPIClient } from '@/api-client';
import { backendURL } from '@/urls';
import { useSocketStatus } from '@/stores/socket-status';

import ApiSpinner from '@/components/ApiSpinner.vue';

const DEFAULT_FLAMENCO_NAME = 'Flamenco';
const DEFAULT_FLAMENCO_VERSION = 'unknown';

export default {
  name: 'App',
  components: {
    ApiSpinner,
  },
  data: () => ({
    flamencoName: DEFAULT_FLAMENCO_NAME,
    flamencoVersion: DEFAULT_FLAMENCO_VERSION,
    backendURL: backendURL,
  }),
  mounted() {
    window.app = this;
    this.fetchManagerInfo();

    const sockStatus = useSocketStatus();
    this.$watch(
      () => sockStatus.isConnected,
      (isConnected) => {
        if (!isConnected) return;
        if (!sockStatus.wasEverDisconnected) return;
        this.socketIOReconnect();
      }
    );
  },
  methods: {
    fetchManagerInfo() {
      const metaAPI = new API.MetaApi(getAPIClient());
      metaAPI.getVersion().then((version) => {
        this.flamencoName = version.name;
        this.flamencoVersion = version.version;
        document.title = version.name;
      });
    },

    socketIOReconnect() {
      const metaAPI = new API.MetaApi(getAPIClient());
      metaAPI.getVersion().then((version) => {
        if (version.name === this.flamencoName && version.version == this.flamencoVersion) return;
        console.log(`Updated from ${this.flamencoVersion} to ${version.version}`);
        location.reload();
      });
    },
  },
};
</script>

<style>
@import 'assets/base.css';
@import 'assets/tabulator.css';
</style>
