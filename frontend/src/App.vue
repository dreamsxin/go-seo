<script setup>
import { EventsOn } from '../wailsjs/runtime/runtime'
import { ref, reactive, useTemplateRef, watch, onMounted, computed } from 'vue'
import { ElNotification } from 'element-plus'
import { Aim } from '@element-plus/icons-vue'
import { GetConfig, SetConfig } from '../wailsjs/go/main/App'
import SitemapTab from './tabs/sitemap.vue'
import ConfigTab from './tabs/config.vue'


onMounted(() => {
  GetConfig().then(config => {
  })
})

const activeName = ref('first')

const handleClick = (tab, event) => {
  console.log(tab, event)
}

EventsOn("error", function (v) {
  ElNotification({
    title: 'Error',
    message: v.Message,
    type: 'error',
  })
});

</script>

<template>
  <el-tabs type="border-card" class="demo-tabs" height="100vh">
    <el-tab-pane>
      <template #label>
        <span class="custom-tabs-label">
          <el-icon>
            <Aim />
          </el-icon>
          <span>网站地图</span>
        </span>
      </template>
      <SitemapTab></SitemapTab>
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