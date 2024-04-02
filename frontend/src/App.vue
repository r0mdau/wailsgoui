<script setup>
import Todolist from './components/Todolist.vue'
import Menu from './components/Menu.vue'
</script>

<template>
  <div class="container-fluid">
    <div class="row">
      <div class="col-5">
        <Menu/>
      </div>
      <div class="col-7">
        <Todolist/>
      </div>
    </div>
  </div>
</template>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
