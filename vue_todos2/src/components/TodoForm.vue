<script lang="ts" setup>
import { reactive } from 'vue'

const formInline = reactive({
    title: '',
    description: '',
})

const onSubmit = () => {
    console.log('submit!')
}
</script>

<template>
    <el-form :inline="true" :model="formInline">
        <el-form-item label="Title">
            <el-input v-model="formInline.title" placeholder="To do title" />
        </el-form-item>
        <el-form-item label="Description">
            <el-input type="textarea" :rows="2" v-model="formInline.description" placeholder="To do description" />
        </el-form-item>
        <el-form-item>
            <el-button type="primary" color="#06D6A0" @click="onSubmit">Submit</el-button>
        </el-form-item>
    </el-form>
</template>
  
