import Vue from 'vue'
import Router from "vue-router";
import Config from "@/components/config/Config";
import Index from "@/components/Index";
import Initialize from "@/components/initialize/Initialize";
import OfflineDownload from "@/components/task/OfflineDownload";
import NetDisk from "@/components/netdisk/NetDisk";

Vue.use(Router);

export default new Router({
    routes: [
        {
            path: '/',
            name: 'index',
            component: Index
        },
        {
            path: '/net_disk',
            name: 'net_disk',
            component: NetDisk,
        },
        {
            path: '/config',
            name: 'config',
            component: Config
        },
        {
            path: '/initialize',
            name: 'Initialize',
            component: Initialize,
            meta: {
                navShow: false
            }
        },
        {
            path: '/offline_download',
            name: 'OfflineDownload',
            component: OfflineDownload,
        },
    ]
})
