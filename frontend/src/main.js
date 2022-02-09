import Vue from 'vue'
import App from './App.vue'
import router from "@/router";

import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
Vue.use(BootstrapVue)
Vue.use(IconsPlugin)

import lightweightRestful from "vue-lightweight_restful";
import consts from "@/consts";
lightweightRestful.api.initClient(consts.BaseUrl)
Vue.use(lightweightRestful)

import VueClipboard from 'vue-clipboard2'
VueClipboard.config.autoSetContainer = true
Vue.use(VueClipboard)

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App),
  data() {
    return {
      should_initialize: false,
    }
  }
}).$mount('#app')
