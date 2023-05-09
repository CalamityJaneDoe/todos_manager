<script lang="ts" setup>
const tableData = [
  {
    title: '2016-05-03',
    description: 'Tom',
  },
  {
    title: '2016-05-02',
    description: 'Tom',
  }

]
</script>

<template>
    <el-table :data="tableData" height="250" style="width: 100%">
      <el-table-column prop="title" label="Title" width="180" />
      <el-table-column prop="description" label="Description" width="180" />
    </el-table>
  </template>