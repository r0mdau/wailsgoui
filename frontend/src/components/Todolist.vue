<script setup>
import {reactive} from 'vue'
import {Add} from '../../wailsjs/go/main/App'
import {Reset} from '../../wailsjs/go/main/App'
import {Remove} from '../../wailsjs/go/main/App'
import {Load} from '../../wailsjs/go/main/App'

const data = reactive({
  items: {},
})


Load().then(result => {
  data.items = result
  data.name= ""
})

function add() {
  Add(data.name).then(result => {
    data.items = result
    data.name= ""
  })
}

function reset() {
  Reset().then(result => {
    data.items = result
    data.name= ""
  })
}

function remove(index) {
  Remove(parseInt(index)).then(result => {
    data.items = result
    data.name= ""
  })
}

</script>

<template>
  <div class="container-sm">
    <div id="result" class="result">Manage todolist items</div>
  </div>
  <div class="form-inline">
    <div class="form-group">
      <label for="item">Item</label>
      <input id="name" v-model="data.name" autocomplete="off" class="form-control" type="text" v-on:keyup.enter="add"/>
    </div>
    <button class="btn btn-primary" @click="add">Add</button>
    <button class="btn btn-danger" @click="reset">Reset</button>
  </div>
  <table class="table table-striped">
    <thead>
      <tr>
        <th scope="col">#</th>
        <th scope="col">Item</th>
        <th scope="col">Action</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(item, index) in data.items">
        <td>{{ index }}</td>
        <td>{{ item }}</td>
        <td>
          <div class="btn btn-sm btn-danger" @click="remove(index)">
            <i class="bi-trash"></i>
          </div>
        </td>
      </tr>
    </tbody>
  </table>
</template>
