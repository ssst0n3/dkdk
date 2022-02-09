<template>
  <div>
    <b-input-group prepend="address">
      <b-form-input v-model="address" class="col" placeholder="eg: 127.0.0.1:14005/dkdk/dkdk:v0.1"/>
    </b-input-group>
    <div class="row mt-3">
      <div class="col"/>
      <b-form-checkbox class="col-3" v-model="model.insecure">insecure</b-form-checkbox>
      <b-form-checkbox class="col-3" v-model="model.private">private</b-form-checkbox>
    </div>
    <div v-if="model.private">
      <b-input-group prepend="username" class="mt-3">
        <b-form-input v-model="model.username"/>
      </b-input-group>
      <b-input-group prepend="password" class="mt-3">
        <b-form-input type="password" v-model="model.secret"/>
        <template #append v-if="update">
          <b-btn @click="update_secret">Update</b-btn>
        </template>
      </b-input-group>
    </div>
    <div class="mt-3 row" style="float: right">
      <b-btn size="sm" @click="clean">Clean</b-btn>
      <b-btn size="sm" v-if="IsConfig" @click="submit" class="ml-2">
        <span>Submit <small v-if="update">(No Secret Include)</small></span>
      </b-btn>
      <div v-else>
        <b-btn class="ml-2 mr-2" @click="fetch">fetch</b-btn>
      </div>
    </div>
  </div>
</template>

<script>
import lightweightRestful from "vue-lightweight_restful";
import consts from "@/consts";

export default {
  name: "RepositoryConfig",
  props: {
    IsConfig: Boolean,
    update: Boolean,
    Repository: Object,
  },
  data: function () {
    return {
      address: '',
      model: {
        service_address: '',
        insecure: false,
        username: '',
        secret: '',
        name: '',
        reference: '',
        private: false,
      }
    }
  },
  methods: {
    parse() {
      let slash = this.address.indexOf("/")
      let colon = this.address.lastIndexOf(":")
      this.model.service_address = this.address.slice(0, slash)
      if (colon > 0) {
        this.model.name = this.address.slice(slash + 1, colon)
        this.model.reference = this.address.slice(colon + 1, this.address.length)
      } else {
        this.model.name = this.address.slice(slash + 1, this.address.length)
        this.model.reference = ''
      }
    },
    submit() {
      this.parse()
      if (this.update) {
        lightweightRestful.api.updateResource(consts.api.v1.repository.repository, this.model.id, this.model, {
          caller: this,
        })
      } else {
        lightweightRestful.api.createResource(consts.api.v1.repository.repository, this.model, {
          caller: this,
        })
      }
    },
    fetch() {
      this.parse()
      this.$emit('fetch')
    },
    clean() {
      this.model = {}
      this.$emit('clean')
    },
    update_secret() {
      lightweightRestful.api.updateResource(consts.api.v1.repository.repository, this.model.id, {"secret": this.model.secret}, {
        caller: this,
      })
    }
  }
}
</script>

<style scoped>

</style>
