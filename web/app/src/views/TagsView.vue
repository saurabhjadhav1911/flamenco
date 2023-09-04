<template>
  <div class="col col-workers-list">
    <h2 class="column-title">Tag Details</h2>

    <div class="action-buttons btn-bar-group">
      <div class="btn-bar">
        <button @click="fetchTags">Refresh</button>
        <button @click="deleteTag" :disabled="!selectedTag">Delete Tag</button>
      </div>
    </div>

    <div class="action-buttons btn-bar">
      <form @submit="createTag">
        <div class="create-tag-container">
          <input
            type="text"
            name="newtagname"
            v-model="newTagName"
            placeholder="New Tag Name"
            class="create-tag-input"
          />
          <button
            id="submit-button"
            type="submit"
            :disabled="newTagName.trim() === ''"
          >
            Create Tag
          </button>
        </div>
      </form>
    </div>

    <div id="tag-table-container"></div>
  </div>
  <footer class="app-footer"></footer>
</template>

<style scoped>
.create-tag-container {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.create-tag-input {
  flex: 1;
  margin-right: 10px;
  height: 30px;
}
</style>

<script>
import { TabulatorFull as Tabulator } from "tabulator-tables";
import { useWorkers } from "@/stores/workers";
import { useNotifs } from "@/stores/notifications";
import { WorkerMgtApi } from "@/manager-api";
import { WorkerTag } from "@/manager-api";
import { getAPIClient } from "@/api-client";
import TabItem from "@/components/TabItem.vue";
import TabsWrapper from "@/components/TabsWrapper.vue";

export default {
  components: {
    TabItem,
    TabsWrapper,
  },
  data() {
    return {
      tags: [],
      selectedTag: null,
      newTagName: "",
      workers: useWorkers(),
      activeRowIndex: -1,
    };
  },

  mounted() {
    this.fetchTags();

    const tag_options = {
      columns: [
        { title: "Name", field: "name", sorter: "string", editor: "input" },
        {
          title: "Description",
          field: "description",
          sorter: "string",
          editor: "input",
        },
      ],
      layout: "fitData",
      layoutColumnsOnNewData: true,
      height: "82%",
      selectable: true,
    };

    this.tabulator = new Tabulator("#tag-table-container", tag_options);
    this.tabulator.on("rowClick", this.onRowClick);
    this.tabulator.on("tableBuilt", () => {
      this.fetchTags();
    });
    this.tabulator.on("cellEdited", (cell) => {
      const editedTag = cell.getRow().getData();
      this.updateTagInAPI(editedTag);
    });
  },

  methods: {
    _onTableBuilt() {
      this.fetchTags();
    },

    fetchTags() {
      this.workers
        .refreshTags()
        .then(() => {
          this.tags = this.workers.tags;
          this.tabulator.setData(this.tags);
        })
        .catch((error) => {
          const errorMsg = JSON.stringify(error);
          useNotifs().add(`Error: ${errorMsg}`);
        });
    },

    createTag(event) {
      event.preventDefault();

      const api = new WorkerMgtApi(getAPIClient());
      const newTag = new WorkerTag(this.newTagName);

      newTag.description = "Default Description...";

      api
        .createWorkerTag(newTag)
        .then(() => {
          this.fetchTags(); // Refresh table data
          this.newTagName = "";
        })
        .catch((error) => {
          const errorMsg = JSON.stringify(error);
          useNotifs().add(`Error: ${errorMsg}`);
        });
    },

    updateTagInAPI(tag) {
      const { id: tagId, ...updatedTagData } = tag;
      const api = new WorkerMgtApi(getAPIClient());

      api
        .updateWorkerTag(tagId, updatedTagData)
        .then(() => {
          // Update the local state with the edited data without requiring a page refresh
          this.tags = this.tags.map((tag) => {
            if (tag.id === tagId) {
              return { ...tag, ...updatedTagData };
            }
            return tag;
          });
          console.log("Tag updated successfully");
        })
        .catch((error) => {
          const errorMsg = JSON.stringify(error);
          useNotifs().add(`Error: ${errorMsg}`);
        });
    },

    deleteTag() {
      if (!this.selectedTag) {
        return;
      }

      const api = new WorkerMgtApi(getAPIClient());
      api
        .deleteWorkerTag(this.selectedTag.id)
        .then(() => {
          this.selectedTag = null;
          this.tabulator.setData(this.tags);

          this.fetchTags();
        })
        .catch((error) => {
          const errorMsg = JSON.stringify(error);
          useNotifs().add(`Error: ${errorMsg}`);
        });
    },

    onRowClick(event, row) {
      const tag = row.getData();
      const rowIndex = row.getIndex();

      this.tabulator.deselectRow();
      this.tabulator.selectRow(rowIndex);

      this.selectedTag = tag;
      this.activeRowIndex = rowIndex;
    },
  },
};
</script>
