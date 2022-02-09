<template>
  <div style="border: darkcyan 1px dashed; padding: 20px">
    <b-form-file
        v-model="file"
        :state="Boolean(file)"
        placeholder="Choose a file or drop it here..."
        drop-placeholder="Drop file here..."
    />
    <div class="mt-3">Selected file: {{ file ? file.name : '' }}</div>
    <div style="text-align: right">
      <b-btn @click="upload">upload</b-btn>
    </div>
  </div>
</template>

<script>
import consts from "@/consts";
// import axios from "axios";
import lightweightRestful from "vue-lightweight_restful";

export default {
  name: "Uploader",
  props: {
    repository: Object,
  },
  data() {
    return {
      file: null,
    }
  },
  methods: {
    async upload() {
      let api
      let formData = new FormData()
      if (this.$root.$children[0].$refs.nav.$refs.auth.loggedIn) {
        let id = this.$parent.$refs.repository_config_id.selected
        api = consts.BaseUrl + consts.api.v1.repository.upload_by_id(id)
      } else {
        api = consts.BaseUrl + consts.api.v1.repository.upload
        let repository_config = JSON.stringify(this.repository)
        formData.append('repository', repository_config)
      }
      formData.append('file', this.file)
      let resp = await lightweightRestful.api.client.post(
          api,
          formData,
          {
            headers: {
              'Content-Type': 'multipart/form-data'
            }
          }
      )
      let msg = resp.data.msg
      this.$bvToast.toast(msg, {
        title: api,
        variant: 'success',
        solid: true
      })
      console.log(resp)
    }
  }
}
</script>

<style scoped>

</style>
