<script lang="ts">
import { defineComponent, type PropType } from 'vue'
import { Check, Edit, Warning } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import TodoForm from './TodoForm.vue'

import type ToDo from '../models/ToDo'

export default defineComponent({
  components: {
    TodoForm
  }, 
    props: {
        todo: Object as PropType<ToDo>
    },
    mounted() {
        this.todo;
    },
    data() {
        return {
            dialogFormVisible: false
        };
    },
    methods: {
        openDeleteDialog() {
            ElMessageBox.confirm("Are you sure you want to delete this todo ?", "Delete todo", {
                confirmButtonText: "Yes",
                cancelButtonText: "No",
                type: "warning"
            })
                .then(() => {
                // @TODO : insert delete methods here !
                console.log("insert here delete methods");
                ElMessage({
                    type: "success",
                    message: "Delete completed",
                });
            })
                .catch(() => {
                ElMessage({
                    type: "info",
                    message: "Delete canceled",
                });
            });
        }
    },
})
</script>

<template>
  <el-col :span="6">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>{{ todo?.title }}</span>
          <div>
            <el-button class="edit-button" type="primary" color="#073B4C" @click="dialogFormVisible = true" >
              <el-icon>
                <Edit />
              </el-icon>
            </el-button>
            <el-dialog v-model="dialogFormVisible" title="Edit To do">
              <TodoForm />
            </el-dialog>
            <el-button type="primary" color="#06D6A0" @click="$emit('todoDone', todo)">
              <el-icon>
                <Check />
              </el-icon>
            </el-button>
          </div>
        </div>
      </template>
      <div>{{ todo?.description || 'No description for this todo' }}</div>
      <el-button class="delete-button" type="primary" color="#EF476F" @click="openDeleteDialog">
        <el-icon>
          <Delete />
        </el-icon>
      </el-button>
    </el-card>
  </el-col>
</template>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}


.box-card {
  width: 250px;
  height: 200px;
}

.el-col {
  margin-top: 20px;
  margin-bottom: 20px;
}
.delete-button {
  margin-top : 40px;
}
.edit-button {
  margin-right: 10px;
}
</style>
