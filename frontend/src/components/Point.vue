<template>
<!--  <el-space wrap>-->
  <el-card class="box-card">

    <div>{{ side }}</div>
    <el-row>
    <el-col :span="11">
    <div v-for="(item,index) in items" :key="index" class="text item">{{ item }}</div>
    </el-col>
      <el-col :span="2">
        <el-divider direction="vertical" content-position="center" />
      </el-col>
    <el-col :span="11">
    <div v-for="(item,index) in items" :key="index" class="text item">{{ item }}</div>
    </el-col>
    </el-row>
    <el-divider direction="horizontal" />
    <button @click="printSide">1111</button>
  </el-card>
<!--  </el-space>-->
  <!-- -->
</template>

<script setup>
import {watch, ref, computed} from "vue";

const props = defineProps(['res'])
let items = props.res.split('\n');
let side = ref()
let tmp = ref()
function printSide() {
  console.log(items)
  console.log(items[0])
}


watch(() => props.res, (newVal, oldVal) => {
  tmp = newVal.split('\n');
  side = tmp[0]
  items = tmp.slice(1);

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