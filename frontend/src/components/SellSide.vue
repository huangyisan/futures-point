<template>
  <div>
    <div v-for="(item,index) in items" :key="index">{{ item }}</div>
  </div>

  <el-card class="box-card">
    <div>sell side</div>
    <div v-for="(item,index) in items" :key="index" class="text item">{{ item }}</div>
  </el-card>

    <button @click="printRes"> print </button>
</template>

<script setup>
import {watchEffect} from "vue";


const props = defineProps(['sellRes'])

let items = props.sellRes.split('\n');


watchEffect(() => {
  items = props.sellRes.split('\n');
})

function printRes() {
  console.log(props.sellRes)
}
</script>

<style scoped>

</style>