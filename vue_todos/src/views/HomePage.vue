<script setup lang="ts">
import { ref } from 'vue'
import TodoCard from '../components/TodoCard.vue'
import ToDo from '../models/ToDo'

type Todos = {
  [key: string]: ToDo
}

const todos: Todos = ref({
  '0': { id: 0, title: 'Todo0', description: 'descritpion0', completed: false },
  '1': { id: 1, title: 'Todo1', description: 'descritpion1', completed: false },
  '2': { id: 2, title: 'Todo2', description: 'descritpion2', completed: false },
  '3': { id: 3, title: 'Todo3', description: 'descritpion3', completed: false },
  '4': { id: 4, title: 'Todo4', description: 'descritpion4', completed: false },
  '5': { id: 5, title: 'Todo5', description: 'descritpion5', completed: false },
  '6': { id: 6, title: 'Todo6', description: 'descritpion6', completed: false },
  '7': { id: 7, title: 'Todo7', description: 'descritpion7', completed: false },
  '8': { id: 8, title: 'Todo8', completed: false },
  '9': { id: 9, title: 'Todo9', completed: false },
  '10': { id: 10, title: 'Todo10', description: 'descritpion10', completed: false },
  '11': { id: 11, title: 'Todo11', description: 'descritpion11', completed: false },
  '12': { id: 12, title: 'Todo12', description: 'descritpion12', completed: false },
  '13': { id: 13, title: 'Todo13', description: 'descritpion13', completed: false },
  '14': { id: 14, title: 'Todo14', completed: false },
  '15': { id: 15, title: 'Todo15', completed: false }
})

</script>

<template>
  <el-row :gutter="20">
    <div v-for="todo in todos" :key="todo.id">
      <TodoCard :todo="todo" @todo-done="(n: ToDo) => {n.completed = true}" />
    </div>
  </el-row>
</template>

<style>
.el-row {
  margin-top: 70px;
  margin-bottom: 50px;
}
</style>
