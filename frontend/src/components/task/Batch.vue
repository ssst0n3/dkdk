<template>
  <div>
    <b-card class="border-0">
      <div class="text-left">
        <b-link class="text-info small">download template file</b-link>
      </div>
      <b-form-file
          v-model="file"
          :state="Boolean(file)"
          placeholder="Upload Json Format Tasks"
          drop-placeholder="Drop file here..."
          class="mt-3"
      ></b-form-file>
      <div class="text-right">
        <b-btn class="mt-3">Submit</b-btn>
      </div>
    </b-card>
  </div>
</template>

<script>
export default {
  name: "Batch",
  data() {
    return {
      file: null
    }
  }
}
</script>

<style scoped>

</style>
