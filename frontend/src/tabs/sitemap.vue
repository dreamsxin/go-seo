<template>
    <el-form :model="form" label-width="auto">
        <el-form-item label="网址">
            <el-input v-model="form.dsturl" placeholder="请输入网址" />
        </el-form-item>
        <el-form-item label="保存路径">
            <el-input v-model="form.savepath" placeholder="请选择路径">
                <template #append><el-button @click="selectSavepath">选择目录</el-button></template>
            </el-input>
        </el-form-item>
        <el-form-item label="生成文件名">
            <el-input v-model="form.filename" placeholder="请输入文件名" />
        </el-form-item>
        <el-form-item>
            <el-col :span="8">
                <el-form-item label="并发数">
                    <el-input-number v-model="form.concurrency" :min="1" :max="100" />
                </el-form-item>
            </el-col>
            <el-col :span="8">
                <el-form-item label="爬取总时长（秒）">
                    <el-input-number v-model="form.crawltimeout" :min="1" :max="3600" />
                </el-form-item>
            </el-col>
            <el-col :span="8">
                <el-form-item label="单个请求超时（秒）">
                    <el-input-number v-model="form.requesttimeout" :min="1" :max="300" />
                </el-form-item>
            </el-col>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="startGenerateSitemap">开始</el-button>
            <el-button>Cancel</el-button>
        </el-form-item>
    </el-form>
</template>

<script lang="ts" setup>
import { EventsOn, EventsOff } from '../../wailsjs/runtime/runtime'
import { defineProps, ref, reactive, onMounted, onUnmounted } from 'vue'
import { ElLoading, ElNotification } from 'element-plus'
import { OpenDirectoryDialog, GenerateSitemap } from '../../wailsjs/go/main/App'

const emit = defineEmits(['run'])

// do not use same name with ref
const form = reactive({
    dsturl: '',
    savepath: '',
    filename: 'sitemap.xml',
    concurrency: 10,
    crawltimeout: 1800,
    requesttimeout: 10,
})

const onSubmit = () => {
    console.log('submit!')
}

const selectSavepath = () => {
    OpenDirectoryDialog().then((path) => {
        form.savepath = path
    })
}

const startGenerateSitemap = () => {
    emit('run')
    GenerateSitemap(form).then(() => {

    })
}
</script>