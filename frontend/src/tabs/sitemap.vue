<template>
    <el-form label-width="auto">
        <el-form-item label="网址">
            <el-input v-model="data.config.Dsturl" placeholder="请输入网址" />
        </el-form-item>
        <el-form-item label="保存路径">
            <el-input v-model="data.config.Savepath" placeholder="请选择路径">
                <template #append><el-button @click="selectSavepath">选择目录</el-button></template>
            </el-input>
        </el-form-item>
        <el-form-item label="生成文件名">
            <el-input v-model="data.config.Filename" placeholder="请输入文件名" />
        </el-form-item>
        <el-form-item>
            <el-col :span="8">
                <el-form-item label="并发数">
                    <el-input-number v-model="data.config.Concurrency" :min="1" :max="100" />
                </el-form-item>
            </el-col>
            <el-col :span="8">
                <el-form-item label="爬取总时长（秒）">
                    <el-input-number v-model="data.config.Crawltimeout" :min="1" :max="3600" />
                </el-form-item>
            </el-col>
            <el-col :span="8">
                <el-form-item label="单个请求超时（秒）">
                    <el-input-number v-model="data.config.Requesttimeout" :min="1" :max="300" />
                </el-form-item>
            </el-col>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="startGenerateSitemap">开始</el-button>
        </el-form-item>
    </el-form>
</template>

<script lang="ts" setup>
import { EventsOn, EventsOff } from '../../wailsjs/runtime/runtime'
import { ref, reactive, onMounted, onUnmounted } from 'vue'
import { ElLoading, ElNotification } from 'element-plus'
import { GetSitemapConfig, OpenDirectoryDialog, GenerateSitemap } from '../../wailsjs/go/main/App'

const data = reactive({
    config: {
        Dsturl: '',
        Savepath: '',
        Filename: 'sitemap.xml',
        Concurrency: 10,
        Crawltimeout: 1800,
        Requesttimeout: 10,
    }
})
const emit = defineEmits(['run'])

const selectSavepath = () => {
    OpenDirectoryDialog().then((path) => {
        data.config.Savepath = path
    })
}

const startGenerateSitemap = () => {
    emit('run')
    GenerateSitemap(data.config).then(() => {

    })
}


onMounted(() => {
    GetSitemapConfig().then(config => {
        console.log("GetConfig", config)
        data.config = config
    })
})
</script>