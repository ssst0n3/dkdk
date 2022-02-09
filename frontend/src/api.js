import lightweightRestful from "vue-lightweight_restful";
import consts from "@/consts";

export default {
    async check_initialize() {
        let response = await lightweightRestful.api.get(consts.api.v1.initialize.initialize)
        return response.should_initialize
    },
}
