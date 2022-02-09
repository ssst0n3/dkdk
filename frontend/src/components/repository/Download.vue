<template>
  <div class="d-inline">
    <b-button size="sm"
              @click="download(loggedIn, repositoryId, filename, digest, repository_config, password!==undefined&&password.length>0)"
              class="mr-2">
      <b-spinner v-show="downloading" small type="grow"></b-spinner>
      Download
      <b-badge v-if="download_count!==undefined" class="ml-1" variant="light"> {{download_count}} <span class="sr-only">download count</span></b-badge>
    </b-button>
    <b-btn size="sm" :id="'copy_password'+digest" class="ml-2" v-show="password!==undefined&&password.length>0"
           v-clipboard:copy="password" v-clipboard:success="OnCopySuccess">Copy Password
    </b-btn>
    <b-tooltip triggers="click" :target="'copy_password'+digest" ref="tooltip">
      Copied!
    </b-tooltip>
    <b-modal :id="'show_link'+digest" :title="`download [${filename}]`" hide-footer>
      <div style="text-align: center">
        <b-btn @click="download_from_link">Download</b-btn>
        <b-btn id="copy" class="ml-2" v-clipboard:copy="link" v-clipboard:success="OnCopySuccess">Copy Address</b-btn>
        <b-tooltip triggers="click" target="copy" ref="tooltip">
          Copied!
        </b-tooltip>
      </div>
    </b-modal>
  </div>
</template>

<script>
// import lightweightRestful from "vue-lightweight_restful";
// import {saveAs} from 'file-saver'
import consts from "@/consts";
import axios from "axios";

export default {
  name: "Download",
  props: {
    repositoryId: Number,
    filename: String,
    digest: String,
    repository_config: Object,
    password: String,
    download_count: Number,
  },
  data() {
    return {
      link: '',
      loggedIn: false,
      downloading: false,
    }
  },
  mounted() {
    if (this.$root.$children[0].$refs.nav.$refs.auth.loggedIn) {
      this.loggedIn = true
    } else {
      if (this.$root.should_initialize === false) {
        this.$bvModal.show('sign')
      }
    }
  },
  methods: {
    async download(loggedIn, repositoryId, filename, digest, repository_config, archived) {
      this.downloading = true
      let response
      if (loggedIn) {
        let url = consts.BaseUrl + consts.api.v1.repository.download_by_id(repositoryId, digest)
        response = await axios.get(url, {responseType: 'arraybuffer', withCredentials: true}).then((resp) => {
          this.$bvToast.toast(resp.data.msg, {
            title: 'download',
            variant: 'success',
            solid: true
          })
          return resp.data
        })
      } else {
        let url = consts.BaseUrl + consts.api.v1.repository.download(digest)
        response = await axios.post(url, repository_config, {
          responseType: 'arraybuffer',
          withCredentials: true
        }).then((resp) => {
          this.$bvToast.toast(resp.data.msg, {
            title: 'download',
            variant: 'success',
            solid: true
          })
          return resp.data
        })
      }
      let header = String.fromCharCode.apply(null, new Uint8Array(response.slice(0, 4)))
      if (header.startsWith("http")) {
        this.link = String.fromCharCode.apply(null, new Uint8Array(response))
        this.$bvModal.show('show_link' + this.digest)
      } else {
        const blob = new Blob([response], {type: 'application/octet-stream'})
        const link = document.createElement('a')
        link.href = URL.createObjectURL(blob)
        link.download = filename
        if (archived) {
          link.download += ".zip"
        }
        link.click()
        URL.revokeObjectURL(link.href)
      }
      this.downloading = false
      if (this.download_count !== undefined) {
        this.$emit("add_download_count")
      }
    },
    download_from_link() {
      // saveAs(this.link, this.filename)
      const a = document.createElement('a')
      a.href = this.link
      a.download = this.filename
      a.click()
    },
    OnCopySuccess() {
      this.$refs.tooltip.$emit('open')
    },
  }
}
</script>

<style scoped>

</style>
