<template>
    <el-form :model="form" label-width="auto">
        <el-form-item label="网址">
            <el-input v-model="form.url" placeholder="请输入网址" />
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
            <el-button type="primary" @click="onSubmit">Create</el-button>
            <el-button>Cancel</el-button>
        </el-form-item>
    </el-form>
</template>

<script lang="ts" setup>
import { reactive } from 'vue'
import { OpenDirectoryDialog } from '../../wailsjs/go/main/App'

// do not use same name with ref
const form = reactive({
    url: '',
    savepath: '',
    filename: 'sitemap.xml',
})

const onSubmit = () => {
    console.log('submit!')
}

const selectSavepath = () => {
    OpenDirectoryDialog().then((path) => {
        form.savepath = path
    })
}
</script>