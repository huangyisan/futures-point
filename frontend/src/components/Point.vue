<template>
<!--  <el-space wrap>-->
  <el-card class="box-card">

    <div>{{ buySide }}</div>
    <el-row>
    <el-col :span="11">
    <div v-for="(item,index) in sellItems" :key="index" class="text item">{{ item }}</div>
    </el-col>
      <el-col :span="2">
        <el-divider direction="vertical" content-position="center" />
      </el-col>
    <el-col :span="11">
<!--      <div>{{ buySide }}</div>-->
    <div v-for="(item,index) in props.res.buyData" :key="index" class="text item">{{ item }}</div>
    </el-col>
    </el-row>
    <el-divider direction="horizontal" />
    <button @click="printSide">1111</button>
  </el-card>
<!--  </el-space>-->
  <!-- -->
</template>

<script setup>
import {watch, ref} from "vue";

const props = defineProps(['res'])
let buyItems = ref([])
let buySide = ref("ccc")
let sellItems = ref([])
let sellSide = ref()
// let buyItems = props.res.buyData;
// let sellItems = props.res.sellData;
// let buySide = ref(props.res.buyData.side)
// let sellSide = ref(props.res.sellData.side)
// let tmp = ref()
function printSide() {
  console.log(buyItems)
  console.log(buyItems[0])
}


watch(() => props.res.buyData, (newVal, oldVal) => {
  buySide = newVal.side
  buyItems = newVal.data
  console.log(buySide)
  console.log(buyItems)
  // tmp = newVal.split('\n');
  // side = tmp[0]
  // buyItems = tmp.slice(1);

})

watch(() => props.res.sellData, (newVal, oldVal) => {
  // tmp = newVal.split('\n');
  // side = tmp[0]
  // sellItems = tmp.slice(1);
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