<template>
  <div>
    <b-form-group>
      <b-input-group prepend="username">
        <b-form-input name="username" v-model="model.username"/>
      </b-input-group>
      <b-input-group class="mt-3" prepend="password">
        <b-form-input name="password" v-model="model.password"/>
      </b-input-group>
      <b-form-checkbox class="text-right mt-3" name="is_admin" v-model="model.is_admin" :value="true"
                       :unchecked-value="false">
        Is Admin
      </b-form-checkbox>
      <div class="text-right">
        <b-button class="mt-3" :variant="button_variant" :disabled="loading" @click="click">
          <b-spinner small type="grow" v-show="loading"></b-spinner>
          {{ button_text }}
        </b-button>
      </div>
    </b-form-group>
  </div>
</template>

<script>
import consts from '@/consts'

export default {
  name: "CreateUser",
  data() {
    return {
      loading: false,
      model: {
        is_admin: true
      },
      button_text: 'Create User',
      button_variant: 'info',
    }
  },
  methods: {
    async click() {
      this.loading = true
      this.button_variant = 'info'
      this.button_text = 'Loading'
      await this.create_user()
      this.button_variant = 'info'
      this.button_text = 'Create User'
      this.loading = false
    },
    async create_user() {
      await this.$rest_api.post(consts.api.v1.initialize.create_user, null, this.model, {
        caller: this
      })
    }
  }
}
</script>

<style scoped>

</style>
