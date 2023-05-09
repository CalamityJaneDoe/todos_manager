<script>
import {defineComponent} from 'vue'
import { Plus } from '@element-plus/icons-vue'

import TodoForm from '../TodoForm.vue'

export default defineComponent({
  components: {
    TodoForm
  }, 
  data() {
    return {
      dialogFormVisible: false
    };
  },
})
</script>

<template>
  <nav id="vue">
    <!-- @TODO : understand why docker can't build 'cos of the import'-->
    <!-- <img class="logo" src="src/assets/logo.png" alt="Logo for to do manager" /> -->
    <h2 class="nav-title">Todo Manager</h2>
    <div class="menu-item"><router-link to="/">Home</router-link></div>
    <div class="menu-item"><router-link to="/stats">Stats</router-link></div>
    <div class="menu-button">
      <el-button type="primary" color="#073B4C" @click="dialogFormVisible = true"><el-icon>
          <Plus />
        </el-icon>Add to do</el-button>
      <el-dialog v-model="dialogFormVisible" title="Add To do"><TodoForm /></el-dialog>
    </div>
  </nav>
</template>

<style>
nav {
  position: fixed;
  left: 0;
  top: 0;
  width: 100vw;
  background-color: #118ab2;
  display: flex;
  align-items: center;
  justify-content: left;
  z-index: 200;
  height: 70px;
}

nav .menu-item {
  color: #fff;
  padding: 25px 25px;
  position: relative;
  text-align: center;
  border-bottom: 5px solid transparent;
  display: flex;
  transition: 0.4s;
}

nav .menu-item a {
  color: inherit;
  text-decoration: none;
}

nav .menu-button {
  position: absolute;
  top: 20px;
  right: 50px;
}

/* nav .logo {
  height: 65px;
} */

nav .nav-title {
  padding-left: 15px;
  color : black
}
</style>
