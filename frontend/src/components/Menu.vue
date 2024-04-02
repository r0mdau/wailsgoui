<script setup>
import {reactive} from 'vue'
import {GetDatastorePath} from '../../wailsjs/go/main/App'
import {ChangeDatastorePath} from '../../wailsjs/go/main/App'

const data = reactive({
  datastore: "",
})

GetDatastorePath().then(result => {
  data.datastore = result
})

function apply() {
  ChangeDatastorePath(data.datastore).then(result => {
    data.datastore = result
  })
}

</script>
<template>
    <div class="mb-3">
        <div class="input-group">
            <span class="input-group-text" id="basic-addon3">
                Datastore path
            </span>
            <input type="text" v-on:keyup.enter="apply" v-model="data.datastore" class="form-control" id="datastore-path" aria-describedby="basic-addon3 basic-addon4">
            <button class="btn btn-outline-secondary" type="button" @click="apply">Apply</button>
        </div>
    </div>
</template>