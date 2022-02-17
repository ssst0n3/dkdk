<template>
  <div>
    <b-card class="border-0">
      <div class="text-left">
        <b-link class="text-info small" @click="downloadTemplate">download template file</b-link>
      </div>
      <b-form-file
          v-model="file"
          :state="Boolean(file)"
          placeholder="Upload Json Format Tasks"
          drop-placeholder="Drop file here..."
          class="mt-3"
      ></b-form-file>
      <div class="text-right">
        <b-btn class="mt-3" @click="batchTaskCreate">Submit</b-btn>
      </div>
    </b-card>
  </div>
</template>

<script>
import lightweightRestful from "vue-lightweight_restful";
import consts from "@/consts";

export default {
  name: "Batch",
  data() {
    return {
      file: null,
      template: `[
  {
      "repository_id": 0,
      "type": 0,
      "origin_url": "",
      "filename_in_dkdk": "",
      "directory_in_dkdk": 0,
      "filename": ""
  }
]`
    }
  },
  methods: {
    downloadTemplate() {
      const blob = new Blob([this.template], {type: 'application/octet-stream'})
      const link = document.createElement('a')
      link.href = URL.createObjectURL(blob)
      link.download = "template.json"
      link.click()
      URL.revokeObjectURL(link.href)
    },
    batchTaskCreate() {
      const reader = new FileReader();
      reader.onload = async e => await lightweightRestful.api.post(consts.api.v1.task.batchTaskCreate, null, e.target.result, {
        caller: this,
      })
      reader.readAsText(this.file)
    }
  }
}
</script>

<style scoped>

</style>
