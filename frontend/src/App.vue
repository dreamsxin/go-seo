<script setup>
import { EventsOn } from '../wailsjs/runtime/runtime'
import { ref, reactive, useTemplateRef, watch, onMounted, computed } from 'vue'
import { ElNotification } from 'element-plus'
import { Aim } from '@element-plus/icons-vue'
import { GetConfig, SetConfig } from '../wailsjs/go/main/App'
import SitemapTab from './tabs/sitemap.vue'
import ConfigTab from './tabs/config.vue'


const fullscreenLoading = ref(false)
const data = reactive({
  config: {}
})

onMounted(() => {
  GetConfig().then(config => {
    console.log("GetConfig", config)
    data.config = config
  })
})

EventsOn("response", function (v) {
  console.log("response", v)
  fullscreenLoading.value = false
  ElNotification({
    title: '',
    message: v.Message,
    type: 'success',
  })
});
EventsOn("error", function (v) {
  console.log("error", v)
  fullscreenLoading.value = false
  ElNotification({
    title: 'Error',
    message: v.Message,
    type: 'error',
  })
});


</script>

<template>
  <el-tabs type="border-card" class="demo-tabs" height="100vh"  v-loading.fullscreen.lock="fullscreenLoading"
        element-loading-background="rgba(0, 0, 0, 0.7)">
    <el-tab-pane>
      <template #label>
        <span class="custom-tabs-label">
          <el-icon>
            <Aim />
          </el-icon>
          <span>网站地图</span>
        </span>
      </template>
      <SitemapTab @run="fullscreenLoading=true" v-model="data"></SitemapTab>
    </el-tab-pane>
    <el-tab-pane label="配置">
      <ConfigTab></ConfigTab>
    </el-tab-pane>
    <el-tab-pane label="Role">Role</el-tab-pane>
    <el-tab-pane label="Task">Task</el-tab-pane>
  </el-tabs>
</template>
<style scoped>
.demo-tabs>.el-tabs__content {
  padding: 32px;
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}

.demo-tabs .custom-tabs-label .el-icon {
  vertical-align: middle;
}

.demo-tabs .custom-tabs-label span {
  vertical-align: middle;
  margin-left: 4px;
}
</style>