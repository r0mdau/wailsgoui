<script setup>
import {reactive} from 'vue'
import {Add} from '../../wailsjs/go/main/App'
import {Reset} from '../../wailsjs/go/main/App'

const data = reactive({
  name: "",
  resultText: "",
})

function add() {
  Add(data.name).then(result => {
    data.items = result
    data.name = ""
  })
}

function reset() {
  Reset().then(result => {
    data.items = result
    data.name = ""
  })
}

</script>

<template>
  <main>
    <div id="result" class="result">Manage todolist items</div>
    <div id="input" class="input-box">
      <input id="name" v-model="data.name" autocomplete="off" class="input" type="text" v-on:keyup.enter="add"/>
      <button class="btn" @click="add">Add</button>
      <button class="btn" @click="reset">Reset</button>
    </div>
    <div id="todolist" class="todolist">
      <ul>
        <li v-for="item in data.items" :key="item.id">
          <span>{{ item }}</span>
        </li>
      </ul>
    </div>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
