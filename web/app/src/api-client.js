import { ApiClient } from "@/manager-api";
import { CountingApiClient } from "@/stores/api-query-count";
import { api as apiURL } from '@/urls'

/**
 * Scrub the custom User-Agent header from the API client, for those webbrowsers
 * who do not want to have it set.
 *
 * It's actually scrubbed for all webbrowsers, as those privacy-first
 * webbrowsers also make it hard to fingerprint which browser you're using (for
 * good reason).
 *
 * @param {ApiClient} apiClient
 */
export function scrubAPIClient(apiClient) {
  delete apiClient.defaultHeaders['User-Agent'];
}

/**
 * @returns {ApiClient} Bare API client that is not connected to the UI in any way.
 */
export function newBareAPIClient() {
  const apiClient = new ApiClient(apiURL());
  scrubAPIClient(apiClient);
  return apiClient;
}

let apiClient = null;

/**
 * @returns {ApiClient} API client that updates the UI to show long-running queries.
 */
export function getAPIClient() {
  if (apiClient == null) {
    apiClient = new CountingApiClient(apiURL());
    scrubAPIClient(apiClient);
  }
  return apiClient;
}
