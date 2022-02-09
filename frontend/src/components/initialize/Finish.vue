<template>
  <div>
    <b-card-text>
      <p> Are you sure to
        <b-spinner small type="grow" v-show="loading"></b-spinner>
        <b-button @click="click" variant="info"> {{ loading ? 'Loading' : 'FINISH' }}</b-button>
        initialize stage ?
      </p>
    </b-card-text>
  </div>
</template>

<script>
import consts from '@/consts'

export default {
  name: "Finish",
  data() {
    return {
      loading: false,
    }
  },
  methods: {
    async click() {
      this.loading = true
      await this.finish()
      this.loading = false
    },
    async finish() {
      let response = await this.$rest_api.post(consts.api.v1.initialize.end, null, null, {
        caller: this
      })
      if (response.success) {
        await this.$router.push(consts.Path.home)
      }
    }
  }
}
</script>

<style scoped>

</style>
