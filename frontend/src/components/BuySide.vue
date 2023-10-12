<template>
  <el-card class="box-card">
    <div>buy side</div>
    <div v-for="(item,index) in items" :key="index" class="text item">{{ item }}</div>
  </el-card>

<!--  <button @click="printRes"> print </button>-->
</template>

<script setup>
import {watchEffect} from "vue";

const props = defineProps(['buyRes'])
console.log(props.buyRes)

let items = props.buyRes.split('\n');


watchEffect(() => {
  items = props.buyRes.split('\n');
})

</script>

<style scoped>
.text {
  font-size: 14px;
}

.item {
  padding: 18px 0;
}

/*.box-card {*/
/*  width: 480px;*/
/*}*/
</style>